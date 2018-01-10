package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	modelDB "github.com/andreslab/prj_api_crop/database/model"
)

func scanPlant(s rowScanner) (*modelDB.PlantModelDB, error) {
	var (
		id          int64
		name        sql.NullString
		kingdom     sql.NullString
		season      sql.NullString
		temperature sql.NullString
		idnutrient  int64
		idpests     int64
	)
	if err := s.Scan(&id,
		&name,
		&kingdom,
		&season,
		&temperature,
		&idnutrient,
		&idpests); err != nil {
		return nil, err
	}

	plant := &modelDB.PlantModelDB{
		ID:             id,
		Name:           name.String,
		Kingdom:        kingdom.String,
		SeasonEnv:      season.String,
		TemperatureEnv: temperature.String,
		IDNutrient:     idnutrient,
		IDPests:        idpests,
	}
	return plant, nil
}

// newMySQLDB creates a new UserDatabase backed by a given MySQL server.
func NewMySQLDBPlant() (*mysqlDB, error) {

	// Check database and table exists. If not, create it.
	err := ensureTableExists(tableNamePlant, createTableStatementsPlant)
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
	if db.list, err = conn.Prepare(listStatementPlant); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}
	if db.get, err = conn.Prepare(getStatementPlant); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatementPlant); err != nil {
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

func (db *mysqlDB) AddPLant(u *modelDB.PlantModelDB) (id int64, err error) {
	r, err := execAffectingOneRow(
		db.insert,
		u.Name,
		u.Kingdom,
		u.SeasonEnv,
		u.TemperatureEnv,
		u.IDNutrient,
		u.IDPests)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil
}

func (db *mysqlDB) ListPlant() ([]*modelDB.PlantModelDB, error) {
	rows, err := db.list.Query()
	if err != nil {
		fmt.Print("error")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var plants []*modelDB.PlantModelDB
	for rows.Next() {
		plant, err := scanPlant(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		plants = append(plants, plant)
	}
	return plants, nil
}

func (db *mysqlDB) GetPlant(id int64) {

}

func (db *mysqlDB) UpdatePlant(b *modelDB.PlantModelDB) error {
	if b.ID == 0 {
		return errors.New("mysql: book with unassigned ID passed into updateBook")
	}

	_, err := execAffectingOneRow(db.update,
		b.Name,
		b.Kingdom,
		b.SeasonEnv,
		b.TemperatureEnv,
		b.IDNutrient,
		b.IDPests,
		b.ID)
	return err
}

func (db *mysqlDB) DeletePlant(id int64) error {
	if id == 0 {
		return errors.New("mysql: Plant with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}
