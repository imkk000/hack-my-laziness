package http

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

type RequestOptions struct {
	URL         string
	Headers     []string
	Body        string
	BodyFile    string
	ContentType string
}

var contentTypeMap = map[string]string{
	"json": "application/json",
	"form": "application/x-www-form-urlencoded",
	"xml":  "application/xml",
	"text": "text/plain",
}

func doRequest(method string, opts RequestOptions) error {
	client := resty.New()
	req := client.R()

	// Set headers
	for _, h := range opts.Headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid header format %q, expected Key: Value", h)
		}
		req.SetHeader(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	// Resolve content type
	ct := ""
	if opts.ContentType != "" {
		mapped, ok := contentTypeMap[strings.ToLower(opts.ContentType)]
		if !ok {
			return fmt.Errorf("unsupported content type %q, supported: json, form, xml, text", opts.ContentType)
		}
		ct = mapped
		req.SetHeader("Content-Type", ct)
	}

	// Set body
	if opts.BodyFile != "" {
		data, err := os.ReadFile(opts.BodyFile)
		if err != nil {
			return fmt.Errorf("read body file: %w", err)
		}
		req.SetBody(data)
	} else if opts.Body != "" {
		req.SetBody(opts.Body)
	}

	resp, err := req.Execute(method, opts.URL)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	printResponse(resp, ct)
	return nil
}

func printResponse(resp *resty.Response, contentType string) {
	fmt.Printf("Status: %s\n", resp.Status())
	fmt.Println("Headers:")
	for k, v := range resp.Header() {
		fmt.Printf("  %s: %s\n", k, strings.Join(v, ", "))
	}
	fmt.Println()

	body := resp.Body()
	respCT := resp.Header().Get("Content-Type")

	if strings.Contains(respCT, "application/json") || strings.Contains(contentType, "json") {
		var pretty any
		if err := json.Unmarshal(body, &pretty); err == nil {
			out, _ := json.MarshalIndent(pretty, "", "  ")
			fmt.Println(string(out))
			return
		}
	}

	fmt.Println(string(body))
}
