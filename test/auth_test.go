package test

import (
	"encoding/json"
	"fmt"
	"go-shopping/internal/model"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

var (
	registerUrl = "/api/v1.0/auth/register"
)

func TestRegister(t *testing.T) {
	testItems := map[string]TestProps{
		"auth_register_ok": {
			"request_name":     "aaaaa",
			"request_email":    "aaaaa@test.com",
			"request_password": "aaaaa",
			"expected_code":    http.StatusOK,
		},
		"auth_register_conflict": {
			"request_name":     "aaaaa",
			"request_email":    "customer1@test.com",
			"request_password": "aaaaa",
			"expected_code":    http.StatusConflict,
		},
		"auth_register_bad-request": {
			"request_name":     "",
			"request_email":    "a",
			"request_password": "a",
			"expected_code":    http.StatusBadRequest,
		},
	}

	for testName, testProps := range testItems {
		t.Run(testName, func(t *testing.T) {
			requestBody :=
				fmt.Sprintf(`{"name": "%s", "email": "%s", "password": "%s"}`,
					testProps["request_name"],
					testProps["request_email"],
					testProps["request_password"])
			request := newRequest(fiber.MethodPost, registerUrl, requestBody)

			response, err := app.Test(request)
			require.Nil(t, err)
			require.Equal(t, testProps["expected_code"], response.StatusCode)

			responseBody, err := io.ReadAll(response.Body)
			require.Nil(t, err)

			messageResponse := new(model.MessageResponse)
			err = json.Unmarshal(responseBody, messageResponse)
			require.Nil(t, err)

			require.True(t, len(messageResponse.Messages) > 0)
		})
	}
}
