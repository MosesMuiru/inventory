package routes

import (
	"encoding/json"
	"net/http"
)

// dfine the different rout you want

type ApisHandler struct{}
type LoginHandler struct{}
type SignInHandler struct{}

type SignParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"Token"`
}

func (a ApisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from moses"))
}

func (l LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (s SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// unmarshal the incoming data
	sign := SignParams{}
	err := json.NewDecoder(r.Body).Decode(&sign)
	if err != nil {
		http.Error(w, "Invalid Params", http.StatusBadRequest)
	}

	// the response will be stored to the db

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := SignParams{
		Username: sign.Username,
		Password: sign.Password,
		Email:    sign.Email,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed", http.StatusInternalServerError)
	}

}

func Routers() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/api", ApisHandler{})
	mux.Handle("POST /api/v0/signin", SignInHandler{})
	mux.Handle("POST /api/v0/login", LoginHandler{})

	return mux
}
