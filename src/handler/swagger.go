package handler

import (
	"api-gateway/src/bootstrap"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SwaggerHandler struct {
	env *bootstrap.Env
}

func NewSwaggerHandler(env *bootstrap.Env) *SwaggerHandler {
	return &SwaggerHandler{env: env}
}

func (s *SwaggerHandler) ServeSwaggerJSON(c *gin.Context, swaggerPath string) {
	swaggerData, err := os.ReadFile(swaggerPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read " + swaggerPath})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(swaggerData))
}

func (s *SwaggerHandler) ServeSwaggerUI(c *gin.Context, swaggerURL string) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui.css" />
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *, *:before, *:after {
            box-sizing: inherit;
        }
        body {
            margin:0;
            background: #fafafa;
        }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '/` + swaggerURL + `',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
