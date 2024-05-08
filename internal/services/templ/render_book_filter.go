package templ

import (
	"io"
	"sort"
)

type Option struct {
	Key   string
	Value string
}

type BookFilterData struct {
	CurrentCatalog string
	CurrentStatus  string
	Catalogs       []Option
	Status         []Option
	Size           int
}

func (s *service) RenderBookFilter(wr io.Writer, search, catalog, status string) error {
	books := getBookList()

	data := BookFilterData{
		CurrentCatalog: catalog,
		CurrentStatus:  status,
		Catalogs:       getBookCatalogs(books),
		Status:         getBookStatus(),
		Size:           len(books),
	}
	s.log.I("Render book filter data", data)

	if err := s.templates.ExecuteTemplate(wr, "book_filter.html", data); err != nil {
		s.log.Err(err).E("Error executing book filter template")
		return err
	}
	return nil
}

func getBookCatalogs(books []Book) []Option {
	catalogs := make([]Option, 0)
	catalogs = append(catalogs, Option{Key: "ทั้งหมด", Value: ""})
	for _, book := range books {
		catalogs = appendUnique(catalogs, book.Catalog)
	}

	sort.Slice(catalogs, func(i, j int) bool {
		return catalogs[i].Key < catalogs[j].Key
	})

	return catalogs
}

func getBookStatus() []Option {
	return []Option{
		{Key: "กำลังจะซื้อ", Value: "wishlist"},
		{Key: "กำลังอ่าน", Value: "reading"},
		{Key: "อ่านจบแล้ว", Value: "done"},
		{Key: "ทั้งหมด", Value: ""},
	}
}

func appendUnique(arr []Option, data string) []Option {
	uniqueSet := make(map[string]struct{})
	result := make([]Option, len(arr))
	for i, obj := range arr {
		uniqueSet[obj.Value] = struct{}{}
		result[i] = obj
	}

	if _, exists := uniqueSet[data]; !exists {
		result = append(arr, Option{Key: data, Value: data})
	}

	return result
}
