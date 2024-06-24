package service

import (
	"mailbox-app/internal/entity"
	"mailbox-app/internal/service/adapter"
)

type SearchEngineService struct {
	adapter adapter.SearchEngine
}

func NewSearchEngineService(adapter adapter.SearchEngine) *SearchEngineService {
	return &SearchEngineService{adapter: adapter}
}

func (svc *SearchEngineService) SaveEmails(emails []entity.Email) error {
	return svc.adapter.SendBulkEmails(emails)
}
