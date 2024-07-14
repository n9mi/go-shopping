package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TestProps map[string]interface{}

func newRequest(method string, url string, requestBody string) *http.Request {
	request := httptest.NewRequest(method, url, strings.NewReader(requestBody))
	request.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return request
}

func newRequestWithToken(method string, url string, requestBody string, token string) *http.Request {
	request := newRequest(method, url, requestBody)

	bearer := fmt.Sprintf("Bearer %s", token)
	request.Header.Add(fiber.HeaderAuthorization, bearer)

	return request
}
