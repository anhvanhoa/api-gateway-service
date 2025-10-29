package constants

import "time"

const (
	AccessTokenCookieName  = "at"
	RefreshTokenCookieName = "rt"
)

const (
	AccessTokenExpires  = 6 * time.Hour
	RefreshTokenExpires = 15 * 24 * time.Hour
)
