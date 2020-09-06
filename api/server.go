package api

import (
	"github.com/Ilhamsat/go/api/controllers"
)

var server = controllers.Server{}

func Run() {

	server.Initialize()

	server.Run(":8090")

}
