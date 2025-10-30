package handler

import (
	"api-gateway/src/bootstrap"
	"api-gateway/src/constants"
	"context"
	"log"
	"net/http"

	proto_auth "github.com/anhvanhoa/sf-proto/gen/auth/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleAuth(
	ctx context.Context,
	env *bootstrap.Env,
	c *gin.Engine,
	swaggerHandler *SwaggerHandler,
	service ServiceHandler,
) {
	conn, err := grpc.NewClient(
		env.AuthService.Host+":"+env.AuthService.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("Không thể kết nối đến gRPC server %s: %v", "auth", err)
		return
	}
	authClient := proto_auth.NewAuthServiceClient(conn)

	c.POST(env.AuthService.Route+"/login", func(c *gin.Context) {
		var loginRequest proto_auth.LoginRequest
		loginRequest.Os = c.GetHeader("user-agent")
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		loginResponse, err := authClient.Login(ctx, &loginRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		expiresAccessToken := int(constants.AccessTokenExpires.Seconds())
		expiresRefreshToken := int(constants.RefreshTokenExpires.Seconds())
		isSecure := env.IsProduction()
		setCookie(c, constants.AccessTokenCookieName, loginResponse.AccessToken, env.Domain, "/", isSecure, true, expiresAccessToken)
		setCookie(c, constants.RefreshTokenCookieName, loginResponse.RefreshToken, env.Domain, "/", isSecure, true, expiresRefreshToken)
		c.JSON(http.StatusOK, loginResponse)
	})

	// Register endpoint
	c.POST(env.AuthService.Route+"/register", func(c *gin.Context) {
		var registerRequest proto_auth.RegisterRequest
		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		registerResponse, err := authClient.Register(ctx, &registerRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusCreated, registerResponse)
	})

	// Refresh token endpoint
	c.POST(env.AuthService.Route+"/refresh", func(c *gin.Context) {
		refreshToken := getCookie(c, constants.RefreshTokenCookieName)
		if refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Vui lòng đăng nhập"})
			return
		}
		refreshRequest := proto_auth.RefreshTokenRequest{
			RefreshToken: refreshToken,
			Os:           c.GetHeader("user-agent"),
		}
		refreshResponse, err := authClient.RefreshToken(ctx, &refreshRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		expiresAccessToken := int(constants.AccessTokenExpires.Seconds())
		expiresRefreshToken := int(constants.RefreshTokenExpires.Seconds())
		isSecure := env.IsProduction()
		setCookie(c, constants.AccessTokenCookieName, refreshResponse.AccessToken, env.Domain, "/", isSecure, true, expiresAccessToken)
		setCookie(c, constants.RefreshTokenCookieName, refreshResponse.RefreshToken, env.Domain, "/", isSecure, true, expiresRefreshToken)
		c.JSON(http.StatusOK, refreshResponse)
	})

	// Logout endpoint
	c.POST(env.AuthService.Route+"/logout", func(c *gin.Context) {
		logoutRequest := proto_auth.LogoutRequest{
			RefreshToken: getCookie(c, constants.RefreshTokenCookieName),
			AccessToken:  getCookie(c, constants.AccessTokenCookieName),
		}
		if err := c.ShouldBindJSON(&logoutRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		logoutResponse, err := authClient.Logout(ctx, &logoutRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		clearCookie(c, constants.AccessTokenCookieName, env.Domain, "/")
		clearCookie(c, constants.RefreshTokenCookieName, env.Domain, "/")
		c.JSON(http.StatusOK, logoutResponse)
	})

	// Forgot password endpoint
	c.POST(env.AuthService.Route+"/forgot-password", func(c *gin.Context) {
		var forgotPasswordRequest proto_auth.ForgotPasswordRequest
		forgotPasswordRequest.Os = c.GetHeader("user-agent")
		if err := c.ShouldBindJSON(&forgotPasswordRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		forgotPasswordResponse, err := authClient.ForgotPassword(ctx, &forgotPasswordRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, forgotPasswordResponse)
	})

	// Check Token endpoint
	c.GET(env.AuthService.Route+"/check-token/:token", func(c *gin.Context) {
		token := c.Param("token")
		checkTokenResponse, err := authClient.CheckToken(ctx, &proto_auth.CheckTokenRequest{
			Token: token,
		})
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, checkTokenResponse)
	})

	// Reset Password endpoint
	c.POST(env.AuthService.Route+"/verify-account", func(c *gin.Context) {
		var verifyAccountRequest proto_auth.VerifyAccountRequest
		if err := c.ShouldBindJSON(&verifyAccountRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		verifyAccountResponse, err := authClient.VerifyAccount(ctx, &verifyAccountRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, verifyAccountResponse)
	})

	// Reset Password by Code endpoint
	c.POST(env.AuthService.Route+"/reset-password-by-code", func(c *gin.Context) {
		var passwordByCodeRequest proto_auth.ResetPasswordByCodeRequest
		if err := c.ShouldBindJSON(&passwordByCodeRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		passwordByCodeResponse, err := authClient.ResetPasswordByCode(ctx, &passwordByCodeRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, passwordByCodeResponse)
	})

	// Reset Password by Code endpoint
	c.POST(env.AuthService.Route+"/reset-password-by-token", func(c *gin.Context) {
		var passwordByTokenRequest proto_auth.ResetPasswordByTokenRequest
		if err := c.ShouldBindJSON(&passwordByTokenRequest); err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		passwordByTokenResponse, err := authClient.ResetPasswordByToken(ctx, &passwordByTokenRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, passwordByTokenResponse)
	})

	// Check Code endpoint
	c.GET(env.AuthService.Route+"/check-code/:code/:email", func(c *gin.Context) {
		code := c.Param("code")
		email := c.Param("email")
		checkCodeRequest := proto_auth.CheckCodeRequest{
			Code:  code,
			Email: email,
		}
		checkCodeResponse, err := authClient.CheckCode(ctx, &checkCodeRequest)
		if err != nil {
			RespondWithGrpcError(c, err)
			return
		}
		c.JSON(http.StatusOK, checkCodeResponse)
	})

	fileJSON := env.AuthService.Route + ".json"
	c.GET(fileJSON, func(c *gin.Context) {
		swaggerHandler.ServeSwaggerJSON(c, service.Swagger)
	})
	c.GET("/swagger"+env.AuthService.Route, func(c *gin.Context) {
		swaggerHandler.ServeSwaggerUI(c, fileJSON)
	})
}

func setCookie(c *gin.Context, name, value, domain, path string, secure bool, httpOnly bool, maxAge int) {
	c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func clearCookie(c *gin.Context, name, domain, path string) {
	c.SetCookie(name, "", -1, path, domain, false, true)
}

func getCookie(c *gin.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie
}
