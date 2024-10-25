package server

import (
	"fmt"
)

func main() {
	registerHandler("/api/get_user_data", getHandler)
	go server()
	fmt.Println("Press Ctrl+C to exit")
}
