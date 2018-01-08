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
const listStatementUser = `SELECT * FROM ` + tableNameUser + `ORDER BY name`
const getStatementUser = `SELECT * FROM ` + tableNameUser + ` WHERE email = ? ORDER BY name`
const updateStatementUser = `
UPDATE ` + tableNameUser + `
SET name=?, lastname=?, email=?, pass=?, created=?
WHERE id = ?`
const deleteStatementUser = `DELETE FROM ` + tableNameUser + ` WHERE id = ?`

var createTableStatementsUser = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameUser + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		lastname VARCHAR(255) NULL,
		email VARCHAR(255) NULL,
		pass VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Note
const tableNameNote = "note"
const insertStatementNote = `
  INSERT INTO ` + tableNameNote + ` (
    title, body, tag, alarm, created, edit, iduser
  ) VALUES (?, ?, ?, ?, ?, ?, ?)`
const listStatementNote = `SELECT * FROM ` + tableNameNote + `ORDER BY created`
const getStatementNote = `SELECT * FROM ` + tableNameNote + ` WHERE iduser = ? ORDER BY created`
const updateStatementNote = `
UPDATE ` + tableNameNote + `
SET title=?, body=?, tag=?, alarm=?, created=?,
		edit=?, iduser=?
WHERE id = ?`
const deleteStatementNote = `DELETE FROM ` + tableNameNote + ` WHERE id = ?`

var createTableStatementsNote = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameNote + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NULL,
		body VARCHAR(255) NULL,
		tag VARCHAR(255) NULL,
		alarm VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		edit VARCHAR(255) NULL,
		iduser INT,
		PRIMARY KEY (id)
	);`,
}

//Alarm
const tableNameAlarm = "alarm"
const insertStatementAlarm = `
  INSERT INTO ` + tableNameAlarm + ` (
    title, created, alarm, iduser
  ) VALUES (?, ?, ?, ?)`
const listStatementAlarm = `SELECT * FROM ` + tableNameAlarm + `ORDER BY created`
const getStatementAlarm = `SELECT * FROM ` + tableNameAlarm + ` WHERE iduser = ? ORDER BY created`
const updateStatementAlarm = `
UPDATE ` + tableNameAlarm + `
SET title=?, created=?, alarm=?, iduser=?
WHERE id = ?`
const deleteStatementAlarm = `DELETE FROM ` + tableNameAlarm + ` WHERE id = ?`

var createTableStatementsAlarm = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameAlarm + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		alarm VARCHAR(255) NULL,
		iduser INT,
		PRIMARY KEY (id)
	);`,
}

//Info
const tableNameInfo = "info"
const insertStatementInfo = `
  INSERT INTO ` + tableNameInfo + ` (
    crop, created, sowing, city, address, iduser
  ) VALUES (?, ?, ?, ?, ?, ?)`
const listStatementInfo = `SELECT * FROM ` + tableNameInfo + `ORDER BY created`
const getStatementInfo = `SELECT * FROM ` + tableNameInfo + ` WHERE iduser = ? ORDER BY created`
const updateStatementInfo = `
UPDATE ` + tableNameInfo + `
SET crop=?, created=?, sowing=?, city=?, address=?, iduser=?
WHERE id = ?`
const deleteStatementInfo = `DELETE FROM ` + tableNameInfo + ` WHERE id = ?`

