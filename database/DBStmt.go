package database

import (
	"github.com/andreslab/prj_api_crop/utils"
)

//User
const tableNameUser = "user"
const insertStatementUser = `
  INSERT INTO ` + tableNameUser + ` (
    name, lastname, email, pass, created
  ) VALUES (?, ?, ?, ?, ?)`
const listStatementUser = `SELECT * FROM ` + tableNameUser
const getStatementUser = `SELECT * FROM ` + tableNameUser

var createTableStatementsUser = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + tableNameUser + `;`,
	`CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		lastname VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		pass VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Info
const tableNameInfo = "info"
const insertStatementInfo = `
  INSERT INTO ` + tableNameInfo + ` (
    name, lastname, email, pass, created
  ) VALUES (?, ?, ?, ?, ?)`
const listStatementInfo = `SELECT * FROM ` + tableNameInfo
const getStatementInfo = `SELECT * FROM ` + tableNameInfo

var createTableStatementsInfo = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + tableNameUser + `;`,
	`CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		lastname VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		pass VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Plant
const tableNamePlant = "plant"
const insertStatementPlant = `
  INSERT INTO ` + tableNamePlant + ` (
    name, lastname, email, pass, created
  ) VALUES (?, ?, ?, ?, ?)`
const listStatementPlant = `SELECT * FROM ` + tableNamePlant
const getStatementPlant = `SELECT * FROM ` + tableNamePlant

var createTableStatementsPlant = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + tableNamePlant + `;`,
	`CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		lastname VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		pass VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Scan
const tableNameScan = "scan"
const insertStatementScan = `
  INSERT INTO ` + tableNameScan + ` (
    name, lastname, email, pass, created
  ) VALUES (?, ?, ?, ?, ?)`
const listStatementScan = `SELECT * FROM ` + tableNameScan
const getStatementScan = `SELECT * FROM ` + tableNameScan

var createTableStatementsScan = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + tableNameScan + `;`,
	`CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		lastname VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		pass VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}
