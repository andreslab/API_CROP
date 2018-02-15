package database

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"log"

// 	"github.com/andreslab/prj_api_crop/model"
// )

// func scanScan(s rowScanner) (*model.ScanModel, error) {
// 	var (
// 		id         int64
// 		created    sql.NullString
// 		distance   sql.NullString
// 		zoneWidth  sql.NullString
// 		zoneHeight sql.NullString
// 		latitude   sql.NullString
// 		longitude  sql.NullString
// 		city       sql.NullString
// 		iduser     int64
// 	)
// 	if err := s.Scan(
// 		&id,
// 		&created,
// 		&distance,
// 		&zoneWidth,
// 		&zoneHeight,
// 		&latitude,
// 		&longitude,
// 		&city,
// 		&iduser); err != nil {
// 		return nil, err
// 	}

// 	scan := &model.ScanModel{
// 		ID:         id,
// 		Created:    created.String,
// 		Distance:   distance.String,
// 		ZoneWidth:  zoneWidth.String,
// 		ZoneHeight: zoneHeight.String,
// 		Latitude:   latitude.String,
// 		Longitude:  longitude.String,
// 		City:       city.String,
// 		IDUser:     iduser,
// 	}
// 	return scan, nil
// }

// // newMySQLDB creates a new UserDatabase backed by a given MySQL server.
// func NewMySQLDBScan() (*mysqlDB, error) {

// 	// Check database and table exists. If not, create it.
// 	err := ensureTableExists(tableNameScan, createTableStatementsScan)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conn := ConnDB()
// 	if err := conn.Ping(); err != nil {
// 		conn.Close()
// 		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
// 	}

// 	db := &mysqlDB{
// 		conn: conn,
// 	}
// 	// Prepared statements. The actual SQL queries are in the code near the
// 	// relevant method (e.g. addBook).
// 	if db.list, err = conn.Prepare(listStatementScan); err != nil {
// 		return nil, fmt.Errorf("mysql: prepare list: %v", err)
// 	}
// 	if db.get, err = conn.Prepare(getStatementScan); err != nil {
// 		return nil, fmt.Errorf("mysql: prepare get: %v", err)
// 	}
// 	if db.insert, err = conn.Prepare(insertStatementScan); err != nil {
// 		return nil, fmt.Errorf("mysql: prepare insert: %v", err)
// 	}
// 	/*
// 		if db.listBy, err = conn.Prepare(listByStatement); err != nil {
// 			return nil, fmt.Errorf("mysql: prepare listBy: %v", err)
// 		}
// 		if db.update, err = conn.Prepare(updateStatement); err != nil {
// 			return nil, fmt.Errorf("mysql: prepare update: %v", err)
// 		}
// 		if db.delete, err = conn.Prepare(deleteStatement); err != nil {
// 			return nil, fmt.Errorf("mysql: prepare delete: %v", err)
// 		}*/

// 	return db, nil
// }

// func (db *mysqlDB) AddScan(u *model.ScanModel) (id int64, err error) {
// 	r, err := execAffectingOneRow(
// 		db.insert,
// 		u.Created,
// 		u.Distance,
// 		u.ZoneWidth,
// 		u.ZoneHeight,
// 		u.Latitude,
// 		u.Longitude,
// 		u.City,
// 		u.IDUser,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	lastInsertID, err := r.LastInsertId()
// 	if err != nil {
// 		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
// 	}
// 	return lastInsertID, nil
// }

// func (db *mysqlDB) ListScan() ([]*model.ScanModel, error) {
// 	rows, err := db.list.Query()
// 	if err != nil {
// 		fmt.Print("error")
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var scans []*model.ScanModel
// 	for rows.Next() {
// 		scan, err := scanScan(rows)
// 		if err != nil {
// 			return nil, fmt.Errorf("mysql: could not read row: %v", err)
// 		}

// 		scans = append(scans, scan)
// 	}
// 	return scans, nil
// }

// func (db *mysqlDB) GetScan(id int64) {

// }

// func (db *mysqlDB) UpdateScan(b *model.ScanModel) error {
// 	if b.ID == 0 {
// 		return errors.New("mysql: book with unassigned ID passed into updateBook")
// 	}

// 	_, err := execAffectingOneRow(db.update,
// 		b.Created,
// 		b.Distance,
// 		b.ZoneWidth,
// 		b.ZoneHeight,
// 		b.Latitude,
// 		b.Longitude,
// 		b.City,
// 		b.IDUser,
// 		b.ID)
// 	return err
// }

// func (db *mysqlDB) DeleteScan(id int64) error {
// 	if id == 0 {
// 		return errors.New("mysql: Scan with unassigned ID passed into deleteBook")
// 	}
// 	_, err := execAffectingOneRow(db.delete, id)
// 	return err
// }
