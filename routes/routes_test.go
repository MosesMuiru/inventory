package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SignIn struct {
	Username string
	Password string
	Email    string
}

func TestRoutes(t *testing.T) {
	a := ApisHandler{}
	s := SignInHandler{}
	l := LoginHandler{}
	user := SignIn{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "test",
	}
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

	t.Run("Sign up as a new user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v0/signin", structToReader(user))

		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)

		fmt.Println("---->response code", response.Code)

		var requestResponse SignIn
		if err := json.NewDecoder(response.Body).Decode(&requestResponse); err != nil {
			t.Fatalf("cannot decode response: %v", err)
		}

		want := user

		if requestResponse != want {
			t.Errorf("got %v, want %v", requestResponse, want)
		}
	})

	loginUserTest := LoginParams{
		Username: "moses",
		Password: "password",
	}
	t.Run("Login up a new User", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/v0/login", structToReader(loginUserTest))

		response := httptest.NewRecorder()
		l.ServeHTTP(response, request)
		var requestResponse LoginResponse

		if err := json.NewDecoder(request.Body).Decode(&requestResponse); err != nil {
			t.Fatalf("Cannot decode response: %v", err)
		}

		if response.Code != http.StatusOK {
			t.Fatalf("Buda check your error, %v", response.Code)
		}

	})
}

func structToReader(v any) io.Reader {
	data, err := json.Marshal(v)

	if err != nil {
		fmt.Println("This is wroking")
	}

	return bytes.NewReader(data)
}
