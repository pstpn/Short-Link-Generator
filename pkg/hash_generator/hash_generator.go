package hash_generator

import (
	"crypto/sha256"
)

// GenerateShortUrl - Функция, реализующая создание уникальной короткой ссылки с помощью алгоритма шифрования SHA256
func GenerateShortUrl(url string) string {

	tmp := sha256.Sum256([]byte(url))

	for i := 0; i < 10; i++ {
		if tmp[i] < 26 || tmp[i] == 96 ||
			(tmp[i] > 90 && tmp[i] < 95) {
			tmp[i] = 65 + tmp[i]%26
		} else if (tmp[i] >= 26 && tmp[i] < 48) ||
			(tmp[i] > 57 && tmp[i] < 65) {
			tmp[i] = 48 + tmp[i]%10
		} else if tmp[i] > 122 {
			tmp[i] = 97 + tmp[i]%26
		}
	}

	return "http://exmpl.lnk/" + string(tmp[:10])
}
