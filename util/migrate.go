package util

import (
	"database/sql"
)

func GetFilesExecute(dir string, db *sql.DB) ([]string, error) {
	executed, err := GetExecutedIDs(db)
	if err != nil {
		return nil, err
	}

	files, err := GetFiles(dir)
	return Diff(files, executed), err
}
