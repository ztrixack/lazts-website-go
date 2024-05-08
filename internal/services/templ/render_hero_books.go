package templ

import (
	"io"
	"math/rand/v2"
)

type BookCover struct {
	Books []Book
}

type BookData struct {
	All       int
	Read      int
	Unread    int
	Percent   float64
	Books     []Book
	BookShelf []BookCover
}

func (s *service) RenderHeroBooks(wr io.Writer) error {
	books := getBooks()
	all, read, percent := stats(books)
	data := BookData{
		All: all, Read: read,
		Unread:  all - read,
		Percent: percent,
		Books:   books,
		BookShelf: []BookCover{
			{Books: getRandomBook(books, 2)},
			{Books: getRandomBook(books, 3)},
			{Books: getRandomBook(books, 4)},
		},
	}

	if err := s.templates.ExecuteTemplate(wr, "hero_books.html", data); err != nil {
		s.log.Err(err).E("Error executing hero books template")
		return err
	}
	return nil
}

func getRandomBook(books []Book, count int) []Book {
	result := make([]Book, count)

	for i := range result {
		r := rand.IntN(len(books))
		result[i] = books[r]
	}

	return result
}

func getBooks() []Book {
	// Example data
	return []Book{
		{Cover: "https://picsum.photos/300/200?random=1", Status: "completed"},
		{Cover: "https://picsum.photos/300/200?random=2", Status: "pending"},
		{Cover: "https://picsum.photos/300/200?random=3", Status: "completed"},
		{Cover: "https://picsum.photos/300/200?random=4", Status: "pending"},
		{Cover: "https://picsum.photos/300/200?random=5", Status: "completed"},
		{Cover: "https://picsum.photos/300/200?random=6", Status: "pending"},
		{Cover: "https://picsum.photos/300/200?random=7", Status: "completed"},
		{Cover: "https://picsum.photos/300/200?random=8", Status: "pending"},
		{Cover: "https://picsum.photos/300/200?random=9", Status: "pending"},
	}
}

func stats(books []Book) (int, int, float64) {
	all := len(books)
	read := 0
	for _, book := range books {
		if book.Status == "completed" {
			read++
		}
	}
	percent := float64(read) * 100 / float64(all)
	return all, read, percent
}
