package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

)
type Movie struct{
	//We need to convert all the fields into JSON
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director  struct{
	FirstName string `json:"firstname"`
	LastName string  `json:"lastname"`

}

var movies[] Movie

func main(){
	r :=mux.NewRouter()
	r.HandleFunc()


}