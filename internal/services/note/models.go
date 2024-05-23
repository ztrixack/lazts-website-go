package note

import (
	"fmt"
	"lazts/internal/utils"
	"path/filepath"
	"strings"
	"time"
)

type Note struct {
	Title            string
	Slug             string
	Excerpt          string
	FeaturedImage    string
	FeaturedImageAlt string
	PublishedAt      string
	LastUpdatedAt    string
	Published        bool
	Tags             []string
	ReadTime         int
}

type NoteHTML struct {
	Title         string
	Excerpt       string
	Image         string
	ImageAlt      string
	Link          string
	Tags          []TagHTML
	DateTime      string
	ShowTime      string
	ShowDateMonth string
	ShowYear      string
	ReadTime      int
}

type TagHTML struct {
	Name  string
	Link  string
	Count int
}

func (n Note) ToHTML() NoteHTML {
	publishedAt, err := time.Parse("2006-01-02", n.PublishedAt)
	if err != nil {
		publishedAt = time.Now()
	}

	return NoteHTML{
		Title:         n.Title,
		Excerpt:       n.Excerpt,
		Image:         utils.UpdateFeaturedImagePaths(n.FeaturedImage, utils.GetContentPath("notes", n.Slug)),
		ImageAlt:      n.FeaturedImageAlt,
		Link:          filepath.Join("/", "notes", n.Tags[0], n.Slug),
		Tags:          ToTags(n.Tags),
		DateTime:      publishedAt.Format(time.RFC3339),
		ShowTime:      publishedAt.Format("2016-01-02"),
		ShowDateMonth: utils.ConvertShowDayMonth(publishedAt),
		ShowYear:      fmt.Sprintf("%d", publishedAt.Year()),
		ReadTime:      n.ReadTime,
	}
}

func ToTags(tags []string) []TagHTML {
	var t []TagHTML
	for _, tag := range tags {
		t = append(t, TagHTML{
			Name:  tag,
			Link:  filepath.Join("/", "notes", strings.ToLower(tag)),
			Count: 1,
		})
	}
	return t
}

func ToBreadcrumbs(tag string) []TagHTML {
	return []TagHTML{
		{
			Name: "Home",
			Link: "/",
		},
		{
			Name: "Notes",
			Link: "/notes",
		},
		{
			Name: strings.ToUpper(string(tag[0])) + tag[1:],
			Link: filepath.Join("/", "notes", tag),
		},
	}
}
