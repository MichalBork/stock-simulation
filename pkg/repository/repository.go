package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	ID   int
	Data string
}

type Repository interface {
	FindBy(table string, column string, value string) ([]Item, error)
	Save(table string, data map[string]interface{}) error
}

type SQLRepository struct {
	DB *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		DB: db,
	}
}
func (r *SQLRepository) FindBy(table string, column string, value string) (*sql.Rows, error) {
	if numValue, err := strconv.Atoi(value); err == nil {
		// Wartość jest liczbą, więc nie używamy cudzysłowów
		query := fmt.Sprintf("SELECT * FROM %s WHERE %s = %d", table, column, numValue)
		fmt.Println(query)
		return r.DB.Query(query)
	} else {
		// Wartość jest ciągiem znaków, więc używamy cudzysłowów
		query := fmt.Sprintf("SELECT * FROM %s WHERE %s = '%s'", table, column, value)
		fmt.Println(query)
		fmt.Println(r.DB.Query(query))
		return r.DB.Query(query)
	}
}

func (r *SQLRepository) Save(table string, data map[string]interface{}) error {
	// Buduj zapytanie INSERT
	keys := []string{}
	values := []string{}
	valArgs := []interface{}{}

	for k, v := range data {
		keys = append(keys, k)
		values = append(values, "?")
		valArgs = append(valArgs, v)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(keys, ", "), strings.Join(values, ", "))

	_, err := r.DB.Exec(query, valArgs...)
	return err
}
