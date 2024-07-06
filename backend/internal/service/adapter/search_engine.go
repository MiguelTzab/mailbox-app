package adapter

import "mailbox-app/internal/entity"

type SearchEngine interface {
	SendBulkEmails(emails []entity.Email) error
	SearchEmails(query entity.SearchQuery) (entity.SearchResult, error)
}
