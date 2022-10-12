package logic

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/models"
	"go.opentelemetry.io/otel/trace"
)

func (logic *Logic) AddAgreedTerms(ctx context.Context, account *models.Account, urls []string) error {
	_, tracer := logic.tracer.Start(ctx, "AddAgreedTerms", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	var acceptedTermsURLs []*models.AcceptedTermsURL
	for _, u := range urls {
		acceptedTerm := models.AcceptedTermsURL{
			Account:   account,
			AccountID: account.ID,
			URL:       u,
		}

		acceptedTermsURLs = append(acceptedTermsURLs, &acceptedTerm)
	}

	if err := logic.db.CreateAcceptedTermsURL(ctx, acceptedTermsURLs...); err != nil {
		return fmt.Errorf("db create: %s", err)
	}

	acceptedTermsURLs, err := logic.db.ReadAcceptedTermsURLForAccount(ctx, account.ID)
	if err != nil {
		return fmt.Errorf("db read: %s", err)
	}

	var allAccepted []string
	for _, a := range acceptedTermsURLs {
		urls = append(urls, a.URL)
	}

	terms := logic.GetTerms(ctx)
	if terms.IsFullyAgreed(allAccepted) {
		account.ConsentVersion = terms.MasterVersion

		if err := logic.db.UpdateAccount(ctx, account); err != nil {
			return fmt.Errorf("db update: %s", err)
		}
	}

	return nil
}

/*func (logic *Logic) GetAgreedTerms(ctx context.Context, account *models.Account) ([]string, error) {
	_, tracer := logic.tracer.Start(ctx, "GetAgreedTerms", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	acceptedTermsURLs, err := logic.db.ReadAcceptedTermsURLForAccount(ctx, account.ID)
	if err != nil {
		return nil, fmt.Errorf("db read: %s", err)
	}

	var urls []string
	for _, a := range acceptedTermsURLs {
		urls = append(urls, a.URL)
	}

	return urls, nil
}*/

func (logic *Logic) GetTerms(ctx context.Context) Terms {
	_, tracer := logic.tracer.Start(ctx, "GetTerms", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	terms := Terms{
		MasterVersion: viper.GetString(config.Keys.TermsMasterVersion),
		Policies:      map[string]TermsPolicies{},
	}

	privacyVersion := viper.GetString(config.Keys.PrivacyVersion)
	if privacyVersion != "" {
		privacyURLs := viper.GetStringMap(config.Keys.PrivacyURLs)
		privacyURLs["version"] = privacyVersion
		terms.Policies["privacy_policy"] = privacyURLs
	}

	termsVersion := viper.GetString(config.Keys.TermsVersion)
	if termsVersion != "" {
		termsURLs := viper.GetStringMap(config.Keys.TermsURLs)
		termsURLs["version"] = termsVersion
		terms.Policies["terms_of_service"] = termsURLs
	}

	return terms
}

type TermsPolicies map[string]interface{}

func (tp TermsPolicies) HasURL(url string) bool {
	for _, doc := range tp {
		d, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		u, ok := d["url"].(string)
		if !ok {
			continue
		}

		if u == url {
			return true
		}
	}

	return false
}

type Terms struct {
	MasterVersion string                   `json:"-"`
	Policies      map[string]TermsPolicies `json:"policies"`
}

func (t *Terms) GetMasterVersion() (string, error) {
	if t.MasterVersion == "" {
		return "", ErrNotFound
	}

	return t.MasterVersion, nil
}

func (t *Terms) GetURLs() []string {
	var urls []string
	for _, policy := range t.Policies {
		for _, link := range policy {
			l, ok := link.(map[string]interface{})
			if !ok {
				continue
			}

			u, ok := l["url"].(string)
			if !ok {
				continue
			}

			urls = append(urls, u)
		}
	}

	return urls
}

func (t *Terms) IsFullyAgreed(userAccepts []string) bool {
	var found bool
	for _, p := range t.Policies {
		found = false
		for _, u := range userAccepts {
			if p.HasURL(u) {
				found = true
				continue
			}
		}

		if !found {
			return false
		}
	}

	return true
}
