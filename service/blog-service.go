package service

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/manuelxantony/blog/app"
	"github.com/manuelxantony/blog/blogerrors"
	"github.com/manuelxantony/blog/entity"
	"github.com/manuelxantony/blog/models"
)

type BlogService interface {
	ShowAll() ([]*entity.BlogPost, error)
	ShowById(int) (*entity.BlogPost, error)
	Create(*entity.BlogPost) (string, error)
	Update(*entity.BlogPost) error
	Delete(int) error
}

type service struct {
	app   *app.App
	model *models.BlogPosts
}

func New(app *app.App) BlogService {
	model := &models.BlogPosts{
		DB:     app.DB,
		Config: app.Config,
	}
	return &service{
		app:   app,
		model: model,
	}
}

func (s *service) ShowAll() ([]*entity.BlogPost, error) {
	posts, err := s.model.ShowAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) ShowById(id int) (*entity.BlogPost, error) {
	post, err := s.model.ShowById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, blogerrors.ErrNoRecord
		}
		return &entity.BlogPost{}, err
	}
	return post, nil
}

func (s *service) Create(entity *entity.BlogPost) (string, error) {
	s.model.Entity = entity
	id, err := s.model.Create()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(id), nil
}

func (s *service) Update(post *entity.BlogPost) error {
	s.model.Entity = post
	rowsAffected, err := s.model.Update()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return blogerrors.ErrNoRecord
	}
	return nil
}

func (s service) Delete(id int) error {
	rowsAffected, err := s.model.Delete(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return blogerrors.ErrNoRecord
	}
	return nil
}
