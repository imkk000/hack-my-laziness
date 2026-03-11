package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v5"
)

func registerRoutes(e *echo.Echo, routes []RouteConfig) {
	for _, route := range routes {
		r := route // capture
		method := strings.ToUpper(r.Method)
		if method == "" {
			method = http.MethodGet
		}

		e.Add(method, r.Path, makeHandler(r))

		slog.Info("add path",
			"method", method,
			"path", r.Path,
		)
	}
}

func makeHandler(route RouteConfig) echo.HandlerFunc {
	return func(c *echo.Context) error {
		status := route.Response.Status
		if status == 0 {
			status = http.StatusOK
		}

		for k, v := range route.Response.Headers {
			c.Response().Header().Set(k, v)
		}

		// file takes priority over inline body
		if route.Response.File != "" {
			data, err := os.ReadFile(route.Response.File)
			if err != nil {
				return c.String(http.StatusInternalServerError,
					fmt.Sprintf("cannot read file: %s", route.Response.File))
			}

			contentType := detectContentType(route.Response.File, route.Response.Headers)
			return c.Blob(status, contentType, data)
		}

		body := interpolateParams(route.Response.Body, c)
		contentType := detectContentType("", route.Response.Headers)
		return c.Blob(status, contentType, []byte(body))
	}
}

func interpolateParams(body string, c *echo.Context) string {
	for _, name := range c.PathValues() {
		body = strings.ReplaceAll(body, ":"+name.Value, c.Param(name.Value))
	}
	return body
}

func detectContentType(filename string, headers map[string]string) string {
	if ct, ok := headers["Content-Type"]; ok {
		return ct
	}
	if strings.HasSuffix(filename, ".json") {
		return "application/json"
	}
	if strings.HasSuffix(filename, ".xml") {
		return "application/xml"
	}
	if strings.HasSuffix(filename, ".html") {
		return "text/html"
	}

	return "application/json"
}
