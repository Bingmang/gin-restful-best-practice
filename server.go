package main

import (
	"gin-restful-best-practice/conf"
	"gin-restful-best-practice/routes"
	"log"
)

func main() {
	config := conf.Conf()
	engine := routes.New()
	log.Fatal(engine.Run(config.URL))
}
