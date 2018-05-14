package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	args := os.Args
	databaseUrl := args[1]
	con := databaseConnection(databaseUrl)
	defer con.Close()
	_, err := con.Exec("CREATE TABLE events(id INT AUTO_INCREMENT PRIMARY KEY, uuid VARCHAR(255) NOT NULL, payload TEXT NOT NULL, recorded_on VARCHAR(255) NOT NULL, type VARCHAR(255) NOT NULL)ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	println("Events table created.")

}

func databaseConnection(url string) *sql.DB {
	con, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return con
}
