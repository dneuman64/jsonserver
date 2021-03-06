package sqlToJson

import (
	"./../structs"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//returns *Rows from queried database
func GetRows(username string, password string, environment string, tableName string) *sql.Rows {
	//opens database
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * from asn")
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

//from given rows, returns array of byte arrays
func MakeJsonByteArray(rows *sql.Rows) [][]byte {
	byteArray := make([][]byte, 0)
	for rows.Next() {
		//currently hardcoded to "asn"
		var a int
		var b int
		var c int
		var d string
		err := rows.Scan(&a, &b, &c, &d)
		if err != nil {
			log.Fatal(err)
		}
		tableInstance := structs.AsnStruct{a, b, c, d}
		tableInstanceJson, _ := json.MarshalIndent(tableInstance, "", "  ")
		tableInstanceJson = append(tableInstanceJson, ","...)
		byteArray = append(byteArray, tableInstanceJson)
	}

	return byteArray
}
