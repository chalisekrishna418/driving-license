package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var confFile = "config/db.json"

type mongoConfig struct {
	Host     string `json:"hostname" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Database string `json:"database" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *mgo.Database

// MongoInit initialize mongodb
func MongoInit() {

	dbconf := mongoConfig{}
	jsonFile, err := os.Open(confFile)
	mongoConf, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(mongoConf)
	json.Unmarshal(mongoConf, &dbconf)

	fmt.Println("data:")
	fmt.Println(dbconf.Host)
	fmt.Println(dbconf.Port)
	fmt.Println(dbconf.Database)

	conn, err := mgo.Dial("mongodb://" + dbconf.Host + ":" + dbconf.Port)
	if err != nil {
		fmt.Print("db connection error")
		fmt.Printf("Error: %s", err)
	}
	fmt.Print("db connection successful")
	db = conn.DB(dbconf.Database)
}

// Insert inserts data to a colletion
func Insert(data interface{}, collection string) error {
	MongoInit()
	err := db.C(collection).Insert(data)
	return err
}

// List lists data from a collection
func List(data *interface{}, filter bson.M, collection string) error {
	MongoInit()
	err := db.C(collection).Find(filter).All(&data)
	return err
}

// Update updates data in a collection
func Update(data interface{}, newData interface{}, collection string) error {
	MongoInit()
	err := db.C(collection).Update(data, newData)
	return err
}

// Select selects data from a collection
func Select(data *interface{}, filter bson.M, collection string) error {
	MongoInit()
	err := db.C(collection).Find(filter).Limit(1).All(&data)
	return err
}

// Delete deletes data from a collection
func Delete(data *interface{}, collection string) error {
	MongoInit()
	err := db.C(collection).Remove(data)
	return err
}
