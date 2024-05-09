package book

import (
	"io"
	"lazts/internal/app/models"
)

type FilterData struct {
	CurrentCatalog string
	CurrentStatus  string
	Catalogs       models.Options
	Status         models.Options
	Size           int
}

func (s *service) RenderFilter(wr io.Writer, search, catalog, status string) error {
	books, err := getList("books")
	if err != nil {
		s.log.Err(err).E("Error getting book list")
		return err
	}

	data := FilterData{
		CurrentCatalog: catalog,
		CurrentStatus:  status,
		Catalogs:       getCatalogs(books),
		Status:         getStatus(),
		Size:           len(books),
	}
	s.log.Fields("data", data).I("Render book filter data")

	if err := s.templates.ExecuteTemplate(wr, "filter.html", data); err != nil {
		s.log.Err(err).E("Error executing book filter template")
		return err
	}
	return nil
}
