package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/andreslab/prj_api_crop/model"
)

func scanInfo(s rowScanner) (*model.InfoModel, error) {

	var (
		id      int64
		created sql.NullString
		crop    sql.NullString
		sowing  sql.NullString
		city    sql.NullString
		address sql.NullString
		iduser  int64
	)
	if err := s.Scan(&id,
		&created,
		&crop,
		&sowing,
		&city,
		&address,
		&iduser); err != nil {
		return nil, err
	}

	info := &model.InfoModel{
		ID:         id,
		Created:    created.String,
		Crop:       crop.String,
		SowingType: sowing.String,
		City:       city.String,
		Address:    address.String,
		IDUser:     iduser,
	}
	return info, nil
}

// newMySQLDB creates a new UserDatabase backed by a given MySQL server.
func NewMySQLDBInfo() (*mysqlDB, error) {

	// Check database and table exists. If not, create it.
	err := ensureTableExists(tableNameInfo, createTableStatementsInfo)
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
	if db.list, err = conn.Prepare(listStatementInfo); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.get, err = conn.Prepare(getStatementInfo); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatementInfo); err != nil {
		return nil, fmt.Errorf("mysql: prepare insert: %v", err)
	}
	/*
		if db.listBy, err = conn.Prepare(listByStatement); err != nil {
			return nil, fmt.Errorf("mysql: prepare listBy: %v", err)
		}
		if db.update, err = conn.Prepare(updateStatement); err != nil {
			return nil, fmt.Errorf("mysql: prepare update: %v", err)
		}
		if db.delete, err = conn.Prepare(deleteStatement); err != nil {
			return nil, fmt.Errorf("mysql: prepare delete: %v", err)
		}*/

	return db, nil
}

func (db *mysqlDB) AddInfo(u *model.InfoModel) (id int64, err error) {
	r, err := execAffectingOneRow(
		db.insert,
		u.Created,
		u.Crop,
		u.SowingType,
		u.City,
		u.Address,
		u.IDUser)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil
}

func (db *mysqlDB) ListInfo() ([]*model.InfoModel, error) {
	rows, err := db.list.Query()
	if err != nil {
		fmt.Print("error")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var infos []*model.InfoModel
	for rows.Next() {
		info, err := scanInfo(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		infos = append(infos, info)
	}
	return infos, nil
}

func (db *mysqlDB) GetInfo(id int64) {

}

func (db *mysqlDB) UpdateInfo(b *model.InfoModel) error {
	if b.ID == 0 {
		return errors.New("mysql: book with unassigned ID passed into updateBook")
	}

	_, err := execAffectingOneRow(db.update,
		b.Created,
		b.Crop,
		b.SowingType,
		b.City,
		b.Address,
		b.IDUser,
		b.ID)
	return err
}

func (db *mysqlDB) DeleteInfo(id int64) error {
	if id == 0 {
		return errors.New("mysql: Info with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}
