package main

import (
	"testing"
)

func TestSearchEmails(t *testing.T) {
	dir := "../../fixtures"

	expectedEmails := 2

	emails, err := searchEmails(dir)
	if err != nil {
		t.Fatalf("Error parsing emails: %v", err)
	}

	if len(emails) != expectedEmails {
		t.Fatalf("Expected %d emails, got %d", expectedEmails, len(emails))
	}
}
