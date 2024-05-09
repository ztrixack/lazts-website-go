package book

import (
	"encoding/json"
	"lazts/internal/app/models"
	"lazts/internal/utils"
	"math/rand/v2"
	"os"
)

func getCatalogs(books []Book) []models.Option {
	catalogs := models.Options{models.Option{Key: "ทั้งหมด", Value: ""}}
	for _, book := range books {
		catalogs.AppendUnique(book.Catalog)
	}
	catalogs.Sort()

	return catalogs
}

func getList(name string) ([]Book, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	books := make([]Book, 0)
	for _, dir := range dirs {
		if !dir.IsDir() {
			bytes, err := os.ReadFile(utils.GetContentDir(name, dir.Name()))
			if err != nil {
				return nil, err
			}

			var book []Book
			if err := json.Unmarshal(bytes, &book); err != nil {
				return nil, err
			}

			books = append(books, book...)
		}
	}

	return books, nil
}

func getStatus() []models.Option {
	return []models.Option{
		{Key: "กำลังจะซื้อ", Value: "wishlist"},
		{Key: "กำลังอ่าน", Value: "reading"},
		{Key: "อ่านจบแล้ว", Value: "done"},
		{Key: "ทั้งหมด", Value: ""},
	}
}

func randomizeOne(books []Book, count int) []Book {
	result := make([]Book, count)

	for i := range result {
		r := rand.IntN(len(books))
		result[i] = books[r]
	}

	return result
}

func getStats(books []Book) (int, int, float64) {
	all := len(books)
	read := 0
	for _, book := range books {
		if book.Status == "done" {
			read++
		}
	}
	percent := float64(read) * 100 / float64(all)
	return all, read, percent
}
