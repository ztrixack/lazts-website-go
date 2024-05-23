package note

import (
	"io"
	"time"
)

type HeaderData struct {
	Title           string
	PublishedDate   string
	Published       string
	LastUpdatedDate string
	LastUpdated     string
	ReadTime        int
	Tags            []TagHTML
	Breadcrumbs     []TagHTML
}

func (s *service) RenderHeader(wr io.Writer, name string) error {
	item, err := s.getOne("notes", name)
	if err != nil {
		s.logger.Err(err).E("Error getting notes")
		return err
	}

	publishedAt, err := time.Parse("2006-01-02", item.PublishedAt)
	if err != nil {
		publishedAt = time.Now()
	}

	lastUpdatedAt, err := time.Parse("2006-01-02", item.LastUpdatedAt)
	if err != nil {
		lastUpdatedAt = publishedAt
	}

	data := HeaderData{
		Title:           item.Title,
		PublishedDate:   publishedAt.Format(time.RFC3339),
		Published:       publishedAt.Format("2016-01-02"),
		LastUpdatedDate: lastUpdatedAt.Format(time.RFC3339),
		LastUpdated:     lastUpdatedAt.Format("2016-01-02"),
		ReadTime:        item.ReadTime,
		Tags:            ToTags(item.Tags),
		Breadcrumbs:     ToBreadcrumbs(item.Tags[0]),
	}

	if err := s.templates.ExecuteTemplate(wr, "header.html", data); err != nil {
		s.logger.Err(err).E("Error executing note header template")
		return err
	}

	return nil
}
