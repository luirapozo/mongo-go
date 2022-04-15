package main

import (
	"fmt"
	"log"
	"mongo-go/models"
	"mongo-go/repositories"
	"net/http"
	"time"
)

var currentUser models.User

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", createUHandler)
	http.HandleFunc("/signin", verifyUHandler)
	http.HandleFunc("/add", testHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func createUHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Succesfull")
	user := models.User{
		Name:        r.FormValue("name"),
		UserName:    r.FormValue("username"),
		Password:    r.FormValue("password"),
		Rank:        0,
		Time_create: time.Now(),
		Characters:  make(models.Megucas, 0, 50),
	}
	err := repositories.CreateU(user)
	if err != nil {
		log.Fatal(err)
	}
	currentUser = repositories.FindUserByUName(r.FormValue("username"), r.FormValue("password"))

}

func verifyUHandler(w http.ResponseWriter, r *http.Request) {
	currentUser = repositories.FindUserByUName(r.FormValue("username"), r.FormValue("password"))
	fmt.Printf("%v\n", currentUser.Id)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	meguca := repositories.FindMegucaByName(r.FormValue("meguca"))
	currentUser.Characters = append(currentUser.Characters, &meguca)
	repositories.UpdateU(currentUser, currentUser.Id)
}
