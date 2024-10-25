package mongoentities

type Link struct {
	link string `json:"link"`
}

type Dir struct {
	name    string `json:"name"`
	subdirs []Dir  `json:"subdirs"`
	links   []Link `json:"links"`
}

type User struct {
	id   string `json:"_id"`
	root Dir    `json:"root"`
}
