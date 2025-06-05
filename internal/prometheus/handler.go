package prometheus

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewMetricsHandler はPrometheusメトリクス用のFiberハンドラーを返します
func NewMetricsHandler() fiber.Handler {
	handler := promhttp.Handler()

	return adaptor.HTTPHandler(handler)
}
