package singleton

import (
	"fmt"
	"sync"
)


var lock = &sync.Mutex{}

type Database struct {
}

var databaseInstance *Database

func GetDatabaseInstance() *Database {
	if databaseInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if databaseInstance == nil {
			fmt.Println("Creating Database Instance Now")
			databaseInstance = &Database{}
		} else {
			fmt.Println("Database Instance already created-1")
		}
	} else {
		fmt.Println("Database Instance already created-2")
	}
	return databaseInstance
}