package note

import (
	"lazts/internal/utils"
	"path/filepath"
	"strings"
)

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

func (n Note) ToHTML() NoteHTML {
	return NoteHTML{
		Title:    n.Title,
		Excerpt:  n.Excerpt,
		Image:    utils.UpdateFeaturedImagePaths(n.FeaturedImage, filepath.Join("", "static", "notes", n.Slug)),
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
			Link:  filepath.Join("", "notes", strings.ToLower(tag)),
			Count: 1,
		})
	}
	return t
}
