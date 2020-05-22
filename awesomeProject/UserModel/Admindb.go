package UserModel

import "github.com/syndtr/goleveldb/leveldb"

// var dbpath = "Database/HeartBeatStruct"
var AdminDB *leveldb.DB
var AdminOpen = false

//------------------------------------------------------------------------------------------------------------
// create or open db if exist
//------------------------------------------------------------------------------------------------------------
func OpenAdmindatabase() bool {
	if !AdminOpen {
		AdminOpen = true
		dbpath := "Database/AdminDB"
		var err error
		AdminDB, err = leveldb.OpenFile(dbpath, nil)
		if err != nil {
			return false
		}
		return true
	}
	return true
}