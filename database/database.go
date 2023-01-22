package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"my_project/urlgen/config"
	"os"
)

type RowData struct {
	Id       int    // (serial, primary_key, not null)
	Url      string // (text, primary_key, not null)
	ShortUrl string // (text, primary_key, not null)
}

type Connection struct {
	conn *pgx.Conn
}

func GetConnection() (Connection, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
		return Connection{}, err
	}

	newConnection := Connection{conn}

	return newConnection, nil
}

func (c Connection) GetOriginalUrlRow(shortUrl string) (*RowData, error) {
	row := c.conn.QueryRow(context.Background(), "SELECT * FROM "+config.TableNameDB+" WHERE "+config.ShortUrlColName+
		" = $1", shortUrl)

	r := RowData{}

	err := row.Scan(&r.Id, &r.Url, &r.ShortUrl)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &r, nil
}

func (c Connection) SaveShortUrl(row RowData) error {
	_, err := c.conn.Exec(context.Background(), "INSERT INTO "+config.TableNameDB+
		" (url, short_url) VALUES ($1, $2)", row.Url, row.ShortUrl)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func (c Connection) CloseConnection() error {
	err := c.conn.Close(context.Background())
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
