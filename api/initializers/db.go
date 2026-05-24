package initializers

import (
	"database/sql"
	"os"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	DB, err := sql.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		return nil, err
	}
	err = SetupDB(DB)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func SetupDB(db *sql.DB) error {
	dbInit := `SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
CREATE TABLE IF NOT EXISTS ` + "`task`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`date_add`" + ` datetime DEFAULT NULL,
  ` + "`date_to`" + ` datetime DEFAULT NULL,
  ` + "`title`" + ` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  ` + "`content`" + ` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  ` + "`is_done`" + ` tinyint(1) NOT NULL,
  ` + "`is_common`" + ` tinyint(1) NOT NULL DEFAULT 0,
  ` + "`ref_user`" + ` int DEFAULT NULL,
  PRIMARY KEY (` + "`id`" + `),
  KEY ` + "`ref_user`" + ` (` + "`ref_user`" + `)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE IF NOT EXISTS ` + "`common_task`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`ref_user`" + ` int NOT NULL,
  ` + "`title`" + ` varchar(255) NOT NULL,
  ` + "`date_add`" + ` datetime DEFAULT NULL,
  PRIMARY KEY (` + "`id`" + `),
  KEY ` + "`ref_user`" + ` (` + "`ref_user`" + `)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE IF NOT EXISTS ` + "`user`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`username`" + ` varchar(255) NOT NULL,
  ` + "`password`" + ` varchar(255) NOT NULL,
  ` + "`date_add`" + ` datetime DEFAULT NULL,
  ` + "`is_active`" + ` tinyint(1) NOT NULL,
  ` + "`token`" + ` varchar(255) NOT NULL,
  PRIMARY KEY (` + "`id`" + `)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS ` + "`note`" + ` (
  ` + "`id`" + ` int NOT NULL AUTO_INCREMENT,
  ` + "`ref_user`" + ` int DEFAULT NULL,
  ` + "`title`" + ` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  ` + "`content`" + ` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  ` + "`date_add`" + ` datetime DEFAULT NULL,
  PRIMARY KEY (` + "`id`" + `),
  KEY ` + "`ref_user`" + ` (` + "`ref_user`" + `)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

	_, err := db.Exec(dbInit)
	return err
}

func ExecFlushDB(db *sql.DB) error {
	dbQuery := "DROP TABLE task;DROP TABLE user"
	if _, err := db.Exec(dbQuery); err != nil {
		return err
	}
	if err := SetupDB(db); err != nil {
		return err
	}
	return nil
}
