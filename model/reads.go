package model

import (
	"GoAPI/views"
)

func ReadAll() ([]views.PostRequest, error) {
	// this is the database query
	rows, err := con.Query("SELECT username, mail FROM users")
	if err != nil {
		return nil, err

	}
	// we create an Array of PostRequest using the struct we created in views
	users := []views.PostRequest{}
	// we loop through the rows.Next() and append the data to the users array
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Username, &data.Mail)
		users = append(users, data)
	}
	// we return the users array and nil for the error
	return users, nil
}