var createTableStatementsInfo = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameInfo + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		crop VARCHAR(255) NULL,
		created VARCHAR(255) NULL,
		sowing VARCHAR(255) NULL,
		city VARCHAR(255) NULL,
		address VARCHAR(255) NULL,
		iduser INT,
		PRIMARY KEY (id)
	);`,
}

//Plant
const tableNamePlant = "plant"
const insertStatementPlant = `
  INSERT INTO ` + tableNamePlant + ` (
    name, kingdom, season, temperature, idnutrient, idpests
  ) VALUES (?, ?, ?, ?, ?,?)`
const listStatementPlant = `SELECT * FROM ` + tableNamePlant + `ORDER BY name`
const getStatementPlant = `SELECT * FROM ` + tableNamePlant + ` WHERE name = ? ORDER BY name`
const updateStatementPlant = `
UPDATE ` + tableNamePlant + `
SET name=?, kingdom=?, idnutrient=?, idweather=?, idpests=?
WHERE id = ?`
const deleteStatementPlant = `DELETE FROM ` + tableNamePlant + ` WHERE id = ?`

var createTableStatementsPlant = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNamePlant + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NULL,
		kingdom VARCHAR(255) NULL,
		season VARCHAR(255) NULL,
		temperature VARCHAR(255) NULL,
		idnutrient INT NULL,
		idpests INT NULL,
		PRIMARY KEY (id)
	);`,
}

//Plant - Nutrient
const tableNameNutrient = "nutrient"
const insertStatementNutrient = `
  INSERT INTO ` + tableNameNutrient + ` (
    formula, advantage
  ) VALUES (?, ?)`
const listStatementNutrient = `SELECT * FROM ` + tableNameNutrient + `ORDER BY formula`
const getStatementNutrient = `SELECT * FROM ` + tableNameNutrient + ` WHERE id = ? ORDER BY formula`
const updateStatementNutrient = `
UPDATE ` + tableNameNutrient + `
SET formula=?, advantage=?
WHERE id = ?`
const deleteStatementNutrient = `DELETE FROM ` + tableNameNutrient + ` WHERE id = ?`

var createTableStatementsNutrient = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameNutrient + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		formula VARCHAR(255) NULL,
		advantage VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Plant - Pests
const tableNamePests = "pests"
const insertStatementPests = `
  INSERT INTO ` + tableNamePests + ` (
    category, causes, consequences, solution
  ) VALUES (?, ?, ?, ?)`
const listStatementPests = `SELECT * FROM ` + tableNamePests + `ORDER BY category`
const getStatementPests = `SELECT * FROM ` + tableNamePests + ` WHERE category = ? ORDER BY category`
const updateStatementPests = `
UPDATE ` + tableNamePests + `
SET category=?, causes=?, consequences=?, solution=?
WHERE id = ?`
const deleteStatementPests = `DELETE FROM ` + tableNamePests + ` WHERE id = ?`

var createTableStatementsPests = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNamePests + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		category VARCHAR(255) NULL,
		causes VARCHAR(255) NULL,
		consequences VARCHAR(255) NULL,
		solution VARCHAR(255) NULL,
		PRIMARY KEY (id)
	);`,
}

//Scan
const tableNameScan = "scan"
const insertStatementScan = `
  INSERT INTO ` + tableNameScan + ` (
    created, distance, zone_width, zone_height, latitude, longitude, city, iduser
  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
const listStatementScan = `SELECT * FROM ` + tableNameScan + `ORDER BY created`
const getStatementScan = `SELECT * FROM ` + tableNameScan + ` WHERE iduser = ? ORDER BY created`
const updateStatementScan = `
UPDATE ` + tableNameScan + `
SET created=?, distance=?, zone_width=?, zone_height=?, latitude=?, longitude=?, city=?, iduser=?
WHERE id = ?`
const deleteStatementScan = `DELETE FROM ` + tableNameScan + ` WHERE id = ?`

var createTableStatementsScan = []string{
	`CREATE DATABASE IF NOT EXISTS ` + utils.Namedb + ` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE ` + utils.Namedb + `;`,
	`CREATE TABLE IF NOT EXISTS ` + tableNameScan + `  (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		created VARCHAR(255) NULL,
		distance VARCHAR(255) NULL,
		zone_width VARCHAR(255) NULL,
		zone_height VARCHAR(255) NULL,
		latitude VARCHAR(255) NULL,
		longitude VARCHAR(255) NULL,
		city VARCHAR(255) NULL,
		iduser INT,
		PRIMARY KEY (id)
	);`,
}
