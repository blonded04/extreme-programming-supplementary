package mongoentities

import "encoding/json"

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

func SerializeUser(user *User) ([]byte, error) {
	return json.Marshal(user)
}

func DeserializeUser(data []byte) (*User, error) {
	var user User
	err := json.Unmarshal(data, &user)
	return &user, err
}
