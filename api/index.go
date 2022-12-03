package handler

import (
	"fmt"
	"net/http"
	"os"
)

const API_BASE_URL = "https://www.speedrun.com/api/v1"
const AUTH_HEADER = "X-API-Key"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintln(w, "See API documentation at: "+os.Getenv("VERCEL_URL"))
}
