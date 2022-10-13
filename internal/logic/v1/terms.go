package v1

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

func (logic *Logic) GetTerms(ctx context.Context) models.Terms {
	_, tracer := logic.tracer.Start(ctx, "GetTerms", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	return logic.terms
}

func genTerms() models.Terms {
	terms := models.Terms{
		MasterVersion: viper.GetString(config.Keys.TermsMasterVersion),
		Policies:      map[string]models.TermsPolicies{},
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
