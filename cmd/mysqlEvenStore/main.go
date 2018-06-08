package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database := flag.String("database", "", "-database [DATABASE_NAME]")
	host := flag.String("host", "", "-host [HOST]")
	port := flag.String("port", "", "-port [PORT]")
	user := flag.String("user", "", "-user [USER]")
	pass := flag.String("pass", "", "-pass [PASSWORD]")
	flag.Parse()
	if "" == *database || "" == *host || "" == *port || "" == *user || "" == *pass{
		fmt.Println("Usage : -host [HOST] -port [PORT] -user [USER] -pass [PASSWORD] -database [DATABASE_NAME]")
		return
	}

	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *user, *pass, *host, *port, *database)

	con := databaseConnection(databaseUrl)
	defer con.Close()
	_, err := con.Exec("CREATE TABLE events(id INT AUTO_INCREMENT PRIMARY KEY, uuid VARCHAR(255) NOT NULL, payload TEXT NOT NULL, recorded_on VARCHAR(255) NOT NULL, type VARCHAR(255) NOT NULL)ENGINE = InnoDB;")

	if err != nil {
		panic(err)
	}

	println("Mysql event store table created.")

}

func databaseConnection(url string) *sql.DB {
	con, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return con
}
