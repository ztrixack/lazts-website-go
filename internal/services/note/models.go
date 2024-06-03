package note

import (
	"fmt"
	"lazts/internal/utils"
	"path/filepath"
	"strings"
	"time"
)

type Note struct {
	Title         string   `json:"title"`
	Slug          string   `json:"slug"`
	Excerpt       string   `json:"excerpt"`
	FeaturedImage string   `json:"featured_image"`
	PublishedAt   string   `json:"published_at"`
	LastUpdatedAt string   `json:"last_updated_at"`
	Published     bool     `json:"published"`
	Tags          []string `json:"tags"`
	ReadTime      int      `json:"-"`
}

type NoteHTML struct {
	Title            string
	Excerpt          string
	FeaturedImage    string
	Link             string
	Tags             []TagHTML
	ReadTime         int
	DateTimeISO      string
	DateTimeReadable string
	DayMonth         string
	Year             string
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
		Title:            n.Title,
		Excerpt:          n.Excerpt,
		FeaturedImage:    utils.UpdateFeaturedImagePaths(n.FeaturedImage, utils.GetContentPath("notes", n.Slug)),
		Link:             filepath.Join("/", "notes", n.Tags[0], n.Slug),
		Tags:             ToTags(n.Tags),
		ReadTime:         n.ReadTime,
		DateTimeISO:      publishedAt.Format(time.RFC3339),
		DateTimeReadable: utils.ToYearMonthDay(publishedAt),
		DayMonth:         utils.ToDayMonth(publishedAt),
		Year:             fmt.Sprintf("%d", publishedAt.Year()),
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
