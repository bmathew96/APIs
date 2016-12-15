package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	env "personal-project/APIs/go-project/environment"
)

var config env.EnvConfig
var gerr error

var db *gorm.DB

func init() {

	log.Println("Grabbing Environment Variables")
	config, gerr = env.GetEnvConfig()
	if gerr != nil {
		log.Fatal(gerr)
		os.Exit(0)
	}

	log.Println("Connecting to database with ", config.DBUsername)
	connectionString := fmt.Sprintf("%s:%s@(%s:%d)/Common?parseTime=True&loc=%s", config.DBUsername, config.DBPassword, config.DBHostname, config.DBPort, "America%2FChicago")
	db, gerr = gorm.Open("mysql", connectionString)
	if gerr != nil {
		log.Fatal(gerr)
		os.Exit(0)
	}
}

func main() {
	defer db.Close()

	getHandlers()
	s := &http.Server{
		Addr:           ":8083",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Serving on port 8083")
	log.Fatal(s.ListenAndServe())
}
