package UserModel

import (
	"html/template"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"awesomeProject.example.com/Structures"
	"time"
)

/*----------function to convert integer to fixed digits of string----------*/
func ConvertIntToFixedLengthString(key int, length int) (stringform string, err bool) {
	stringform = strconv.Itoa(key)
	stringlen := len(stringform)
	if stringlen > length {
		return "", false
	}
	for i := 0; i < length-stringlen; i++ {
		stringform = "0" + stringform
	}
	return stringform, true
}

/*----------function to convert Convert Fixed Length String to Int----------*/
func ConvertFixedLengthStringtoInt(key string) (stringform int) {
	for index := 0; index < len(key); index++ {
		if key[index:index+1] != "0" {
			number, _ := strconv.Atoi(key[index:len(key)])
			return number
		}

	}
	return 0
}
//add user to database
func AddUser(ls Structures.LoginStruct) int {
	Allusers := GetAllUsersName()
	for _, user := range Allusers {
		if ls.UserName == user {
			return 3
		}
	}
	fmt.Println(ls,"ls")
	Us := Structures.UserStruct{
		ID:getCurrentID(),
		Login :Structures.LoginStruct{
			UserName:      ls.UserName,
			PassWord:      ls.PassWord,
			Lastlogintime: ls.Lastlogintime,
		},
		RegisterTime : time.Now().UTC(),
		Active : true,
		Online : false,
	}
	fmt.Println(Us)
	fmt.Println("Us.Login", Us.Login)
	fmt.Println(Us.Login.Lastlogintime, "Us.Lastlogintime")
	value, _ := json.Marshal(Us)
	Opendatabase()
	err := DB.Put([]byte(Us.ID), []byte(value), nil)
	fmt.Println(GetAllUsersName(),"GetAllUsersName")

	if err != nil {
		return 0
	}
	return 2
}
//get user from database
func GetUser(ls Structures.LoginStruct) (r int , d Structures.UserStruct) {
	userlst := GetAllUsers()
	if userlst == nil{
		return 3, d //user not found
	}
	for i, userData := range userlst {
		if userData.Login.UserName == ls.UserName {
			if userData.Login.PassWord == ls.PassWord {
				return 2 , userData //registersuccessfully
			}
		}else if userData.Login.UserName == ls.UserName && userData.Login.PassWord != ls.PassWord{
			return 4 , d // errorpassworrd
		}else if userData.Login.UserName != ls.UserName && i != len(userlst)-1{
			continue

		} else{
			return 3, d //user not found
		}

	}
	return 0, d
}


//ID autoincrement
//TODO : used in AddUser to set the ID in database
func getCurrentID() string {

	allIDs := GetAllUsersID()
	var CurrentID string
	if allIDs == nil {
		CurrentID,_ = ConvertIntToFixedLengthString(0,20)
		return CurrentID
	}else{
		lastID := ConvertFixedLengthStringtoInt(allIDs[len(allIDs)-1])+1
		CurrentID,_ = ConvertIntToFixedLengthString(lastID,20)
	}
	return CurrentID
}
//get all users from database
func GetAllUsers() (values []Structures.UserStruct) {
	Opendatabase()
	iter := DB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata Structures.UserStruct
		json.Unmarshal(value, &newdata)
		values = append(values, newdata)
	}
	return values
}
//get all Users IDs from database
func GetAllUsersID() (ID []string) {
	Opendatabase()
	iter := DB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata Structures.UserStruct
		json.Unmarshal(value, &newdata)
		ID = append(ID, newdata.ID)
	}
	return ID
}

//get all Users Names from database
func GetAllUsersName() (Name []string) {
	Opendatabase()
	iter := DB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata Structures.UserStruct
		json.Unmarshal(value, &newdata)
		Name = append(Name, newdata.Login.UserName)
	}
	return Name
}
//redirection to client home page
func redirectionPage(w http.ResponseWriter, s string) {
	render(w, "templates/home.go.html", nil)
}
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
