//url: https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mysql.go

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/andreslab/prj_api_crop/model"
)

func scanUser(s rowScanner) (*model.UserModel, error) {
	var (
		id       int64
		name     sql.NullString
		lastname sql.NullString
		email    sql.NullString
		pass     sql.NullString
		created  sql.NullString
	)
	if err := s.Scan(&id,
		&name,
		&lastname,
		&email,
		&pass,
		&created); err != nil {
		return nil, err
	}

	user := &model.UserModel{
		ID:       id,
		Name:     name.String,
		Lastname: lastname.String,
		Email:    email.String,
		Pass:     pass.String,
		Created:  created.String,
	}
	return user, nil
}

// newMySQLDB creates a new UserDatabase backed by a given MySQL server.
func NewMySQLDBUser() (*mysqlDB, error) {

	// Check database and table exists. If not, create it.
	err := ensureTableExists(tableNameUser, createTableStatementsUser)
	if err != nil {
		return nil, err
	}

	conn := ConnDB()
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	db := &mysqlDB{
		conn: conn,
	}
	// Prepared statements. The actual SQL queries are in the code near the
	// relevant method (e.g. addBook).
	if db.list, err = conn.Prepare(listStatementUser); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.get, err = conn.Prepare(getStatementUser); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatementUser); err != nil {
		return nil, fmt.Errorf("mysql: prepare insert: %v", err)
	}

	/*if db.listBy, err = conn.Prepare(listByStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare listBy: %v", err)
	}*/
	if db.update, err = conn.Prepare(updateStatementUser); err != nil {
		return nil, fmt.Errorf("mysql: prepare update: %v", err)
	}
	if db.delete, err = conn.Prepare(deleteStatementUser); err != nil {
		return nil, fmt.Errorf("mysql: prepare delete: %v", err)
	}

	return db, nil
}

func (db *mysqlDB) AddUser(u *model.UserModel) (id int64, err error) {

	r, err := execAffectingOneRow(
		db.insert,
		u.Name,
		u.Lastname,
		u.Email,
		u.Pass,
		u.Created)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	defer db.conn.Close()
	return lastInsertID, nil
}

func (db *mysqlDB) ListUser() ([]*model.UserModel, error) {

	rows, err := db.list.Query()
	if err != nil {
		fmt.Print("error")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var users []*model.UserModel
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		users = append(users, user)
	}
	defer db.conn.Close()
	return users, nil
}

func (db *mysqlDB) GetUser(id int64) {

}

func (db *mysqlDB) UpdateUser(b *model.UserModel) error {
	if b.ID == 0 {
		return errors.New("mysql: book with unassigned ID passed into updateBook")
	}

	_, err := execAffectingOneRow(db.update,
		b.Name,
		b.Lastname,
		b.Email,
		b.Pass,
		b.Created,
		b.ID)
	return err
}

func (db *mysqlDB) DeleteUser(id int64) error {
	if id == 0 {
		return errors.New("mysql: book with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}
