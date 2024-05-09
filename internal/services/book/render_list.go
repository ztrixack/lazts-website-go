package book

import (
	"io"
	"strings"
)

type ListData struct {
	Items []Book
}

func (s *service) RenderList(wr io.Writer, search, catalog, status string) error {
	books, err := getList("books")
	if err != nil {
		s.log.Err(err).E("Error getting book list")
		return err
	}

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

	if err := s.templates.ExecuteTemplate(wr, "list.html", ListData{items}); err != nil {
		s.log.Err(err).E("Error executing book list template")
		return err
	}
	return nil
}
