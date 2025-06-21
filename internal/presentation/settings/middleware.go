package settings

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewJWTMiddleware() (echo.MiddlewareFunc, error) {
	// Create a new JWT middleware with the default options
	options := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			// Handle errors during key refresh
			// You can log the error or take other actions as needed
			fmt.Println("Error refreshing keys:", err)
		},
		RefreshInterval:   time.Hour,
		RefreshUnknownKID: true, // You can set a rate limiter if needed 今回未知の kid であれば即時リフレッシュ
	}
	jwks, err := keyfunc.Get(os.Getenv("AUTH0_JWKS_URL"), options)
	if err != nil {
		return nil, fmt.Errorf("failed to get JWKS: %w", err)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//Authorizationヘッダ所得
			authz := c.Request().Header.Get("Authorization")
			if authz == "" {
				return echo.NewHTTPError(401, "Authorization header is required")
			}
			parts := strings.SplitN(authz, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(401, "Authorization header must be Bearer token")
			}
			token := parts[1]
			//Parse and Verify
			payload, err := jwt.Parse(token, jwks.Keyfunc)
			if err != nil {
				return echo.NewHTTPError(401, fmt.Sprintf("Invalid token: %v", err))
			}
			claims, ok := payload.Claims.(jwt.MapClaims)
			if !ok || !payload.Valid {
				return echo.NewHTTPError(401, "Invalid token claims")
			}
			//Check issuer
			if claims["iss"] != os.Getenv("AUTH0_ISSUER") {
				return echo.NewHTTPError(401, "Invalid token issuer")
			}
			//check audience
			audience, ok := claims["aud"].([]interface{})
			if !ok || len(audience) == 0 {
				return echo.NewHTTPError(401, "Invalid token audience")
			}
			if audience[0] != os.Getenv("AUTH0_AUDIENCE") {
				return echo.NewHTTPError(401, "Invalid token audience")
			}
			//Extract email
			email, ok := claims["email"].(string)
			if !ok || email == "" {
				return echo.NewHTTPError(401, "Failed to extract email from token")
			}
			// Set user email in context
			c.Set("user_email", email)
			// Call the next handler
			return next(c)
		}
	}, nil
}
