package main

import (
	"flag"
	"fmt"
	"log"
	"mailbox-app/internal/entity"
	"mailbox-app/internal/service"
	"mailbox-app/internal/service/adapter"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	destDir := flag.String("dir", "enron_data", "Directory to search Enron dataset")
	batchSize := flag.Int("batch", 5000, "Number of batches to send")
	flag.Parse()

	fmt.Println("Searching emails...")
	emails, err := SearchEmails(*destDir)
	fmt.Println("Emails found", len(emails))

	if err != nil {
		log.Fatalf("Failed to parse emails: %v", err)
	}

	indexer := service.NewService(adapter.NewZincSearchEngineAdapter())
	for _, batch := range CreateBatch(*batchSize, emails) {
		if err := indexer.SaveEmails(batch); err != nil {
			log.Printf("Failed to export emails: %v", err)
		}
	}
	fmt.Println("Exporting completed.")
}

func SearchEmails(dir string) ([]entity.Email, error) {
	var emails []entity.Email
	var wg sync.WaitGroup
	emailChan := make(chan entity.Email)
	errChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		for email := range emailChan {
			emails = append(emails, email)
		}
		doneChan <- struct{}{}
	}()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				email, err := entity.EmailFromFile(path)
				if err != nil {
					errChan <- fmt.Errorf("error parsing email %s: %v", path, err)
					return
				}
				emailChan <- email
			}(path)
		}
		return nil
	})

	go func() {
		wg.Wait()
		close(emailChan)
	}()

	go func() {
		for err := range errChan {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	<-doneChan
	return emails, err
}

func CreateBatch(batchSize int, emails []entity.Email) [][]entity.Email {
	batches := make([][]entity.Email, 0, (len(emails)+batchSize-1)/batchSize)

	for batchSize < len(emails) {
		emails, batches = emails[batchSize:], append(batches, emails[0:batchSize:batchSize])
	}
	batches = append(batches, emails)

	return batches
}
