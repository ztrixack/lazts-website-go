package templ

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type TagData struct {
	Items []TagHTML
}

func (s *service) RenderNoteTags(wr io.Writer) error {
	data, err := s.getTagList()
	if err != nil {
		s.log.Err(err).E("Error getting tag list")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "note_tags.html", data); err != nil {
		s.log.Err(err).E("Error executing note tags template")
		return err
	}
	return nil
}

func (s *service) getTagList() (TagData, error) {
	dirs, err := os.ReadDir("./contents/notes")
	if err != nil {
		fmt.Println("Error reading directories:", err)
		return TagData{}, err
	}

	tags := make([]TagHTML, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			content, err := s.markdown.ReadFile("notes", dir.Name())
			if err != nil {
				s.log.Err(err).E("Error reading file")
				continue
			}

			var note Note
			err = s.markdown.Metadata(content, &note)
			if err != nil {
				s.log.Err(err).E("Error reading metadata")
				continue
			}

			for _, notetag := range note.Tags {
				exist := false
				for _, tag := range tags {
					if tag.Name == notetag {
						tag.Count++
						exist = true
						break
					}
				}

				if !exist {
					tags = append(tags, TagHTML{
						Name:  notetag,
						Link:  fmt.Sprintf("/notes/%s", strings.ToLower(notetag)),
						Count: 1,
					})
				}
			}
		}
	}

	data := TagData{Items: tags}
	return data, nil
}
