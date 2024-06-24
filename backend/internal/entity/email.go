package entity

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func EmailFromFile(path string) (Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return Email{}, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return Email{}, err
	}
	fileSize := fileInfo.Size()

	reader := bufio.NewReader(file)
	var email Email
	var bodyLines []string
	isBody := false
	hasRequiredHeaders := false
	lineCount := 0

	for {
		lineCount++
		if reader.Buffered() == 0 && int64(lineCount) >= fileSize {
			break
		}

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			return Email{}, err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			isBody = true
			continue
		}

		if isBody {
			bodyLines = append(bodyLines, line)
		} else {
			if strings.HasPrefix(line, "From:") {
				email.From = strings.TrimSpace(line[5:])
				hasRequiredHeaders = true
			} else if strings.HasPrefix(line, "To:") {
				email.To = strings.Split(strings.TrimSpace(line[3:]), ",")
				hasRequiredHeaders = true
			} else if strings.HasPrefix(line, "Subject:") {
				email.Subject = strings.TrimSpace(line[8:])
				hasRequiredHeaders = true
			}
		}
	}

	if !hasRequiredHeaders {
		return Email{}, fmt.Errorf("missing required headers in email: %s", path)
	}

	email.Body = strings.Join(bodyLines, "\n")
	return email, nil
}
