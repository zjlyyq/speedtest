package main

import (
	"backend/config"
	"backend/database"
	"backend/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.Load()

	database.SetDBInfo(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}
