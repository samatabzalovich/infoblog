package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var ctx = context.Background()

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

func (m *InfoBlogsModel) Insert(title string, content string, created int, img string) (int, error) {

	stmt := `INSERT INTO infoblogs (title, content, created, img)
	VALUES ($1, $2, current_timestamp, &3) returning id`

	var id int
	err := m.DB.QueryRow(ctx, stmt, title, content, created, img).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *InfoBlogsModel) Get(id int) (*InfoBlog, error) {

	info := &InfoBlog{}

	stmt := `SELECT id, title, content, created, img FROM infoblogs where id = $1`

	row := m.DB.QueryRow(ctx, stmt, id)

	err := row.Scan(&info.ID, &info.Title, &info.Content, &info.Created, &info.Img)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	//shortly version
	//err := m.DB.QueryRow(ctx, stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return nil, ErrNoRecord
	//	} else {
	//		return nil, err
	//	}
	//}
	return info, nil
}
