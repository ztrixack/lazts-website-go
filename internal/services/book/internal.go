package book

import (
	"encoding/json"
	"lazts/internal/app/models"
	"lazts/internal/utils"
	"math/rand/v2"
	"os"
	"strings"
)

func getCatalogs(books []Book) []models.Option {
	catalogs := models.Options{models.Option{Key: "ทั้งหมด", Value: ""}}
	for _, book := range books {
		catalogs = catalogs.AppendUnique(book.Catalog)
	}

	return catalogs.Sort()
}

func getList(name string) ([]Book, error) {
	files, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	books := make([]Book, 0)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			bytes, err := os.ReadFile(utils.GetContentDir(name, file.Name()))
			if err != nil {
				return nil, err
			}

			var list []Book
			if err := json.Unmarshal(bytes, &list); err != nil {
				return nil, err
			}

			for i, book := range list {
				if book.Cover == "" {
					list[i].Cover = "https://picsum.photos/640/480"
					continue
				}
				name, _ := strings.CutSuffix(file.Name(), ".json")
				list[i].Cover = utils.GetContentPath("books", name, book.Cover)
			}

			books = append(books, list...)
		}
	}

	return books, nil
}

func getCount(name string) (int, error) {
	files, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return 0, err
	}

	count := 0
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			bytes, err := os.ReadFile(utils.GetContentDir(name, file.Name()))
			if err != nil {
				return 0, err
			}

			var list []Book
			if err := json.Unmarshal(bytes, &list); err != nil {
				return 0, err
			}

			count += len(list)
		}
	}

	return count, nil
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

func getStats(books []Book) (int, int) {
	all := len(books)
	read := 0
	for _, book := range books {
		if book.Status == "done" {
			read++
		}
	}

	return all, read
}
