package controller

import (
	"net/http"
	// "encoding/json"
	// "2k22go/views"
	// "fmt"
)

func RegisterApi() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/create", create())
	return mux
}