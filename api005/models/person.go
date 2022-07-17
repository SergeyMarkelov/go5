package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./people_db.db") // ./
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Person struct {
	Login    string `json:"Login"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

func CheckPersonByLoginAndPassword(Login string, Password string) Person {

	rows := DB.QueryRow("SELECT Login, Password from people WHERE Login = ? AND Password = ?", Login, Password)

	type tempcolumns struct {
		Login    string
		Password string
	}
	c := tempcolumns{}
	login := rows.Scan(&c.Login)
	password := rows.Scan(c.Password)

	if Login == login.Error() && Password == password.Error() {
		Person := Person{"Login", "is", "Ok"}
		return Person
	} else {
		return Person{"Isn't okay", "Isn't Okay", "Isn't okay"}
	}
}

/*func GetPersons(count int) ([]Person, error) {

	rows, err := DB.Query("SELECT Login, Name, Password from people LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]Person, 0)

	for rows.Next() {
		singlePerson := Person{}
		err = rows.Scan(&singlePerson.Login, &singlePerson.Name, &singlePerson.Password)
		if err != nil {
			return nil, err
		}
		people = append(people, singlePerson)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return people, err
}

func Registration(newPerson Person) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO people (Login, Name, Password) VALUES (?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newPerson.Login, newPerson.Name, newPerson.Password)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
*/
