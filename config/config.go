package config

import "time"

const (
	GenUrl                 = "http://exmpl.lnk/" // Основа генерируемой короткой ссылки
	ServerPort             = ":4000"             // Порт, на котором развернуто приложение
	TableNameDB            = " \"GenTable\""     // Название таблицы в БД (начинается с пробела)
	UrlColName             = "url"               // Название столбца с исходными ссылками в БД
	ShortUrlColName        = "short_url"         // Название столбца с короткими ссылками в БД
	ShortUrlLen            = 10                  // Длина части выходной короткой ссылки после длины основы "GenUrl" (до 32 символов)
	CacheDefaultExpiration = 20 * time.Minute    // Время жизни кеша по умолчанию
	CacheCleanupTime       = 20 * time.Minute    // Время очистки кеша по умолчанию
)
