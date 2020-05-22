package UserModel

import (

	"encoding/json"
	"fmt"
	"net/http"
	"awesomeProject.example.com/Structures"

	"time"
)

func Login(w http.ResponseWriter, req *http.Request)  {
	fmt.Println("req.URL.Method",req.Method)

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	ls := Structures.LoginStruct{}
	ls.UserName = req.PostFormValue("userame")
	ls.PassWord = req.PostFormValue("password")
	ls.Lastlogintime = time.Now().UTC()
	fmt.Println("ls",ls)
	switch req.Method {
	case "POST":
		ret , users := GetUser(ls)
		fmt.Println("ret",ret)
		switch ret {
		case 1 :
			fmt.Fprintf(w, "Sorry, only PUT and POST methods are supported.")
		case 2 :
			ls2 := &users
			Users, err := json.Marshal(ls2)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(Users)
			//w.Write(Users)
			render(w, "templates/examples/welcome.html", ls)
			//fmt.Fprintf(w, string(Users))

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

func Register(w http.ResponseWriter, req *http.Request)  {
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
		ret := AddUser(ls);
		switch ret {
		case 1 :
			fmt.Fprintf(w, "Sorry, only POST methods are supported.")
		case 2 :
			//fmt.Fprintf(w, "Added sucessfully")
			redirectionPage(w,"Added sucessfully")
		case 3 :
			fmt.Fprintf(w, "User allready exists")
		default:
			fmt.Fprintf(w, "unknown error occurs")
		}

		//http.ServeFile(w, req, "form.html")

	default:
		fmt.Fprintf(w, "Sorry, only PUT and POST methods are supported.")
	}
}

