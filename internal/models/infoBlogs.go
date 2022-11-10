package models

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type InfoBlog struct {
	ID      int
	Title   string
	Content string
	Likes   int
	Img     string
	Created time.Time
}

type InfoBlogsModel struct {
	DB *pgxpool.Pool
}
