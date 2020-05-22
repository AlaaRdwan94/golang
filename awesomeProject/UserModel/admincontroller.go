
package UserModel

import (
"encoding/json"
"fmt"
	. "awesomeProject.example.com/Structures"
"time"
)

//add admin to database
func AddAdmin(ls LoginStruct) int {
	fmt.Println("AllAdmins")
	AllAdmins := GetAllAdminsName()
	fmt.Println(AllAdmins)
	for _, admin := range AllAdmins {
		if ls.UserName == admin {
			return 3
		}
	}
	fmt.Println(ls,"ls")
	As := AdminStruct{
		Admindata: UserStruct{
			ID:getCurrentAdminID(),
			Login : LoginStruct{
				UserName:      ls.UserName,
				PassWord:      ls.PassWord,
				Lastlogintime: ls.Lastlogintime,
			},
			RegisterTime : time.Now().UTC(),
			Active : true,
			Online : false,
		},
		Adminrule:true,

	}
	fmt.Println(As)
	fmt.Println("Us.Login", As.Admindata.Login)
	fmt.Println(As.Admindata.Login.Lastlogintime, "Us.Lastlogintime")
	value, _ := json.Marshal(As)
	OpenAdmindatabase()
	err := AdminDB.Put([]byte(As.Admindata.ID), []byte(value), nil)
	fmt.Println(GetAllAdminsName(),"GetAllUsersName")

	if err != nil {
		return 0
	}
	return 2
}
//get user from database
func GetAdmin(ls LoginStruct) (r int , d AdminStruct) {
	userlst := GetAllAdmins()

	for i, userData := range userlst {
		if(userData.Adminrule){
			if userData.Admindata.Login.UserName == ls.UserName {
				if userData.Admindata.Login.PassWord == ls.PassWord {
					return 2 , userData //registersuccessfully
				}
			}else if userData.Admindata.Login.UserName == ls.UserName && userData.Admindata.Login.PassWord != ls.PassWord{
				return 4 , d // errorpassworrd
			}else if userData.Admindata.Login.UserName != ls.UserName && i != len(userlst)-1{
				continue

			} else{
				return 3, d //user not found
			}

		}else {
			return 5, d //not admin
		}

	}
	return 0, d
}

//ID autoincrement
//TODO : used in AddUser to set the ID in database
func getCurrentAdminID() string {

	allIDs := GetAllAdminsID()
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
func GetAllAdmins() (values []AdminStruct) {
	OpenAdmindatabase()
	iter := AdminDB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata AdminStruct
		json.Unmarshal(value, &newdata)
		values = append(values, newdata)
	}
	return values
}
//get all IDs from database
func GetAllAdminsID() (ID []string) {
	OpenAdmindatabase()
	iter := AdminDB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata AdminStruct
		json.Unmarshal(value, &newdata)
		ID = append(ID, newdata.Admindata.ID)
	}
	return ID
}

//get all Names from database
func GetAllAdminsName() (Name []string) {
	OpenAdmindatabase()
	iter := AdminDB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata AdminStruct
		json.Unmarshal(value, &newdata)
		Name = append(Name, newdata.Admindata.Login.UserName)
	}
	return Name
}
