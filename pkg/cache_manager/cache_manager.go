package cache_manager

import (
	"errors"
	"sync"
	"time"
)

const (
	CacheDefaultExpiration = 20 * time.Minute // Время жизни кеша по умолчанию
	CacheCleanupTime       = 20 * time.Minute // Время очистки кеша по умолчанию
)

// Cache - Тип данных, реализующий менеджер кеша для работы с кешируемыми данными
type Cache struct {
	sync.RWMutex                       // Асинхронность для корректного доступа для чтения и записи
	defaultExpiration time.Duration    // Продолжительность жизни кеша по умолчанию
	cleanupTime       time.Duration    // Интервал, после которого запускается очистка
	data              map[string]Value // Непосредственно кешируемые данные
}

// Value - Тип данных, реализующий структуру конкретного элемента кеша
type Value struct {
	CreateTime time.Time // Время создания
	Expiration int64     // Время истечения актуальности
	Value      string    // Непосредственно значение
}

// CacheCreate - Функция, реализующая создание кеша
func CacheCreate() *Cache {

	data := make(map[string]Value)

	cache := Cache{
		data:              data,
		defaultExpiration: CacheDefaultExpiration,
		cleanupTime:       CacheCleanupTime,
	}

	cache.startGC()

	return &cache
}

// Set - Метод, реализующий добавление заданных значений в кеш
func (c *Cache) Set(key string, value string, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()
	defer c.Unlock()

	c.data[key] = Value{
		Value:      value,
		Expiration: expiration,
		CreateTime: time.Now(),
	}

}

// Get - Метод, реализующий получение кеша по заданному ключу
func (c *Cache) Get(key string) (string, bool) {

	c.RLock()
	defer c.RUnlock()

	item, found := c.data[key]

	if !found {
		return "", false
	}

	if item.Expiration > 0 &&
		time.Now().UnixNano() > item.Expiration {
		return "", false
	}

	return item.Value, true
}

// Delete - Метод, реализующий удаление элемента кеша
func (c *Cache) Delete(key string) error {

	c.Lock()
	defer c.Unlock()

	if _, found := c.data[key]; !found {
		return errors.New("error: Key not found")
	}

	delete(c.data, key)

	return nil
}

// startGC - Метод, реализующий запуск очистки кеша
func (c *Cache) startGC() {
	go c.gC()
}

// gC - Метод, реализующий очистку кеша
func (c *Cache) gC() {

	for {
		<-time.After(c.cleanupTime)

		if c.data == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearValues(keys)
		}
	}
}

// expiredKeys - Метод, реализующий поиск неактуального кеша
func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()
	defer c.RUnlock()

	for k, i := range c.data {
		if i.Expiration > 0 &&
			time.Now().UnixNano() > i.Expiration {
			keys = append(keys, k)
		}
	}

	return
}

// clearValues - Метод, реализующий очистку кеша по значению ключей
func (c *Cache) clearValues(keys []string) {

	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.data, k)
	}
}
