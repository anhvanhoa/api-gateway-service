package handler

import (
	"context"

	proto_auth "github.com/anhvanhoa/sf-proto/gen/auth/v1"
	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	proto_fertilizer_type "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	proto_greenhouse "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	proto_greenhouse_installation_log "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	proto_growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	proto_growing_zone_history "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	proto_irrigation_log "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	proto_irrigation_schedule "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"
	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"
	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"
	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
	proto_media "github.com/anhvanhoa/sf-proto/gen/media/v1"
	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"
	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
	proto_plant_variety "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	proto_planting_cycle "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	proto_role "github.com/anhvanhoa/sf-proto/gen/role/v1"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
	proto_session "github.com/anhvanhoa/sf-proto/gen/session/v1"
	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"
	proto_user "github.com/anhvanhoa/sf-proto/gen/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ServiceHandler struct {
	Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	Swagger string
}

// GetServiceHandlers trả về map chứa tất cả service handlers
func GetServiceHandlers() map[string]ServiceHandler {
	return map[string]ServiceHandler{
		"auth": {
			Handler: proto_auth.RegisterAuthServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\auth\\v1\\auth.swagger.json",
		},
		"user": {
			Handler: proto_user.RegisterUserServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\user\\v1\\user.swagger.json",
		},
		"session": {
			Handler: proto_session.RegisterSessionServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\session\\v1\\session.swagger.json",
		},
		"role": {
			Handler: proto_role.RegisterRoleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\role\\v1\\role.swagger.json",
		},
		"greenhouse": {
			Handler: proto_greenhouse.RegisterGreenhouseServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\greenhouse\\v1\\greenhouse.swagger.json",
		},
		"greenhouse_installation_log": {
			Handler: proto_greenhouse_installation_log.RegisterGreenhouseInstallationLogServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\greenhouse_installation_log\\v1\\greenhouse_installation_log.swagger.json",
		},
		"growing_zone": {
			Handler: proto_growing_zone.RegisterGrowingZoneServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\growing_zone\\v1\\growing_zone.swagger.json",
		},
		"growing_zone_history": {
			Handler: proto_growing_zone_history.RegisterGrowingZoneHistoryServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\growing_zone_history\\v1\\growing_zone_history.swagger.json",
		},
		"iot_device": {
			Handler: proto_iot_device.RegisterIoTDeviceServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\iot_device\\v1\\iot_device.swagger.json",
		},
		"iot_device_history": {
			Handler: proto_iot_device_history.RegisterIoTDeviceHistoryServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\iot_device_history\\v1\\iot_device_history.swagger.json",
		},
		"device_type": {
			Handler: proto_device_type.RegisterDeviceTypeServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\device_type\\v1\\device_type.swagger.json",
		},
		"sensor_data": {
			Handler: proto_sensor_data.RegisterSensorDataServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\sensor_data\\v1\\sensor_data.swagger.json",
		},
		"plant_variety": {
			Handler: proto_plant_variety.RegisterPlantVarietyServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\plant_variety\\v1\\plant_variety.swagger.json",
		},
		"planting_cycle": {
			Handler: proto_planting_cycle.RegisterPlantingCycleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\planting_cycle\\v1\\planting_cycle.swagger.json",
		},
		"harvest_record": {
			Handler: proto_harvest_record.RegisterHarvestRecordServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\harvest_record\\v1\\harvest_record.swagger.json",
		},
		"pest_disease_record": {
			Handler: proto_pest_disease_record.RegisterPestDiseaseRecordServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\pest_disease_record\\v1\\pest_disease_record.swagger.json",
		},
		"irrigation_schedule": {
			Handler: proto_irrigation_schedule.RegisterIrrigationScheduleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\irrigation_schedule\\v1\\irrigation_schedule.swagger.json",
		},
		"irrigation_log": {
			Handler: proto_irrigation_log.RegisterIrrigationLogServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\irrigation_log\\v1\\irrigation_log.swagger.json",
		},
		"fertilizer_type": {
			Handler: proto_fertilizer_type.RegisterFertilizerTypeServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\fertilizer_type\\v1\\fertilizer_type.swagger.json",
		},
		"fertilizer_schedule": {
			Handler: proto_fertilizer_schedule.RegisterFertilizerScheduleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\fertilizer_schedule\\v1\\fertilizer_schedule.swagger.json",
		},
		"maintenance_schedule": {
			Handler: proto_maintenance_schedule.RegisterMaintenanceScheduleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\maintenance_schedule\\v1\\maintenance_schedule.swagger.json",
		},
		"environmental_alert": {
			Handler: proto_environmental_alert.RegisterEnvironmentalAlertServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\environmental_alert\\v1\\environmental_alert.swagger.json",
		},
		"status_history": {
			Handler: proto_status_history.RegisterStatusHistoryServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\status_history\\v1\\status_history.swagger.json",
		},
		"mail_provider": {
			Handler: proto_mail_provider.RegisterMailProviderServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\mail_provider\\v1\\mail_provider.swagger.json",
		},
		"mail_status": {
			Handler: proto_mail_status.RegisterMailStatusServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\mail_status\\v1\\mail_status.swagger.json",
		},
		"mail_history": {
			Handler: proto_mail_history.RegisterMailHistoryServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\mail_history\\v1\\mail_history.swagger.json",
		},
		"type_mail": {
			Handler: proto_type_mail.RegisterTypeMailServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\type_mail\\v1\\type_mail.swagger.json",
		},
		"system_configuration": {
			Handler: proto_system_configuration.RegisterSystemConfigurationServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\system_configuration\\v1\\system_configuration.swagger.json",
		},
		"cost_tracking": {
			Handler: proto_cost_tracking.RegisterCostTrackingServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\cost_tracking\\v1\\cost_tracking.swagger.json",
		},
		"module": {
			Handler: proto_module.RegisterModuleServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\module\\v1\\module.swagger.json",
		},
		"module_child": {
			Handler: proto_module_child.RegisterModuleChildServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\module_child\\v1\\module_child.swagger.json",
		},
		"media": {
			Handler: proto_media.RegisterMediaServiceHandler,
			Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\media\\v1\\media.swagger.json",
		},
	}
}
