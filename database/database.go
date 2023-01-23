package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"my_project/urlgen/config"
	"os"
)

type RowData struct {
	Id       int    // (serial, not null)
	Url      string // (text, not null)
	ShortUrl string // (text, primary_key, not null)
}

type Connection struct {
	conn *pgx.Conn
}

func GetConnection() (Connection, error) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return Connection{}, err
	}

	newConnection := Connection{conn}

	return newConnection, nil
}

func (c Connection) GetUrlRow(url string, isShortUrl bool) (*RowData, bool) {
	var row pgx.Row

	if isShortUrl {
		row = c.conn.QueryRow(context.Background(),
			"SELECT * FROM"+config.TableNameDB+" WHERE "+config.ShortUrlColName+" = $1", url)
	} else {
		row = c.conn.QueryRow(context.Background(),
			"SELECT * FROM"+config.TableNameDB+" WHERE "+config.UrlColName+" = $1", url)
	}

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		return nil, false
	}

	return &r, true
}

func (c Connection) SaveShortUrl(row RowData) error {
	_, err := c.conn.Exec(context.Background(), "INSERT INTO"+config.TableNameDB+
		" ("+config.UrlColName+", "+config.ShortUrlColName+") VALUES ($1, $2)", row.Url, row.ShortUrl)
	if err != nil {
		return err
	}

	return nil
}

func (c Connection) CloseConnection() error {
	err := c.conn.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}
