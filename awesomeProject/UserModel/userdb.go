package UserModel

import "github.com/syndtr/goleveldb/leveldb"

// var dbpath = "Database/HeartBeatStruct"
var DB *leveldb.DB
var Open = false

//------------------------------------------------------------------------------------------------------------
// create or open db if exist
//------------------------------------------------------------------------------------------------------------
func Opendatabase() bool {
	if !Open {
		Open = true
		dbpath := "Database/UserDB"
		var err error
		DB, err = leveldb.OpenFile(dbpath, nil)
		if err != nil {
			return false
		}
		return true
	}
	return true
}