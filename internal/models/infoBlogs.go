package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
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

type Comment struct {
	ID       int
	Username string
	Created  time.Time
	text     string
}

type InfoBlogsModel struct {
	DB *pgxpool.Pool
}

func (m *InfoBlogsModel) Insert(title string, content string, created int, img string) (int, error) {

	stmt := `INSERT INTO infoblogs (title, content, created, img) VALUES ($1, $2, current_timestamp, &3) returning id`

	var id int
	err := m.DB.QueryRow(ctx, stmt, title, content, created, img).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *InfoBlogsModel) Get(id int) (*InfoBlog, error) {

	info := &InfoBlog{}

	stmt := "SELECT id, title, content, created, img FROM infoblogs where id = $1"

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
	//  if errors.Is(err, sql.ErrNoRows) {
	//    return nil, ErrNoRecord
	//  } else {
	//    return nil, err
	//  }
	//}
	return info, nil
}

func (m *InfoBlogsModel) GetPopular() ([]*InfoBlog, error) {
	stmt := `SELECT * FROM infoblogs s order by likes desc limit 10`
	result, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
	}
	defer result.Close()
	var infoblogs []*InfoBlog
	for result.Next() {
		s := &InfoBlog{}
		err := result.Scan(&s.ID, &s.Title, &s.Content, &s.Likes, &s.Created, &s.Img)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNoRecord
			}
			return nil, err
		}
		infoblogs = append(infoblogs, s)
	}
	return infoblogs, nil
}

func (m *InfoBlogsModel) Latest() ([]*InfoBlog, error) {
	stmt := `SELECT * FROM infoblogs order by created desc limit 10`
	result, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
	}
	defer result.Close()
	var infoblog []*InfoBlog
	for result.Next() {
		s := &InfoBlog{}
		err := result.Scan(&s.ID, &s.Title, &s.Content, &s.Likes, &s.Created, &s.Img)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNoRecord
			}
			return nil, err
		}
		infoblog = append(infoblog, s)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return infoblog, nil
}

func (m *InfoBlogsModel) IsLiked(blog_id int, user_id int) (liked bool, err error) {
	stmt := `Select exists(SELECT * FROM likes where blog_id = $1 AND user_id = $2)`
	var result bool
	err = m.DB.QueryRow(ctx, stmt, blog_id, user_id).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, ErrNoRecord
		} else {
			return false, err
		}
	}
	return result, nil
}

func (m *InfoBlogsModel) ToLike(blog_id int, user_id int) (err error) {
	stmt := `Select exists(SELECT * FROM likes where blog_id = $1 AND user_id = $2)`
	var result bool
	err = m.DB.QueryRow(ctx, stmt, blog_id, user_id).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNoRecord
		} else {
			return err
		}
	}
	if result == false {
		stmt = `Insert into likes(blog_id, user_id) values($1, $2)`
		var pg pgconn.CommandTag
		pg, err = m.DB.Exec(ctx, stmt, blog_id, user_id)
		if err != nil {
			return err
		}
		pg.Insert()
		stmt = `Update infoblogs set likes = likes + 1 where blog_id = $1`
		pg, err = m.DB.Exec(ctx, stmt, blog_id)
		if err != nil {
			return err
		}
		pg.Update()
	} else {
		stmt = `Delete from likes where blog_id = $1 AND user_id = $2`
		var pg pgconn.CommandTag
		pg, err = m.DB.Exec(ctx, stmt, blog_id, user_id)
		if err != nil {
			return err
		}
		pg.Delete()
		stmt = `Update infoblogs set likes = likes - 1 where blog_id = $1`
		pg, err = m.DB.Exec(ctx, stmt, blog_id)
		if err != nil {
			return err
		}
		pg.Update()
	}
	return nil
}

func (m *InfoBlogsModel) PostComment(user_id int, blog_id int, text string) (bool, error) {
	stmt := `INSERT INTO comments (blog_id, user_id, text, created) VALUES ($1, $2, $3, current_timestamp)`

	pg, err := m.DB.Exec(ctx, stmt, blog_id, user_id, text)
	res := pg.Insert()
	if err != nil {
		return false, err
	}

	return res, nil
}

func (m *InfoBlogsModel) GetComments(blog_id int) ([]*Comment, error) {

	stmt := "SELECT c.comment_id, u.name, c.created, c.text FROM comments c join users u on c.user_id = u.id where blog_id = $1 order by c.created desc"

	result, err := m.DB.Query(ctx, stmt, blog_id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	var comments []*Comment
	for result.Next() {
		s := &Comment{}
		err := result.Scan(&s.ID, &s.Username, &s.Created, &s.text)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNoRecord
			}
			return nil, err
		}
		comments = append(comments, s)
	}
	return comments, nil

}
