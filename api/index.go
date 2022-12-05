package handler

import (
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintln(w, "See API documentation at: "+os.Getenv("VERCEL_URL"))
}
