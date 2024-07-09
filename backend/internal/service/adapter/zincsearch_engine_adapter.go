package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type SearchRequest struct {
	SearchType string         `json:"search_type"`
	Query      ZincSeachQuery `json:"query"`
	SortFields []string       `json:"sort_fields"`
	From       int            `json:"from"`
	MaxResults int            `json:"max_results"`
	Source     []string       `json:"_source"`
}

type ZincSeachQuery struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}

type ZincSearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []struct {
			Email entity.Email `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

const search_type_match = "match"
const search_type_all = "alldocuments"

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

func (zsea *ZincSearchEngineAdapter) SendBulkEmails(emails []entity.Email) error {
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
		return fmt.Errorf("failed to create emails: %s", resp.Status)
	}

	return nil
}

func (zsea *ZincSearchEngineAdapter) SearchEmails(query entity.SearchQuery) (entity.SearchResult, error) {
	data, err := buildSearchRequest(query)
	if err != nil {
		return entity.SearchResult{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/%s/_search", zsea.url, zsea.index), bytes.NewBuffer(data))
	if err != nil {
		return entity.SearchResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(zsea.username, zsea.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entity.SearchResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.SearchResult{}, fmt.Errorf("failed to search emails: %s", resp.Status)
	}

	return buildSearchResult(resp.Body, query)
}

func buildSearchRequest(query entity.SearchQuery) ([]byte, error) {
	zincQuery := SearchRequest{
		SortFields: []string{"-@timestamp"},
		From:       query.Page * query.PageSize,
		MaxResults: query.PageSize,
	}

	if query.Term != "" {
		zincQuery.Query = ZincSeachQuery{Term: query.Term, Field: "_all"}
		zincQuery.SearchType = search_type_match
	} else {
		zincQuery.SearchType = search_type_all
	}

	if len(query.Fields) != 0 {
		for _, field := range query.Fields {
			zincQuery.Source = append(zincQuery.Source, field)
		}
	}

	requestBody, err := json.Marshal(zincQuery)
	if err != nil {
		return []byte{}, err
	}

	return requestBody, nil
}

func buildSearchResult(body io.ReadCloser, query entity.SearchQuery) (entity.SearchResult, error) {
	response := ZincSearchResponse{}
	result := entity.SearchResult{Items: make([]entity.Email, 0)}
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return entity.SearchResult{}, err
	}

	for _, record := range response.Hits.Hits {
		result.Items = append(result.Items, record.Email)
	}

	result.TotalResults = response.Hits.Total.Value
	result.Page = query.Page
	result.PageSize = query.PageSize

	return result, nil
}
