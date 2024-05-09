package page

var DEFAULT_MENU = []MenuItem{
	{
		Label: "Home",
		Path:  "/",
	},
	{
		Label: "Vacations",
		Path:  "/vacations",
	},
	{
		Label: "Books",
		Path:  "/books",
	},
	{
		Label: "Notes",
		Path:  "/notes",
	},
	{
		Label: "About",
		Path:  "/about",
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
