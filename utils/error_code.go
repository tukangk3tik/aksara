package utils

var ErrorCodeMap = map[string]string{
	"BAD_REQUEST":           "Permintaan tidak valid / data tidak lengkap",
	"UNAUTHORIZED":          "Akses tidak diizinkan",
	"FORBIDDEN":             "Akses dibatalkan",
	"NOT_FOUND":             "%s tidak ditemukan",
	"INTERNAL_SERVER_ERROR": "Terjadi Kesalahan",
	"WRONG_PASSWORD":        "Password salah",
	"FOREIGN_KEY_VIOLATION": "Kesalahan pada kunci asing",
	"DUPLICATE_ENTRY":       "Data yang sama sudah ada",
}
