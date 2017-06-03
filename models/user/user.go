package user

import (
	"database/sql"
	"log"
)

// User -- represents database model
type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	password  string
}

func Find(db *sql.DB) ([]User, error) {
	query, err := db.Prepare(`
SELECT id, firstName, lastName, email, password
FROM user;
	`)
	rows, err := query.Query()
	defer rows.Close()

	if err != nil {
		log.Fatal("db error: " + err.Error())
		return nil, err
	}

	var results []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.id, &user.firstName, &user.lastName, &user.email, &user.password); err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println(results)

	return results, nil
}
