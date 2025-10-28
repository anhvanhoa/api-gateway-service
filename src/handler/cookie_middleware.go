package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// CORS headers - cho phép credentials
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))

		// Nếu localhost (development), dùng SameSiteLaxMode
		// Nếu production (HTTPS), dùng SameSiteNoneMode
		if c.Request.Host == "localhost" || c.Request.Host == "127.0.0.1" {
			c.SetSameSite(http.SameSiteLaxMode)
		} else {
			c.SetSameSite(http.SameSiteNoneMode)
		}

		accessToken := GetAuthTokenFromCookie(c)
		fmt.Printf("accessToken: %s\n", accessToken)
		refreshToken := GetRefreshTokenFromCookie(c)
		fmt.Printf("refreshToken: %s\n", refreshToken)

		// Chỉ intercept auth endpoints
		if !strings.HasPrefix(c.Request.URL.Path, "/auth/") {
			c.Next()
			return
		}

		// Wrap response writer để capture response body
		w := &responseWriter{
			ResponseWriter: c.Writer,
			body:           make([]byte, 0),
		}
		c.Writer = w

		c.Next()

		// Parse response và set cookies ngay lập tức
		if w.Status() >= 200 && w.Status() < 300 && len(w.body) > 0 {
			handleAuthResponse(c, w, c.Request.URL.Path)
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func handleAuthResponse(c *gin.Context, w *responseWriter, path string) {
	switch path {
	case "/auth/register", "/auth/login", "/auth/refresh-token":
		setAuthCookies(c, w, path)
	case "/auth/logout":
		clearAuthCookies(c)
	}
}

func setAuthCookies(c *gin.Context, w *responseWriter, endpoint string) {
	var response map[string]any
	if err := json.Unmarshal(w.body, &response); err != nil {
		fmt.Printf("Error parsing JSON: %v, body: %s\n", err, string(w.body))
		return
	}

	switch endpoint {
	case "/auth/register":
		if token, ok := response["token"].(string); ok && token != "" {
			fmt.Printf("Setting auth_token: %s\n", token)
			c.SetCookie("auth_token", token, 24*60*60, "/", "", false, true)
		}
	case "/auth/login":
		if accessToken, ok := response["accessToken"].(string); ok && accessToken != "" {
			fmt.Printf("Setting access_token: %s\n", accessToken)
			c.SetCookie("access_token", accessToken, 60*60, "/", "", false, true)
		}
		if refreshToken, ok := response["refreshToken"].(string); ok && refreshToken != "" {
			fmt.Printf("Setting refresh_token: %s\n", refreshToken)
			c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
		}
	case "/auth/refresh-token":
		if accessToken, ok := response["accessToken"].(string); ok && accessToken != "" {
			fmt.Printf("Setting access_token (refresh): %s\n", accessToken)
			c.SetCookie("access_token", accessToken, 60*60, "/", "", false, true)
		}
		if refreshToken, ok := response["refreshToken"].(string); ok && refreshToken != "" {
			fmt.Printf("Setting refresh_token (refresh): %s\n", refreshToken)
			c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)
		}
	}

	fmt.Printf("Response headers: %v\n", c.Writer.Header())
}

func clearAuthCookies(c *gin.Context) {
	cookies := []string{"auth_token", "access_token", "refresh_token"}
	for _, cookieName := range cookies {
		c.SetCookie(cookieName, "", -1, "/", "", true, true)
	}
}

func GetAuthTokenFromCookie(c *gin.Context) string {
	token, _ := c.Cookie("access_token")
	if token == "" {
		token, _ = c.Cookie("auth_token")
	}
	return token
}

func GetRefreshTokenFromCookie(c *gin.Context) string {
	token, _ := c.Cookie("refresh_token")
	return token
}
