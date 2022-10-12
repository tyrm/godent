package terms

import (
	"encoding/json"
	"net/http"
	"strings"

	gdhttp "github.com/tyrm/godent/internal/http"
)

// TermsGetHandler returns the terms.
func (m *Module) TermsGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "TermsGetHandler")

	terms := m.logic.GetTerms(r.Context())

	err := json.NewEncoder(w).Encode(&terms)
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}

type termsPostRequest struct {
	UserAccepts interface{} `json:"user_accepts"`
}

// TermsPostHandler .
func (m *Module) TermsPostHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "TermsPostHandler")
	ctx := r.Context()

	account, errCode, errString := m.logic.RequireAuth(r)
	if errCode != gdhttp.ErrCodeSuccess {
		gdhttp.ReturnError(w, errCode, errString)

		return
	}

	var data termsPostRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		gdhttp.ReturnError(w, gdhttp.ErrCodeBadJSON, gdhttp.ResponseMalformedJSON)

		return
	}

	var userAccepts []string
	switch obj := data.UserAccepts.(type) {
	case string:
		userAccepts = []string{obj}
	case []string:
		userAccepts = obj
	default:
		gdhttp.ReturnError(w, gdhttp.ErrCodeInvalidParam, "'user_accepts' must be a string or set of strings")

		return
	}

	terms := m.logic.GetTerms(ctx)
	unknown := unknownURLs(terms.GetURLs(), userAccepts)
	if len(unknown) > 0 {
		gdhttp.ReturnError(w, gdhttp.ErrCodeInvalidParam, "Unrecognised URLs: %s", strings.Join(unknown, ", "))

		return
	}

	err = m.logic.AddAgreedTerms(ctx, account, terms.GetURLs())
	if len(unknown) > 0 {
		l.Errorf("add agreed: %s", err.Error())
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, "Database Error")

		return
	}

	err = json.NewEncoder(w).Encode(&struct{}{})
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}

func unknownURLs(termURLs, userAccepts []string) []string {
	var unknown []string

	var found bool
	for _, userAccept := range userAccepts {
		found = false
		for _, termURL := range termURLs {
			if termURL == userAccept {
				found = true
				break
			}
		}

		if !found {
			unknown = append(unknown, userAccept)
		}
	}

	return unknown
}
