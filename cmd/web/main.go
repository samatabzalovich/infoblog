package main

import (
	"flag"
	"infoblog/internal/models"
	"log"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	infoBlogs *models.InfoBlogsModel
}

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://postgres:1qwerty7@localhost:5432/snippetbox", "PostgresSQL data source name")
	dbURL := "postgres://postgres:1qwerty7@localhost:5432/infoblog"
	flag.Parse()

}
