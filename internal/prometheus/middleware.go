package prometheus

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTPリクエスト数のカウンター
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	// HTTPリクエスト処理時間のヒストグラム
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// アクティブな接続数のゲージ
	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_active_connections",
			Help: "Number of active HTTP connections",
		},
	)
)

func init() {
	// Prometheusにメトリクスを登録
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(activeConnections)
}

// Prometheusミドルウェア
func PrometheusMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// アクティブ接続数を増加
		activeConnections.Inc()
		defer activeConnections.Dec()

		// リクエストを処理
		err := c.Next()

		// メトリクスを記録
		duration := time.Since(start).Seconds()
		method := c.Method()
		path := c.Route().Path
		statusCode := strconv.Itoa(c.Response().StatusCode())

		// リクエスト数をカウント
		httpRequestsTotal.WithLabelValues(method, path, statusCode).Inc()

		// 処理時間を記録
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)

		return err
	}
}
