package page

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

type Blackhole struct {
	Size    int
	Rotate  int
	Opacity int
	Width   int
}

type Cloud struct {
	Top    int
	Left   int
	Rotate int
}
