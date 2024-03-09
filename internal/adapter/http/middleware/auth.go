package middleware

import (
	"gocrud/internal/util"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	jwt *util.JWT
}

func NewAuthMiddleware(jwt *util.JWT) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

func (a *AuthMiddleware) Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		cookie := ctx.Cookies("token")
		if cookie == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized access")
		}

		claims, err := a.jwt.ValidateToken(cookie)
		if err != nil {
			return err
		}

		ctx.Locals("claims", claims)

		return ctx.Next()
	}
}
