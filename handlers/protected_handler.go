package handlers

import "net/http"

func Protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected data accessed"))
}