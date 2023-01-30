package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"my_project/urlgen/config"
	"os"
)

// RowData - Тип данных, реализующий структуру для работы с данными в строке БД
type RowData struct {
	Id       int    // (serial, not null)
	Url      string // (text, not null)
	ShortUrl string // (text, primary_key, not null)
}

// Connection - Тип данных, реализующий структуру для более удобной работы с БД и подключением в ней
type Connection struct {
	conn *pgx.Conn
}

// GetConnection - Функция, позволяющая подключиться к БД
func GetConnection() (Connection, error) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return Connection{}, err
	}

	newConnection := Connection{conn}

	return newConnection, nil
}

// GetUrlRow - Метод, позволяющий получить строку из БД по заданной исходной ссылке
func (c Connection) GetUrlRow(url string) (*RowData, bool) {

	var row pgx.Row

	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", config.TableNameDB, config.UrlColName)

	row = c.conn.QueryRow(context.Background(), sql, url)

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		return nil, false
	}

	return &r, true
}

// GetShortUrlRow - Метод, позволяющий получить строку из БД по заданной короткой ссылке
func (c Connection) GetShortUrlRow(shortUrl string) (*RowData, bool) {

	var row pgx.Row

	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", config.TableNameDB, config.ShortUrlColName)

	row = c.conn.QueryRow(context.Background(), sql, shortUrl)

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		return nil, false
	}

	return &r, true
}

// SaveShortUrl - Метод, позволяющий сохранить в БД заданную строку
func (c Connection) SaveShortUrl(row RowData) error {

	_, err := c.conn.Exec(context.Background(), "INSERT INTO"+config.TableNameDB+
		" ("+config.UrlColName+", "+config.ShortUrlColName+") VALUES ($1, $2)", row.Url, row.ShortUrl)
	if err != nil {
		return err
	}

	return nil
}

// CloseConnection - Метод, реализующий закрытие соединения с БД
func (c Connection) CloseConnection() error {

	err := c.conn.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}
