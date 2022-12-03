package handler

import (
	"fmt"
	"log"
	"net/http"
)

func PersonalBests(w http.ResponseWriter, r *http.Request) {
	log.Default().Println(r.Body)
	fmt.Fprintln(w, "testing")
}
