package vacation

import (
	"path/filepath"
	"time"

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
	DateTime string
	Location string
}

func (v Vacation) ToHTML() VacationHTML {
	from, err := time.Parse("2006-01-02", v.DateFrom)
	if err != nil {
		from = time.Now()
	}

	to, err := time.Parse("2006-01-02", v.DateTo)
	if err != nil {
		to = time.Now()
	}

	return VacationHTML{
		Title:    v.Title,
		Excerpt:  v.Excerpt,
		Image:    utils.UpdateFeaturedImagePaths(v.FeaturedImage, utils.GetContentPath("vacations", v.Slug)),
		ImageAlt: v.FeaturedImageAlt,
		Link:     filepath.Join("/", "vacations", v.Slug),
		ShowDate: utils.ConvertShowDate(from, to),
		DateTime: from.Format(time.RFC3339),
		Location: utils.CountryToFlagEmoji(v.Location),
	}
}
