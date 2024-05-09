package vacation

import (
	"fmt"
	"path/filepath"

	"lazts/internal/utils"
)

type Vacation struct {
	Title            string
	Slug             string
	Excerpt          string
	Location         string
	DateFrom         string
	DateTo           string
	FeaturedImage    string
	FeaturedImageAlt string
	PublishedAt      string
	Published        bool
}

type VacationHTML struct {
	Title    string
	Excerpt  string
	Image    string
	ImageAlt string
	Link     string
	ShowDate string
	Location string
}

func (v Vacation) ToHTML() VacationHTML {
	return VacationHTML{
		Title:    v.Title,
		Excerpt:  v.Excerpt,
		Image:    utils.UpdateFeaturedImagePaths(v.FeaturedImage, filepath.Join("", "static", "vacations", v.Slug)),
		ImageAlt: v.FeaturedImageAlt,
		Link:     v.Slug,
		ShowDate: fmt.Sprintf("%s - %s", v.DateFrom, v.DateTo),
		Location: v.Location,
	}
}
