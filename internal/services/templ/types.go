package templ

import (
	"fmt"
	"path/filepath"
	"strings"

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

func (n Vacation) ToHTML() VacationHTML {
	return VacationHTML{
		Title:    n.Title,
		Excerpt:  n.Excerpt,
		Image:    utils.UpdateFeaturedImagePaths(n.FeaturedImage, filepath.Join("/static/vacations", n.Slug)),
		ImageAlt: n.FeaturedImageAlt,
		Link:     n.Slug,
		ShowDate: fmt.Sprintf("%s - %s", n.DateFrom, n.DateTo),
		Location: n.Location,
	}
}

type Book struct {
	Number      int
	Title       string
	Subtitle    string
	Description string
	Authors     []string
	Translator  string
	Publisher   string
	Catalog     string
	Status      string
	Review      string
	Cover       string
}

type Note struct {
	Title            string
	Slug             string
	Excerpt          string
	FeaturedImage    string
	FeaturedImageAlt string
	PublishedAt      string
	Published        bool
	Tags             []string
	ReadTime         int
}

func (n Note) ToHTML() NoteHTML {
	return NoteHTML{
		Title:    n.Title,
		Excerpt:  n.Excerpt,
		Image:    utils.UpdateFeaturedImagePaths(n.FeaturedImage, filepath.Join("/static/notes", n.Slug)),
		ImageAlt: n.FeaturedImageAlt,
		Link:     filepath.Join(n.Tags[0], n.Slug),
		Tags:     ToTags(n.Tags),
		DateTime: n.PublishedAt,
		ShowTime: n.PublishedAt,
		ReadTime: n.ReadTime,
	}
}

func ToTags(tags []string) []TagHTML {
	var t []TagHTML
	for _, tag := range tags {
		t = append(t, TagHTML{
			Name:  tag,
			Link:  fmt.Sprintf("/notes/%s", strings.ToLower(tag)),
			Count: 1,
		})
	}
	return t
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

type NoteHTML struct {
	Title    string
	Excerpt  string
	Image    string
	ImageAlt string
	Link     string
	Tags     []TagHTML
	DateTime string
	ShowTime string
	ReadTime int
}

type TagHTML struct {
	Name  string
	Link  string
	Count int
}
