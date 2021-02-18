package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	log.Println(r.Body)
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	fmt.Fprintf(w, "Hello, World "+u.String())
}
