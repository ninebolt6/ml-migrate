package util

import (
	"database/sql"
	"fmt"
	"log"
	"path"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(connectionStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("cannot open a connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot connect to the database: %w", err)
	}

	if err = createTableIfNotExists(db); err != nil {
		return nil, err
	}
	return db, nil
}

func GetExecutedIDs(db *sql.DB) ([]string, error) {
	var result []string
	rows, err := db.Query("SELECT CONCAT(`id`, '.sql') FROM `_migration`;")
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch executed migration IDs: %w", err)
	}
	for rows.Next() {
		var id string
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	return result, nil
}

func createTableIfNotExists(db *sql.DB) error {
	if err := db.QueryRow("SELECT 1 FROM `_migration` LIMIT 1;").Err(); err != nil {
		log.Println("A migration table not found. Create a new one.")

		createTable := "CREATE TABLE `_migration` (" +
			"`id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL," +
			"`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
			"`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;"

		if _, err := db.Exec(createTable); err != nil {
			return fmt.Errorf("cannot create a migration table: %w", err)
		}
	}
	return nil
}

func ExecMigration(db *sql.DB, folderPath string, files []string) error {
	for _, file := range files {
		filePath := path.Join(folderPath, file)
		data, err := ReadFileAsString(filePath)
		if err != nil {
			return fmt.Errorf("failed to load file: %s", filePath)
		}
		_, err = db.Exec(data)
		if err != nil {
			db.Exec("ROLLBACK")
			return fmt.Errorf("failed to execute query: %w", err)
		}
		putID(db, GetFileNameWithoutExt(file))
	}
	return nil
}

func putID(db *sql.DB, id string) error {
	_, err := db.Exec("INSERT INTO `_migration` (`id`, `created_at`, `updated_at`) VALUES (?, CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())", id)
	if err != nil {
		return fmt.Errorf("ID insertion failed: %w", err)
	}
	fmt.Println("ID insertion succeeded:", id)
	return nil
}
