package routes

import "net/http"

// dfine the different rout you want

type ApisHandler struct{}

func (a ApisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from moses"))
}

func Routers() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/api", ApisHandler{})

	return mux
}
