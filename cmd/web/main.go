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
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	dbURL := "postgres://postgres:190704Samat@localhost:5432/snippetbox"
	flag.Parse()

}
