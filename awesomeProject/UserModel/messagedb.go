package UserModel
import "github.com/syndtr/goleveldb/leveldb"


var MDB *leveldb.DB
var MOpen = false

//------------------------------------------------------------------------------------------------------------
// create or open db if exist
//------------------------------------------------------------------------------------------------------------
func MsgOpendatabase() bool {
	if !MOpen {
		MOpen = true
		dbpath := "Database/Msgdb"
		var err error
		MDB, err = leveldb.OpenFile(dbpath, nil)
		if err != nil {
			return false
		}
		return true
	}
	return true
}