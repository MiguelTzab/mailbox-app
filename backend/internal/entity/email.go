package entity

import (
	"bufio"
	"fmt"
	"io"
	"mailbox-app/internal/utils"
	"net/mail"
	"os"
	"strings"
)

type Email struct {
	MessageId string   `json:"message_id"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Date      string   `json:"date"`
	Body      string   `json:"body"`
}

func EmailFromFile(path string) (Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return Email{}, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	m, err := mail.ReadMessage(r)
	if err != nil {
		return Email{}, err
	}

	header := m.Header

	var body strings.Builder
	chunkSize := 4096 // Tamaño del buffer para leer en partes
	part := make([]byte, chunkSize)

	for {
		n, err := m.Body.Read(part)
		if n > 0 {
			body.Write(part[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return Email{}, fmt.Errorf("error reading body: %w", err)
		}
	}

	email := Email{
		MessageId: header.Get("Message-ID"),
		Date:      header.Get("Date"),
		From:      header.Get("From"),
		To:        utils.SplitAndTrim(header.Get("To"), ","),
		Subject:   header.Get("Subject"),
		Body:      body.String(),
	}

	return email, nil
}
