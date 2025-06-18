package settings

import (
	"Musica-Backend/internal/domain/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// リクエストが来るたびに「本当にログイン済みのユーザーか」をチェックして、問題なければそのユーザー情報をコンテキストに入れて次の処理へ渡す仕組み
func AuthMiddleware(usecase auth.GetCurrentUserUseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Authorizationヘッダーを取り出す
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				// ヘッダが無ければ401 Unauthorizedを返す
				return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
			}
			// "Bearerトークン文字列か検証"
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
			}
			// JWTをユースケースに渡して検証＆ユーザー情報を取得
			in := &auth.GetCurrentUserInput{
				Token: parts[1],
			}
			output, err := usecase.Execute(in)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
			// ユーザー情報をコンテキストにセット
			c.Set("currentUser", output.User)
			return next(c)
		}
	}
}
