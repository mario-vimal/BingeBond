package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Development ConfigDetails
	Test        ConfigDetails
	Production  ConfigDetails
}

type ConfigDetails struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Dialect  string
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	return db.Close()
}

func InitDb() {

	env := os.Getenv("APP_ENV")
	var config Config
	var currentConfig ConfigDetails

	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(color.RedString("Error reading file - config.json"), err)
		return
	}

	fmt.Println(color.GreenString("Reading file - config.json......."))

	json.Unmarshal([]byte(data), &config)

	if env == "production" {
		currentConfig = config.Production
	} else {
		currentConfig = config.Development
	}

	db, err = gorm.Open(currentConfig.Dialect,
		currentConfig.Username+":"+
			currentConfig.Password+"@tcp("+
			currentConfig.Host+":"+
			currentConfig.Port+")/"+
			currentConfig.Database+"?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println(color.RedString("MySQL connection Failed to Open "), err)
	} else {
		fmt.Println(color.GreenString("MySQL connection Established"))
	}

	db.DB().SetMaxOpenConns(1500)
	db.DB().SetMaxIdleConns(1500)

	for _, model := range []interface{}{
		//Add the struct of the Models to be migrated here
		UserDetail{}, Platform{},
	} {
		if err := db.AutoMigrate(model).Error; err != nil {
			fmt.Println(color.RedString("Migration failed")+" "+reflect.TypeOf(model).Name(), err)
		} else {
			fmt.Println(color.CyanString("Auto migrating " + reflect.TypeOf(model).Name() + "..."))
		}
	}

	//Add all the foreign key relations here
	db.Model(&Platform{}).AddForeignKey("user_id", "user_details(user_id)", "RESTRICT", "RESTRICT")
}
