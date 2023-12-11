package model

import (
	"fmt"
	"log"
)

func Create() error {
	// Prepare the SQL statement.
	stmt, err := con.Prepare("INSERT INTO users(username, mail) VALUES($1, $2)")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the SQL statement.
	_, err = stmt.Exec("John Doe", "johndoe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully.")
	return nil
}
