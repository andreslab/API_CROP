package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/andreslab/prj_api_crop/model"
)

func scanAlarm(s rowScanner) (*model.ActionCreateAlarmModel, error) {
	var (
		id      int64
		title   sql.NullString
		created sql.NullString
		alarm   sql.NullString
		iduser  int64
	)
	if err := s.Scan(&id,
		&title,
		&created,
		&alarm,
		&iduser); err != nil {
		return nil, err
	}

	_alarm := &model.ActionCreateAlarmModel{
		ID:          id,
		Title:       title.String,
		DateCreated: created.String,
		DateAlarm:   alarm.String,
		IDUser:      iduser,
	}
	return _alarm, nil
}

// newMySQLDB creates a new UserDatabase backed by a given MySQL server.
func NewMySQLDBAlarm() (*mysqlDB, error) {

	// Check database and table exists. If not, create it.
	err := ensureTableExists(tableNameAlarm, createTableStatementsAlarm)
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
	if db.list, err = conn.Prepare(listStatementAlarm); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.get, err = conn.Prepare(getStatementAlarm); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatementAlarm); err != nil {
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

func (db *mysqlDB) AddAlarm(u *model.ActionCreateAlarmModel) (id int64, err error) {
	r, err := execAffectingOneRow(
		db.insert,
		u.Title,
		u.DateCreated,
		u.DateAlarm,
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

func (db *mysqlDB) ListAlarm() ([]*model.ActionCreateAlarmModel, error) {
	rows, err := db.list.Query()
	if err != nil {
		fmt.Print("error")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var alarms []*model.ActionCreateAlarmModel
	for rows.Next() {
		_alarm, err := scanAlarm(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		alarms = append(alarms, _alarm)
	}
	return alarms, nil
}

func (db *mysqlDB) GetAlarm(id int64) {

}

func (db *mysqlDB) UpdateAlarm(b *model.ActionCreateAlarmModel) error {
	if b.ID == 0 {
		return errors.New("mysql: book with unassigned ID passed into updateBook")
	}

	_, err := execAffectingOneRow(db.update,
		b.Title,
		b.DateCreated,
		b.DateAlarm,
		b.IDUser,
		b.ID)
	return err
}

func (db *mysqlDB) DeleteAlarm(id int64) error {
	if id == 0 {
		return errors.New("mysql: User with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}
