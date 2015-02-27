package main

import (
	"fmt"
	"net/http"
	_ "encoding/json"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/sessions"
)

const (

)

func GetDates(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "")
}

func GetIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Nothing to see here")
}