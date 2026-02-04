package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
// import "hello"

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	
	fmt.Println("Starting go api server ....")

	fmt.Println(`
 ------	--------------------------------------
	|  Go  |    Learning Go Language    |
	 ------	------------------------------
	`)

	err:= http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
