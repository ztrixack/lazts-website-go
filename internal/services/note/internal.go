package note

import (
	"lazts/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func (s *service) getList(name string) ([]NoteHTML, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	notes := make([]NoteHTML, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			content, err := s.markdown.ReadFile(name, dir.Name())
			if err != nil {
				return nil, err
			}

			var note Note
			if err := s.markdown.Metadata(content, &note); err != nil {
				return nil, err
			}

			notes = append(notes, note.ToHTML())
		}
	}

	return notes, nil
}

func (s *service) getTagList(name string) ([]TagHTML, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	tags := make([]TagHTML, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			content, err := s.markdown.ReadFile(name, dir.Name())
			if err != nil {
				return nil, err
			}

			var note Note
			if err := s.markdown.Metadata(content, &note); err != nil {
				return nil, err
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
						Link:  filepath.Join("/", name, strings.ToLower(notetag)),
						Count: 1,
					})
				}
			}
		}
	}

	return tags, nil
}
