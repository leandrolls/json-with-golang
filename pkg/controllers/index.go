package controllers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is the index page!")
}
