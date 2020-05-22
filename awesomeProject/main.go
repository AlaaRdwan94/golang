package main

import (
	"awesomeProject.example.com/UserModel"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)


func handleRequests() {

	router := mux.NewRouter()
	//--------------html pages----------------
	router.Handle("/home", http.HandlerFunc(home)).Methods("GET", "OPTIONS")
	router.Handle("/register", http.HandlerFunc(register)).Methods("GET", "OPTIONS")
	router.Handle("/welcome", http.HandlerFunc(welcome)).Methods("GET", "OPTIONS")
	router.Handle("/homeAdmin", http.HandlerFunc(homeAdmin)).Methods("GET", "OPTIONS")
	router.Handle("/registerAdmin", http.HandlerFunc(registerAdmin)).Methods("GET", "OPTIONS")
	router.Handle("/Adminwelcome", http.HandlerFunc(Adminwelcome)).Methods("GET", "OPTIONS")

	//--------------go APIs--------------------
	router.Handle("/Login", http.HandlerFunc(UserModel.Login)).Methods("POST", "OPTIONS")
	router.Handle("/Register", http.HandlerFunc(UserModel.Register)).Methods("POST", "OPTIONS")
//------------------- Admin APIs -----------------
	router.Handle("/LoginAdmin", http.HandlerFunc(UserModel.LoginAdmin)).Methods("POST", "OPTIONS")
	router.Handle("/RegisterAdmin", http.HandlerFunc(UserModel.RegisterAdmin)).Methods("POST", "OPTIONS")
	router.Handle("/Socket", websocket.Handler(UserModel.Socket))
	router.Handle("/GetAllMessagesRooms", websocket.Handler(UserModel.GetAllMessagesRooms))
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("templates").HTTPBox()))
	//--------------server listening------------
	log.Println("Listening on port 3000 ...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.go.html", nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/register.html", nil)
}
func welcome(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/examples/welcome.html", nil)
}

func homeAdmin(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/homeAdmin.html", nil)
}
func registerAdmin(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/registerAdmin.html", nil)
}
func Adminwelcome(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/examples/Adminwelcome.html", nil)
}
//--------render html page to serve it in golang server---------
func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
