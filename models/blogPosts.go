package models

import (
	"database/sql"
	"strconv"

	"github.com/manuelxantony/blog/config"
	"github.com/manuelxantony/blog/entity"
)

type BlogPosts struct {
	Entity *entity.BlogPost
	DB     *sql.DB
	Config *config.Configuration
}

func (b *BlogPosts) ShowAll() ([]*entity.BlogPost, error) {
	stmt := `SELECT id, title, content, createdby FROM post`

	rows, err := b.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*entity.BlogPost{}

	for rows.Next() {
		post := &entity.BlogPost{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedBy)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (b *BlogPosts) ShowById(id int) (*entity.BlogPost, error) {
	stmt := `SELECT id, title, content, createdby FROM post WHERE id=?`

	rows := b.DB.QueryRow(stmt, id)

	post := &entity.BlogPost{}
	err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedBy)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (b *BlogPosts) Create() (int, error) {
	stmt := `INSERT INTO post (title, content, createdby)
			  VALUES(?, ?, ?)`

	result, err := b.DB.Exec(stmt, b.Entity.Title, b.Entity.Content, b.Entity.CreatedBy)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (b *BlogPosts) Update() (int, error) {
	stmt := `UPDATE post set title=?,content=? WHERE id=?`
	id_int, err := strconv.Atoi(b.Entity.ID)
	if err != nil {
		return -1, nil
	}
	result, err := b.DB.Exec(stmt, b.Entity.Title, b.Entity.Content, id_int)
	if err != nil {
		return -1, err
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(rowsEffected), err
}

func (b *BlogPosts) Delete(id int) (int, error) {
	stmt := `DELETE FROM post WHERE id=?`

	result, err := b.DB.Exec(stmt, id)
	if err != nil {
		return -1, err
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(rowsEffected), nil
}
