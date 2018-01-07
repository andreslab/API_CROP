package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/andreslab/prj_api_crop/utils"
	"github.com/go-sql-driver/mysql"
)

type mysqlDB struct {
	conn *sql.DB

	list   *sql.Stmt
	listBy *sql.Stmt
	insert *sql.Stmt
	get    *sql.Stmt
	update *sql.Stmt
	delete *sql.Stmt
}

func ConnDB() *sql.DB {
	db, err := sql.Open(utils.TypeDataBase, utils.LinkDataBase)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success: conexión exitosa")
	}
	defer db.Close()
	return db
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

// ensureTableExists checks the table exists. If not, it creates it.
func ensureTableExists(tableName string, stmtCreateTable []string) error {
	conn := ConnDB()
	// Check the connection.
	if conn.Ping() == driver.ErrBadConn {
		return fmt.Errorf("mysql: could not connect to the database. " +
			"could be bad address, or this address is not whitelisted for access.")
	}
	if _, err := conn.Exec("USE " + utils.Namedb); err != nil {
		// MySQL error 1049 is "database does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return nil
		}
	}

	if _, err := conn.Exec("DESCRIBE " + tableName); err != nil {
		// MySQL error 1146 is "table does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			//crear la tabla en caso que no exista
			CreateTable(stmtCreateTable)
			return nil
		}
		// Unknown error.
		return fmt.Errorf("mysql: could not connect to the database: %v", err)
	}
	return nil
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func execAffectingOneRow(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.Exec(args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}

func CreateTable(stmtCreateTable []string) error {
	for _, stmt := range stmtCreateTable {
		_, err := ConnDB().Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
