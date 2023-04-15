package appmetrics

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitPrometheus(e *echo.Echo) {
	requestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Number of requests.",
	}, []string{"method", "path", "status"})
	// Register the counter with Prometheus's default registry.
	prometheus.MustRegister(requestsTotal)
	// Use a middleware to count requests and update the Prometheus counter
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Call the next handler
			err := next(c)

			// Increment the counter based on the response status code and path
			statusCode := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path
			requestsTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()

			return err
		}
	})

	// Set up a route to expose the metrics to Prometheus
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
