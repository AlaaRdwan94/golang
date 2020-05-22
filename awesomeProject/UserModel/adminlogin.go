package UserModel

import (
	"awesomeProject.example.com/Structures"
	"fmt"
	"net/http"

	"time"
)
//--------------API for Admin login-----------------------
func LoginAdmin(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("req.URL.Method",req.Method)
	if err := req.ParseForm(); err != nil {
		return
	}
	ls := Structures.LoginStruct{}
	ls.UserName = req.PostFormValue("userame")
	ls.PassWord = req.PostFormValue("password")
	ls.Lastlogintime = time.Now().UTC()
	fmt.Println("ls",ls)
	switch req.Method {
	case "POST":
		ret , _ := GetAdmin(ls)

		switch ret {
		case 1 :
			fmt.Fprintf(w, "Sorry, only PUT and POST methods are supported.")
		case 2 :
			render(w, "templates/examples/Adminwelcome.html", ls)

		case 3 :
			fmt.Fprintf(w, "not found please reguister first")
		case 4 :
			fmt.Fprintf(w, "wrong password")
		default:
			fmt.Fprintf(w, "unknown error occurs")
		}

	default:
		fmt.Fprintf(w, "Sorry, only PUT and POST methods are supported.")
	}
}
//--------------API for Admin register --------------------
func RegisterAdmin(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("Register req.URL.Method",req.Method)
	if err := req.ParseForm(); err != nil {
		return
	}
	ls := Structures.LoginStruct{}
	ls.UserName = req.PostFormValue("userame")
	ls.PassWord = req.PostFormValue("password")
	ls.Lastlogintime = time.Now().UTC()
	fmt.Println("ls",ls)
	switch req.Method {
	case "POST":
		ret := AddAdmin(ls);
		switch ret {
		case 1 :
			fmt.Fprintf(w, "Sorry, only POST methods are supported.")
		case 2 :
			redirectionAdminPage(w,"Added sucessfully")
		case 3 :
			fmt.Fprintf(w, "User allready exists")
		default:
			fmt.Fprintf(w, "unknown error occurs")
		}

	default:
		fmt.Fprintf(w, "Sorry, only PUT and POST methods are supported.")
	}
}
//-------------- redirect to admin home page --------------
func redirectionAdminPage(w http.ResponseWriter, s string) {
	render(w, "templates/homeAdmin.html", nil)
}


