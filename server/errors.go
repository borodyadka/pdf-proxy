package server

import "net/http"

func response(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message))
	return err
}
