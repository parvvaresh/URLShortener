package service

import (
	"errors"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService(repo : repo)
}

func (s *URLService) Shorten(original string) (string, string) {
	shortCode := utils.GenerateCode(6)
	err := s.repo.Save(shortCode, original)

	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func (s *URLService) Resolve(code string) (string, error) {
	url, err := s.repo.Get(code)
	if err != nil {
		return "", errors.New("Not Found")
	}
	return url, nil
}
