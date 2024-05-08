package templ

import (
	"io"
	"strings"
)

type BookListData struct {
	Items []Book
}

func (s *service) RenderBookList(wr io.Writer, search, catalog, status string) error {
	books := getBookList()
	items := make([]Book, 0)

	for _, book := range books {
		if search == "" || strings.Contains(strings.ToLower(book.Title), strings.ToLower(search)) {
			if catalog == "" || book.Catalog == catalog {
				if status == "" || book.Status == status {
					items = append(items, book)
				}
			}
		}
	}

	s.log.Fields("search", search, "catalog", catalog, "status", status, "count", len(items)).I("book list")

	data := BookListData{
		Items: items,
	}

	if err := s.templates.ExecuteTemplate(wr, "book_list.html", data); err != nil {
		s.log.Err(err).E("Error executing book list template")
		return err
	}
	return nil
}

func getBookList() []Book {
	return []Book{
		{Title: "Book Test 01", Cover: "https://picsum.photos/300/200?random=1", Status: "done", Catalog: "Test01"},
		{Title: "Book 02", Cover: "https://picsum.photos/300/200?random=2", Status: "reading", Catalog: "Test01"},
		{Title: "Test 03", Cover: "https://picsum.photos/300/200?random=3", Status: "done", Catalog: "Test01"},
		{Title: "Example 04", Cover: "https://picsum.photos/300/200?random=4", Status: "reading", Catalog: "Test02"},
		{Title: "Sample 05", Cover: "https://picsum.photos/300/200?random=5", Status: "done", Catalog: "Test02"},
		{Title: "New One 06", Cover: "https://picsum.photos/300/200?random=6", Status: "wishlist", Catalog: "Test02"},
		{Title: "Recheck 07", Cover: "https://picsum.photos/300/200?random=7", Status: "done", Catalog: "Test03"},
		{Title: "Number 08", Cover: "https://picsum.photos/300/200?random=8", Status: "wishlist", Catalog: "Test03"},
		{Title: "Book Test 09", Cover: "https://picsum.photos/300/200?random=9", Status: "wishlist", Catalog: "Test03"},
	}
}
