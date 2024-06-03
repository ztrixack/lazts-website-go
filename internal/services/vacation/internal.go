package vacation

import (
	"lazts/internal/utils"
	"os"
)

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

			if !vacation.Published {
				continue
			}

			vacations = append([]VacationHTML{vacation.ToHTML()}, vacations...)
		}
	}

	return vacations, nil
}

func getCount(name string) (int, error) {
	dirs, err := os.ReadDir(utils.GetContentDir(name))
	if err != nil {
		return 0, err
	}

	count := 0
	for _, dir := range dirs {
		if dir.IsDir() {
			count++
		}
	}

	return count, nil
}
