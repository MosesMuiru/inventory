package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiHandler(t *testing.T) {
	t.Run("Moses from the api", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", nil)
		response := httptest.NewRecorder()

	})

}
