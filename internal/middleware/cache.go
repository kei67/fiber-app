package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CacheMiddleware(duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Cache-Control", fmt.Sprintf("max-age=%d", int(duration.Seconds())))
		return c.Next()
	}
}
