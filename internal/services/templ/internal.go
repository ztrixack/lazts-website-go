package templ

import (
	"fmt"
	"os"
)

var DEFAULT_MENU = []MenuItem{
	{
		Label: "Vacations",
		Path:  "/vacations",
		Icon:  "vacations",
	},
	{
		Label: "Books",
		Path:  "/books",
		Icon:  "books",
	},
	{
		Label: "Notes",
		Path:  "/notes",
		Icon:  "notes",
	},
}

func (s *service) getVacation() (*VacationHighlightData, error) {
	dirs, err := os.ReadDir("./contents/vacations")
	if err != nil {
		s.log.Err(err).E("Error reading directories")
		return nil, err
	}

	if dirs[0].IsDir() {
		content, err := s.markdown.ReadFile("vacations", dirs[0].Name())
		if err != nil {
			s.log.Err(err).E("Error reading file")
			return nil, err
		}

		var vacation Vacation
		err = s.markdown.Metadata(content, &vacation)
		if err != nil {
			s.log.Err(err).E("Error reading metadata")
			return nil, err
		}

		return &VacationHighlightData{
			Title:    vacation.Title,
			Excerpt:  vacation.Excerpt,
			Image:    vacation.FeaturedImage,
			Link:     vacation.Slug,
			ShowDate: fmt.Sprintf("%s - %s", vacation.DateFrom, vacation.DateTo),
			Location: vacation.Location,
		}, nil
	}

	return nil, nil
}

func (s *service) getVacationList() (*VacationData, error) {
	dirs, err := os.ReadDir("./contents/vacations")
	if err != nil {
		s.log.Err(err).E("Error reading directories")
		return nil, err
	}

	vacations := make([]VacationHTML, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			content, err := s.markdown.ReadFile("vacations", dir.Name())
			if err != nil {
				s.log.Err(err).E("Error reading file")
				continue
			}

			var vacation Vacation
			err = s.markdown.Metadata(content, &vacation)
			if err != nil {
				s.log.Err(err).E("Error reading metadata")
				continue
			}

			vacations = append(vacations, vacation.ToHTML())
		}
	}

	return &VacationData{Items: vacations}, nil
}

func (s *service) getNoteList() (NoteData, error) {
	dirs, err := os.ReadDir("./contents/notes")
	if err != nil {
		s.log.Err(err).E("Error reading directories")
		return NoteData{}, err
	}

	notes := make([]NoteHTML, 0)
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

			notes = append(notes, note.ToHTML())
		}
	}

	return NoteData{Items: notes}, nil
}
