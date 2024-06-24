package main

import (
	"flag"
	"fmt"
	"log"
	"mailbox-app/internal/entity"
	"mailbox-app/internal/service"
	"mailbox-app/internal/service/adapter"
	"mailbox-app/internal/utils"
	"os"
	"path/filepath"
)

func main() {
	destDir := flag.String("dir", "enron_data", "Directory to search Enron dataset")
	workers := flag.Int("workers", 5, "Number of concurrent workers")
	batchSize := flag.Int("batch-size", 10000, "Number of batches to send")
	profile := flag.Bool("profile", false, "write cpu and memory profile files")
	flag.Parse()

	if *profile {
		defer utils.StartCPUProfile("cpu.prof")()
		defer utils.StartMemProfile("mem.prof")()
	}

	fmt.Println("Searching emails...")
	files, err := searchEmails(*destDir)

	if err != nil {
		log.Fatal("Failed to search emails: ", err)
	}

	if len(files) == 0 {
		log.Fatal("No emails found")
	}

	fmt.Println("Emails found", len(files))
	wp := utils.NewWorkerPool(*workers, len(files))
	wp.Start()

	go func() {
		enqueueFiles(files, *wp)
	}()

	numBatches := (len(files) + *batchSize - 1) / *batchSize

	ep := utils.NewWorkerPool(*workers, numBatches)
	ep.Start()

	handleEmails(files, *batchSize, *wp, *ep)

	wp.Close()
	ep.Close()
}

func searchEmails(dir string) ([]string, error) {
	var emailPaths []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			emailPaths = append(emailPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return emailPaths, nil
}

func enqueueFiles(files []string, wp utils.WorkerPool) {
	for _, file := range files {
		wp.Submit(func(args ...interface{}) (interface{}, error) {
			f := args[0].(string)
			return entity.EmailFromFile(f)
		}, file)
	}
}

func handleEmails(files []string, batchSize int, wp utils.WorkerPool, ep utils.WorkerPool) {
	var emails = make([]entity.Email, batchSize)
	searchEngineService := service.NewSearchEngineService(adapter.NewZincSearchEngineAdapter())

	for range files {
		result := wp.GetResult()
		if result.Err != nil {
			continue
		}

		email, ok := result.Result.(entity.Email)
		if !ok {
			fmt.Println("Error: invalid result type")
			continue
		}

		emails = append(emails, email)
		if len(emails) >= batchSize {
			ep.Submit(func(args ...interface{}) (interface{}, error) {
				b := args[0].([]entity.Email)
				s := args[1].(service.SearchEngineService)

				return nil, s.SaveEmails(b)
			}, emails, *searchEngineService)

			emails = nil
		}
	}
}
