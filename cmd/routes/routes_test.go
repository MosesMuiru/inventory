package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	a := ApisHandler{}
	t.Run("Should test the tests that are there", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api", nil)
		response := httptest.NewRecorder()
		a.ServeHTTP(response, request)

		got := response.Body.String()

		fmt.Println("Got", got)

		fmt.Println("working")
		want := "hello from moses"
		if got != want {
			t.Errorf("got %v, wan %v", got, want)
		}
	})

}
