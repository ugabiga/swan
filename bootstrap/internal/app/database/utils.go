package database

import (
	"fmt"

	"gorm.io/gorm"
)

func DeleteAllDataFromTables(db *gorm.DB, tables []string) error {
	var wrappedErr error

	for _, table := range tables {
		err := db.Exec("DELETE " + "FROM " + table).Error
		if err != nil {
			wrappedErr = fmt.Errorf("error deleting from %s: %w", table, err)
		}
	}

	return wrappedErr
}

func DeleteAllDataFromModels(db *gorm.DB, models []interface{}) error {
	var wrappedErr error

	for _, model := range models {
		table := GetTableName(db, model)

		err := db.Exec("DELETE " + "FROM " + table).Error
		if err != nil {
			wrappedErr = fmt.Errorf("error deleting from %s: %w", table, err)
		}
	}

	return wrappedErr
}

func GetTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(model)
	if err != nil {
		return ""
	}

	return stmt.Table
}
