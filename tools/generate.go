package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ServiceInfo struct {
	PackageName string
	ServiceName string
	HandlerName string
	ImportPath  string
	ConfigName  string
	Folder      string
}

type ServiceConfig struct {
	Name   string `mapstructure:"name"`
	Route  string `mapstructure:"route"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Folder string `mapstructure:"folder"`
}

type Env struct {
	NodeEnv  string          `mapstructure:"node_env"`
	Port     int             `mapstructure:"port"`
	Services []ServiceConfig `mapstructure:"services"`
}

func main() {
	fmt.Println("🚀 Bắt đầu quá trình generate...")

	// Bước 1: Update thư viện sf-proto
	fmt.Println("📦 Đang update thư viện sf-proto...")
	if err := updateLibrary(); err != nil {
		log.Fatalf("Lỗi khi update thư viện: %v", err)
	}
	fmt.Println("✅ Đã update thư viện sf-proto thành công")

	// Bước 2: Đọc config và quét qua Env.Services
	fmt.Println("🔍 Đang đọc config và quét qua Env.Services...")
	services, err := scanServicesFromConfig()
	if err != nil {
		log.Fatalf("Lỗi khi đọc config và quét services: %v", err)
	}
	fmt.Printf("✅ Tìm thấy %d services từ config\n", len(services))

	// Bước 3: Tìm các Register*ServiceHandler
	fmt.Println("🔍 Đang tìm các Register*ServiceHandler...")
	handlers, err := findServiceHandlers(services)
	if err != nil {
		log.Fatalf("Lỗi khi tìm service handlers: %v", err)
	}
	fmt.Printf("✅ Tìm thấy %d service handlers\n", len(handlers))

	// Bước 4: Tạo file service_handlers.go
	fmt.Println("📝 Đang tạo file service_handlers.go...")
	if err := createServiceHandlersFile(handlers); err != nil {
		log.Fatalf("Lỗi khi tạo service_handlers.go: %v", err)
	}
	fmt.Println("✅ Đã tạo file service_handlers.go thành công")

	// Bước 5: Copy swagger files
	fmt.Println("📁 Đang copy swagger files...")
	if err := copySwaggerFiles(handlers); err != nil {
		log.Fatalf("Lỗi khi copy swagger files: %v", err)
	}
	fmt.Println("✅ Đã copy swagger files thành công")

	fmt.Println("🎉 Hoàn thành quá trình generate!")
}

func updateLibrary() error {
	cmd := exec.Command("go", "get", "-u", "github.com/anhvanhoa/sf-proto")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func scanServicesFromConfig() ([]ServiceInfo, error) {
	var services []ServiceInfo

	// Đọc config file
	configPath := "dev.config.yml"
	if _, err := os.Stat("prod.config.yml"); err == nil {
		configPath = "prod.config.yml"
	}

	// Parse YAML config (đơn giản hóa - chỉ đọc folder từ config)
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("không thể đọc config file: %v", err)
	}

	// Parse YAML để lấy services
	env, err := parseConfigYAML(string(content))
	if err != nil {
		return nil, fmt.Errorf("không thể parse config: %v", err)
	}

	// Lặp qua các services trong config
	for _, serviceConfig := range env.Services {
		if serviceConfig.Folder == "" {
			fmt.Printf("⚠️  Bỏ qua service '%s' - không có folder\n", serviceConfig.Name)
			continue
		}

		// Tạo đường dẫn import: github.com/anhvanhoa/sf-proto/gen/folder/v1
		importPath := fmt.Sprintf("github.com/anhvanhoa/sf-proto/gen/%s/v1", serviceConfig.Folder)

		// Kiểm tra xem service có tồn tại không
		serviceInfo, err := checkServiceExists(importPath, serviceConfig)
		if err != nil {
			fmt.Printf("⚠️  Bỏ qua service '%s' - không tìm thấy: %v\n", serviceConfig.Name, err)
			continue
		}

		services = append(services, *serviceInfo)
		fmt.Printf("✅ Tìm thấy service: %s (%s)\n", serviceConfig.Name, importPath)
	}

	return services, nil
}

func parseConfigYAML(content string) (*Env, error) {
	// Parse YAML đơn giản để lấy services
	env := &Env{}

	// Tìm section services
	lines := strings.Split(content, "\n")
	inServices := false
	var services []ServiceConfig
	var currentService ServiceConfig

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "services:" {
			inServices = true
			continue
		}

		if inServices {
			// Nếu gặp dòng bắt đầu với "- name:" thì bắt đầu service mới
			if strings.HasPrefix(line, "- name:") {
				if currentService.Name != "" {
					services = append(services, currentService)
				}
				currentService = ServiceConfig{}
				currentService.Name = strings.TrimSpace(strings.TrimPrefix(line, "- name:"))
			} else if strings.HasPrefix(line, "name:") {
				currentService.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
			} else if strings.HasPrefix(line, "host:") {
				currentService.Host = strings.TrimSpace(strings.TrimPrefix(line, "host:"))
			} else if strings.HasPrefix(line, "port:") {
				currentService.Port = strings.TrimSpace(strings.TrimPrefix(line, "port:"))
			} else if strings.HasPrefix(line, "route:") {
				currentService.Route = strings.TrimSpace(strings.TrimPrefix(line, "route:"))
			} else if strings.HasPrefix(line, "folder:") {
				currentService.Folder = strings.TrimSpace(strings.TrimPrefix(line, "folder:"))
			} else if line == "" || strings.HasPrefix(line, "#") {
				// Bỏ qua dòng trống hoặc comment
				continue
			} else if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
				// Nếu không phải indent thì kết thúc services section
				break
			}
		}
	}

	// Thêm service cuối cùng
	if currentService.Name != "" {
		services = append(services, currentService)
	}

	env.Services = services
	return env, nil
}

func checkServiceExists(importPath string, serviceConfig ServiceConfig) (*ServiceInfo, error) {
	// Tìm thư mục của module trong go mod cache
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", "github.com/anhvanhoa/sf-proto")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("không thể tìm thư mục sf-proto: %v", err)
	}

	// Tạo đường dẫn đến thư mục service
	serviceDir := filepath.Join(strings.TrimSpace(string(output)), "gen", serviceConfig.Folder, "v1")

	// Kiểm tra xem thư mục có tồn tại không
	if _, err := os.Stat(serviceDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("thư mục service không tồn tại: %s", serviceDir)
	}

	// Tìm file .go trong thư mục service
	goFiles, err := filepath.Glob(filepath.Join(serviceDir, "*.go"))
	if err != nil {
		return nil, fmt.Errorf("không thể quét file .go: %v", err)
	}

	// Tìm Register*ServiceHandler trong các file
	var handlerName string
	var packageName string

	for _, goFile := range goFiles {
		serviceInfo, err := parseGoFile(goFile)
		if err != nil {
			continue
		}
		if serviceInfo != nil && serviceInfo.HandlerName != "" {
			handlerName = serviceInfo.HandlerName
			packageName = serviceInfo.PackageName
			break
		}
	}

	if handlerName == "" {
		return nil, fmt.Errorf("không tìm thấy Register*ServiceHandler trong %s", serviceDir)
	}

	// Tạo service name từ config name
	serviceName := strings.ReplaceAll(serviceConfig.Name, " ", "_")
	serviceName = strings.ToLower(serviceName)

	return &ServiceInfo{
		PackageName: packageName,
		ServiceName: serviceName,
		HandlerName: handlerName,
		ImportPath:  importPath,
		ConfigName:  serviceConfig.Name,
		Folder:      serviceConfig.Folder,
	}, nil
}

func parseGoFile(filePath string) (*ServiceInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var serviceInfo *ServiceInfo

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// Tìm các function Register*ServiceHandler
			if strings.HasPrefix(x.Name.Name, "Register") && strings.HasSuffix(x.Name.Name, "ServiceHandler") {
				if serviceInfo == nil {
					serviceInfo = &ServiceInfo{
						PackageName: node.Name.Name,
						HandlerName: x.Name.Name,
						ImportPath:  getImportPath(filePath),
					}

					// Extract service name từ handler name
					serviceName := strings.TrimPrefix(x.Name.Name, "Register")
					serviceName = strings.TrimSuffix(serviceName, "ServiceHandler")
					serviceInfo.ServiceName = serviceName
				}
			}
		}
		return true
	})

	return serviceInfo, nil
}

func getImportPath(filePath string) string {
	// Convert file path to import path
	// Ví dụ: /path/to/sf-proto/gen/iot_device/v1/iot_device.pb.go -> github.com/anhvanhoa/sf-proto/gen/iot_device/v1
	parts := strings.Split(filePath, string(filepath.Separator))

	// Tìm vị trí của "gen"
	genIndex := -1
	for i, part := range parts {
		if part == "gen" {
			genIndex = i
			break
		}
	}

	if genIndex == -1 {
		return ""
	}

	// Lấy phần từ gen trở đi
	genParts := parts[genIndex:]
	// Loại bỏ file name
	genParts = genParts[:len(genParts)-1]

	return "github.com/anhvanhoa/sf-proto/" + strings.Join(genParts, "/")
}

func findServiceHandlers(services []ServiceInfo) ([]ServiceInfo, error) {
	var handlers []ServiceInfo

	for _, service := range services {
		// Kiểm tra xem có function Register*ServiceHandler không
		if service.HandlerName != "" {
			handlers = append(handlers, service)
		}
	}

	return handlers, nil
}

func createServiceHandlersFile(handlers []ServiceInfo) error {
	handlersFilePath := "src/handler/service_handlers.go"

	// Tạo imports cho các services
	var imports []string
	imports = append(imports, `"context"`)
	imports = append(imports, `"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"`)
	imports = append(imports, `"google.golang.org/grpc"`)

	// Thêm imports cho các services
	importMap := make(map[string]string)
	for _, handler := range handlers {
		alias := fmt.Sprintf("proto_%s", handler.Folder)
		importMap[handler.ImportPath] = alias
		imports = append(imports, fmt.Sprintf(`%s "%s"`, alias, handler.ImportPath))
	}

	// Tạo nội dung file
	var newContent strings.Builder

	// Package declaration
	newContent.WriteString("package handler\n\n")

	// Imports
	newContent.WriteString("import (\n")
	for _, imp := range imports {
		newContent.WriteString("\t" + imp + "\n")
	}
	newContent.WriteString(")\n\n")

	// ServiceHandler struct
	newContent.WriteString("type ServiceHandler struct {\n")
	newContent.WriteString("\tHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error\n")
	newContent.WriteString("\tSwagger string\n")
	newContent.WriteString("}\n\n")

	// GetServiceHandlers function
	newContent.WriteString("// GetServiceHandlers trả về map chứa tất cả service handlers\n")
	newContent.WriteString("func GetServiceHandlers() map[string]ServiceHandler {\n")
	newContent.WriteString("\treturn map[string]ServiceHandler{\n")

	// Thêm các handlers
	for _, handler := range handlers {
		alias := importMap[handler.ImportPath]
		// Use relative path for swagger files in Docker container
		swaggerPath := filepath.Join("swagger", handler.Folder, "v1", handler.Folder+".swagger.json")
		// Convert to forward slashes for consistency
		swaggerPath = strings.ReplaceAll(swaggerPath, "\\", "/")
		newContent.WriteString(fmt.Sprintf("\t\t\"%s\": {\n", handler.Folder))
		newContent.WriteString(fmt.Sprintf("\t\t\tHandler: %s.%s,\n", alias, handler.HandlerName))
		newContent.WriteString(fmt.Sprintf("\t\t\tSwagger: \"%s\",\n", "./"+swaggerPath))
		newContent.WriteString("\t\t},\n")
	}

	newContent.WriteString("\t}\n")
	newContent.WriteString("}\n")

	// Ghi file mới
	err := os.WriteFile(handlersFilePath, []byte(newContent.String()), 0644)
	if err != nil {
		return fmt.Errorf("không thể ghi file service_handlers.go: %v", err)
	}

	return nil
}

func copySwaggerFiles(handlers []ServiceInfo) error {
	// Tìm đường dẫn module sf-proto
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", "github.com/anhvanhoa/sf-proto")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("không thể tìm đường dẫn sf-proto module: %v", err)
	}
	moduleDir := strings.TrimSpace(string(output))

	// Tạo thư mục swagger nếu chưa có
	swaggerDir := "swagger"
	if err := os.MkdirAll(swaggerDir, 0755); err != nil {
		return fmt.Errorf("không thể tạo thư mục swagger: %v", err)
	}

	// Copy từng swagger file
	for _, handler := range handlers {
		// Đường dẫn nguồn
		sourcePath := filepath.Join(moduleDir, "gen", handler.Folder, "v1", handler.Folder+".swagger.json")

		// Đường dẫn đích
		destDir := filepath.Join(swaggerDir, handler.Folder, "v1")
		destPath := filepath.Join(destDir, handler.Folder+".swagger.json")

		// Tạo thư mục đích nếu chưa có
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("không thể tạo thư mục đích %s: %v", destDir, err)
		}

		// Copy file
		if err := copyFile(sourcePath, destPath); err != nil {
			fmt.Printf("⚠️  Không thể copy swagger file cho %s: %v\n", handler.Folder, err)
			continue
		}

		fmt.Printf("✅ Đã copy swagger file: %s\n", handler.Folder)
	}

	return nil
}

func copyFile(src, dst string) error {
	// Đọc file nguồn
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Tạo file đích
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy nội dung
	_, err = destFile.ReadFrom(sourceFile)
	return err
}
