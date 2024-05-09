package book

import (
	"io"
)

type HeroData struct {
	All       int
	Read      int
	Unread    int
	Percent   float64
	Books     []Book
	BookShelf []BookCover
}

type BookCover struct {
	Books []Book
}

func (s *service) RenderHero(wr io.Writer) error {
	books, err := getList("books")
	if err != nil {
		s.log.Err(err).E("Error getting book list")
		return err
	}

	all, read, percent := getStats(books)
	data := HeroData{
		All: all, Read: read,
		Unread:  all - read,
		Percent: percent,
		Books:   books,
		BookShelf: []BookCover{
			{Books: randomizeOne(books, 2)},
			{Books: randomizeOne(books, 3)},
			{Books: randomizeOne(books, 4)},
		},
	}

	s.log.Fields("all", all, "read", read, "percent", percent).I("Rendered hero books")

	if err := s.templates.ExecuteTemplate(wr, "hero.html", data); err != nil {
		s.log.Err(err).E("Error executing hero books template")
		return err
	}
	return nil
}
