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

	return info, nil
}
func (m *InfoBlogsModel) Latest() ([]*InfoBlog, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
				WHERE expires > current_timestamp ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}

	infoblogs := []*InfoBlog{}

	defer rows.Close()

	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &InfoBlog{}
		// Use rows.Scan() to copy the values from each field in the row to the
		// new Snippet object that we created. Again, the arguments to row.Scan()
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Img)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		infoblogs = append(infoblogs, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return infoblogs, nil
}
