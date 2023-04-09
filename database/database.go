package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

// RowData - Тип данных, реализующий структуру для работы с данными в строке БД
type RowData struct {
	Id       int    // (serial, not null)
	Url      string // (text, not null)
	ShortUrl string // (text, primary_key, not null)
}

// Database - Тип данных, реализующий структуру для более удобной работы с БД и подключением в ней
type Database struct {
	db *pgx.Conn // База данных
}

// GetConnection - Функция, позволяющая подключиться к БД
func GetConnection() (*Database, error) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return &Database{}, err
	}

	return &Database{conn}, nil
}

// GetUrlRow - Метод, позволяющий получить строку из БД по заданной исходной ссылке
func (c *Database) GetUrlRow(url string) (*RowData, bool) {

	var row pgx.Row

	sql := "SELECT * FROM \"GenTable\" WHERE url = $1"

	row = c.db.QueryRow(context.Background(), sql, url)

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		return nil, false
	}

	return &r, true
}

// GetShortUrlRow - Метод, позволяющий получить строку из БД по заданной короткой ссылке
func (c *Database) GetShortUrlRow(shortUrl string) (*RowData, bool) {

	var row pgx.Row

	sql := "SELECT * FROM \"GenTable\" WHERE short_url = $1"

	row = c.db.QueryRow(context.Background(), sql, shortUrl)

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		return nil, false
	}

	return &r, true
}

// SaveShortUrl - Метод, позволяющий сохранить в БД заданную строку
func (c *Database) SaveShortUrl(row RowData) error {

	sql := "INSERT INTO \"GenTable\" (url, short_url) VALUES ($1, $2)"

	_, err := c.db.Exec(context.Background(), sql, row.Url, row.ShortUrl)
	if err != nil {
		return err
	}

	return nil
}

// CloseConnection - Метод, реализующий закрытие соединения с БД
func (c *Database) CloseConnection() error {

	err := c.db.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}
