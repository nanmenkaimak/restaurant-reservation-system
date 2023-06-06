package handlers

import (
	"net/http"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
