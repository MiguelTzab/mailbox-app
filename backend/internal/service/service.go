package service

import (
	"mailbox-app/internal/entity"
	"mailbox-app/internal/service/adapter"
)

type Service struct {
	adapter adapter.SearchEngine
}

func NewService(adapter adapter.SearchEngine) *Service {
	return &Service{adapter: adapter}
}

func (svc *Service) SaveEmails(emails []entity.Email) error {
	return svc.adapter.Bulk(emails)
}
