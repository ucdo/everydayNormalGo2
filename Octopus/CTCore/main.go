package main

import (
	"Octopus/CTCore/structures"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var cnf *structures.MysqlCnf
var db *sqlx.DB

func main() {

}

func parseConfig() *structures.MysqlCnf {
	var config *structures.MysqlCnf
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("open file with err:", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("unmarshal json config with error:", err)
	}

	return config
}
