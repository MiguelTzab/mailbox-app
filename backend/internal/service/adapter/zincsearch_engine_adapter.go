package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mailbox-app/internal/entity"
	"net/http"
	"os"
)

type ZincSearchEngineAdapter struct {
	url      string
	index    string
	username string
	password string
}

type BulkRequest struct {
	Index   string         `json:"index"`
	Records []entity.Email `json:"records"`
}

func NewZincSearchEngineAdapter() *ZincSearchEngineAdapter {
	adp := ZincSearchEngineAdapter{
		url:      "http://localhost:4080",
		index:    "emails",
		username: "admin",
		password: "atd5NF35mX£k",
	}

	if baseUrl, exists := os.LookupEnv("ZINCSEARCH_BASE_URL"); exists {
		adp.url = baseUrl
	}

	if username, exists := os.LookupEnv("ZINCSEARCH_USERNAME"); exists {
		adp.username = username
	}

	if pass, exists := os.LookupEnv("ZINCSEARCH_PASSWORD"); exists {
		adp.password = pass
	}

	if index, exists := os.LookupEnv("ZINCSEARCH_INDEX"); exists {
		adp.index = index
	}

	return &adp
}

func (zsea *ZincSearchEngineAdapter) Bulk(emails []entity.Email) error {
	request := BulkRequest{Index: zsea.index, Records: emails}
	data, err := json.Marshal(request)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/_bulkv2", zsea.url), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(zsea.username, zsea.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to index email: %s", resp.Status)
	}

	return nil
}
