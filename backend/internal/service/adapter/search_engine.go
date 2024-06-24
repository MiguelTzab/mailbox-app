package adapter

import "mailbox-app/internal/entity"

type SearchEngine interface {
	Bulk(emails []entity.Email) error
}
