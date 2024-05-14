package vacation

import (
	"fmt"
	"lazts/internal/utils"
	"os"
)

func (s *service) getOne(name string) (*VacationHTML, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	if dirs[0].IsDir() {
		content, err := s.markdowner.ReadFile(name, dirs[0].Name())
		if err != nil {
			return nil, err
		}

		var vacation Vacation
		err = s.markdowner.Metadata(content, &vacation)
		if err != nil {
			return nil, err
		}

		result := vacation.ToHTML()
		return &result, nil
	}

	return nil, fmt.Errorf("unable to read file")
}

func (s *service) getList(name string) ([]VacationHTML, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return nil, err
	}

	vacations := make([]VacationHTML, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			content, err := s.markdowner.ReadFile(name, dir.Name())
			if err != nil {
				return nil, err
			}

			var vacation Vacation
			if err := s.markdowner.Metadata(content, &vacation); err != nil {
				return nil, err
			}

			vacations = append(vacations, vacation.ToHTML())
		}
	}

	return vacations, nil
}
