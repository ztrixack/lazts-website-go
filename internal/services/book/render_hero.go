package book

import (
	"io"
)

type HeroData struct {
	All       int
	Read      int
	Unread    int
	Books     []Book
	BookShelf []BookCover
}

type BookCover struct {
	Books []Book
}

func (s *service) RenderHero(wr io.Writer) error {
	books, err := getList("books")
	if err != nil {
		s.logger.Err(err).E("Error getting book list")
		return err
	}

	all, read := getStats(books)
	data := HeroData{
		All: all, Read: read,
		Unread: all - read,
		Books:  books,
		BookShelf: []BookCover{
			{Books: randomizeOne(books, 2)},
			{Books: randomizeOne(books, 3)},
			{Books: randomizeOne(books, 4)},
		},
	}

	s.logger.Fields("all", all, "read", read).I("Rendered hero books")

	if err := s.templates.ExecuteTemplate(wr, "hero.html", data); err != nil {
		s.logger.Err(err).E("Error executing hero books template")
		return err
	}
	return nil
}
