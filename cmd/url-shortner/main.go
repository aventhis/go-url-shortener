package main

import (
	"github.com/aventhis/go-url-shortener/internal/config"
)

func main() {
	// init config
	cfg := config.MustLoad()
	//fmt.Println(cfg)
	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server:

}
