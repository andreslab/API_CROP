package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/andreslab/prj_api_crop/model"
)

func scanNote(s rowScanner) (*model.ActionCreateNoteModel, error) {
	var (
		id      int64
		title   sql.NullString
		body    sql.NullString
		tag     sql.NullString
		alarm   sql.NullString
		created sql.NullString
		edit    sql.NullString
		iduser  int64
	)
	if err := s.Scan(&id,
		&title,
		&body,
		&tag,
		&alarm,
		&created,
		&edit,
		iduser); err != nil {
		return nil, err
	}

	note := &model.ActionCreateNoteModel{
		ID:          id,
		Title:       title.String,
		Body:        body.String,
		Tag:         tag.String,
		Alarm:       alarm.String,
		DateCreated: created.String,
		DateEdit:    edit.String,
		IDUser:      iduser,
	}
	return note, nil
}

// newMySQLDB creates a new UserDatabase backed by a given MySQL server.
func NewMySQLDBNote() (*mysqlDB, error) {

	// Check database and table exists. If not, create it.
	err := ensureTableExists(tableNameNote, createTableStatementsNote)
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
	if db.list, err = conn.Prepare(listStatementNote); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.get, err = conn.Prepare(getStatementNote); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatementNote); err != nil {
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

func (db *mysqlDB) AddNote(u *model.ActionCreateNoteModel) (id int64, err error) {
	r, err := execAffectingOneRow(
		db.insert,
		u.Title,
		u.Body,
		u.Tag,
		u.Alarm,
		u.DateCreated,
		u.DateEdit,
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

func (db *mysqlDB) ListNote() ([]*model.ActionCreateNoteModel, error) {
	rows, err := db.list.Query()
	if err != nil {
		fmt.Print("error")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var notes []*model.ActionCreateNoteModel
	for rows.Next() {
		note, err := scanNote(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		notes = append(notes, note)
	}
	return notes, nil
}

func (db *mysqlDB) GetNote(id int64) {

}

func (db *mysqlDB) UpdateNote(b *model.ActionCreateNoteModel) error {
	if b.ID == 0 {
		return errors.New("mysql: book with unassigned ID passed into updateBook")
	}

	_, err := execAffectingOneRow(db.update,
		b.Title,
		b.Body,
		b.Tag,
		b.Alarm,
		b.DateCreated,
		b.DateEdit,
		b.ID)
	return err
}

func (db *mysqlDB) DeleteNote(id int64) error {
	if id == 0 {
		return errors.New("mysql: User with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}
