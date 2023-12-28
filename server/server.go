package server

import "github.com/bobykurniawan11/starter-go/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
