package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing/teplates/templates/pages"
)

func test_templates() {
  err := pages.Home(os.Stdout, pages.HomeProps{
    Header: "This will be a todo list",
  })
  if err != nil {
    log.Panic(err)
  }
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    err := pages.Home(w, pages.HomeProps{
      Header: "This will be a todo list",
    })
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(fmt.Sprintf("500 - Internal server error\n%v", err)))
    }
  })

  http.ListenAndServe("127.0.0.1:8080", nil)
}
