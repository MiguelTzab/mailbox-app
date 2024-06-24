package main

import (
	"mailbox-app/internal/entity"
	"testing"
)

func TestSearchEmails(t *testing.T) {
	dir := "../../fixtures"

	expectedEmails := 2

	emails, err := SearchEmails(dir)
	if err != nil {
		t.Fatalf("Error parsing emails: %v", err)
	}

	if len(emails) != expectedEmails {
		t.Fatalf("Expected %d emails, got %d", expectedEmails, len(emails))
	}
}

func TestCreateBatch(t *testing.T) {
	expectedBatchSize := 2
	expectedBatches := 2
	emails := []entity.Email{
		{
			From:    "john@doe.com",
			To:      []string{"jane@doe.com"},
			Subject: "Testing",
			Body:    "I'm first test",
		},
		{
			From:    "jane@doe.com",
			To:      []string{"john@doe.com"},
			Subject: "Re: Testing",
			Body:    "Testing another email. Thank you!!!",
		},
		{
			From:    "jane@doe.com",
			To:      []string{"john@doe.com"},
			Subject: "Re: Testing 2",
			Body:    "Testing another email",
		},
	}

	batches := CreateBatch(expectedBatchSize, emails)

	if len(batches) != expectedBatches {
		t.Fatalf("Expected %d batches of emails, got %d", expectedBatches, len(batches))
	}

	for _, batch := range batches {
		if len(batch) > expectedBatchSize {
			t.Fatalf("Expected at least %d emails per batch, got %d", expectedBatchSize, len(batch))
		}
	}
}
