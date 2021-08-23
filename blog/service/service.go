package service

import "unit_test/blog/pkg/blog"

type Service interface {
	ListPosts() ([]blog.Post, error)
}

type service struct {
	blog blog.Blog
}

func NewService(b blog.Blog) Service {
	return &service{
		blog: b,
	}
}

func (s *service) ListPosts() ([]blog.Post, error) {
	return s.blog.ListPosts(), nil
}