package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var dbConfStr string

func Init() {

	dbConfFile, err := os.Open("db.json")
	if err != nil {
		panic(err)
	}

	dbConfByte, _ := ioutil.ReadAll(dbConfFile)
	dbConf := map[string]string{}
	json.Unmarshal(dbConfByte, &dbConf)
	fmt.Print(dbConf)
	dbConfStr = dbConf["user"] + ":" + dbConf["pass"] + "@tcp(" + dbConf["path"] + ")/gam?charset=utf8mb4&parseTime=True"

}

func DBConn() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dbConfStr), &gorm.Config{})
}
