package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AtccData :
type AtccData struct {
	Id                  int    `json:"id" binding:"requied"`
	AtccId              int    `json:"atcc_id" binding:"requied"`
	Ip                  string `json:"utf8mb4_general_ci"  binding:"requied"`
	DateTime            string `json:"date_time"  binding:"requied"`
	TotalCount          int    `json:"total_count"  binding:"requied"`
	Lane                string `json:"lane"  binding:"requied"`
	CarCount            int    `json:"car_count" binding:"requied"`
	TruckCount          int    `json:"truck_count" binding:"requied"`
	MavCount            int    `json:"mac_count"  binding:"requied"`
	BikeCount           int    `json:"bike_count" binding:"requied"`
	BusCount            int    `json:"bus_count" binding:"requied"`
	LcvCount            int    `json:"lcv_count" binding:"requied"`
	AutoCount           int    `json:"auto_count" binding:"requied"`
	TractorCount        int    `json:"tractor_count" binding:"requied"`
	TruckAxle2Count     int    `json:"truck_axle2_count" binding:"requied"`
	TruckAxle3Count     int    `json:"truck_axle3_count" binding:"requied"`
	TruckAxle4Count     int    `json:"truck_axle4_count" binding:"requied"`
	TruckAxle5Count     int    `json:"truck_axle5_count" binding:"requied"`
	TruckAxle6Count     int    `json:"truck_axle6_count" binding:"requied"`
	TruckAxle7PlusCount int    `json:"truck_axle7_count" binding:"requied"`
	DateAdded           string `json:"date_added" binding:"requied"`
}

var DB *gorm.DB

func main() {
	fmt.Println("enter in main ....")

	var err error

	// conntion with database
	DB, _, err = databaseConnection()
	if err != nil {
		return
	}

	router := gin.Default()
	// http://localhost:8080/atcc/atcc-data
	router.GET("/atcc/atcc-data", GetAtccData)

	router.Run("localhost:8080")

}

/*func append() {


	f, err := os.OpenFile(AtccData, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	src, err := os.Open("AtccData")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.OpenFile("atcc", "AtccData")
	if err != nil {
		panic(err)

	}
	defer dst.Close()
	data, err := ioutil.ReadAll(src)
	if err != nil {
		panic(err)
	}
	_, err = dst.Write(data)
	if err != nil {
		panic(err)
	}
}*/

func GetAtccData(c *gin.Context) {

	var atcc_data []AtccData

	//var atcc AtccData

	//atccData = append(atccData, atcc)
	err := DB.Find(&atcc_data).Error //for insert data in local host
	if err != nil {                  //
		fmt.Println("fetch err:", err) //
		return                         //
	}

	c.JSON(http.StatusOK, gin.H{"message": "this is for atcc data", "data": atcc_data})

}

func databaseConnection() (*gorm.DB, string, error) {
	username := "root"
	password := "root"
	host := "localhost"
	database := "itms"

	dsn := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return db, database, err
	}

	fmt.Println("database connected successfully...")
	return db, database, nil
}

/*
var AtccData atcc_data
var db *gorm.DB
var err error

func CreateTable() {
	query := `CreateTable atcc_data info(
	Id INT(11) NOT NULL AUTO_INCREMENT,
	AtccId INT(11) NOT NULL DEFAULT '0',
	Ip VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	DateTime DATETIME NOT NULL,
	TotalCount INT(11) NOT NULL DEFAULT '0',
	Lane VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	CarCount INT(11) NOT NULL DEFAULT '0',
	TruckCount INT(11) NOT NULL DEFAULT '0',
	MavCount INT(11) NOT NULL DEFAULT '0',
	BikeCount INT(11) NOT NULL DEFAULT '0',
	BusCount INT(11) NOT NULL DEFAULT '0',
	LcvCount INT(11) NOT NULL DEFAULT '0',
	AutoCount INT(11) NOT NULL DEFAULT '0',
	TractorCount INT(11) NOT NULL DEFAULT '0',
	TruckAxle2Count INT(11) NOT NULL DEFAULT '0',
	TruckAxle3Count INT(11) NOT NULL DEFAULT '0',
	TruckAxle4Count INT(11) NOT NULL DEFAULT '0',
	TruckAxle5Count INT(11) NOT NULL DEFAULT '0',
	TruckAxle6Count INT(11) NOT NULL DEFAULT '0',
	TruckAxle7PlusCount INT(11) NOT NULL DEFAULT '0',
	AateAdded DATETIME NULL DEFAULT NULL,
	PRIMARY KEY (Id) USING BTREE,
	UNIQUE INDEX AtccIdDateTimeLane (atcc_id, date_time, lane) USING BTREE,
	INDEX date_time (date_time) USING BTREE
	COLLATE = Utf8mb4Generalci,

	ENGINE = InnENGINE
	AUTO_INCREMENT = 197);`
	_, err := err, db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("atcc_data Table created...")

}

func getAllatcc_data(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getAllatcc_data)
}

/*
	func getatcc_data(db *gorm.DB, atcc_data *atcc_data) (err error) {
		err = db.Find(atcc_data).Error
		if err != nil {
			return err
		}

}

func getatcc_data(w http.ResponseWriter, r *http.Request) {
	vars := gin.Vars(r)
	key := vars["id"]
	// fror _,atcc_data:= _,range() {
		for _, atcc_data := range
		if atcc_data.Id == key{
			json.NewEncoder(w).Encode(atcc_data)
			}
	}

}
func createatcc_data(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var AtccData atcc_data
    json.Unmarshal(reqBody, &atcc_data)
    Employees = append(AtccData, atcc_data)
    json.NewEncoder(w).Encode(atcc_data)
}




func main() {
	databaseConnection()
	CreateTable()
	getatcc_data()
	getAllatcc_data()

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc")
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", h1)
}
*/
