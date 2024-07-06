package handler

import (
	"encoding/json"
	"fmt"
	"mailbox-app/internal/entity"
	"mailbox-app/internal/service"
	"net/http"
	"strconv"
)

type EmailHandler struct {
	SearchEngineService service.SearchEngineService
}

func (eh *EmailHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	term := r.URL.Query().Get("q")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 0
	} else {
		page = page - 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 20
	}

	query := entity.SearchQuery{
		Term:     term,
		Page:     page,
		PageSize: limit,
		Fields:   make(map[string]string),
	}

	fields := r.URL.Query()["cols[]"]

	if len(fields) > 0 {
		for _, field := range fields {
			query.Fields[field] = field
		}
	}

	results, err := eh.SearchEngineService.GetEmails(query)

	if err != nil {
		fmt.Println("[handler] failed to search:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	results.Page = results.Page + 1
	if err := json.NewEncoder(w).Encode(results); err != nil {
		fmt.Println("[handler] failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
