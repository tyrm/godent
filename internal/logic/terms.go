package logic

import (
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
)

type Terms struct {
	Policies TermPolicies `json:"policies"`
}

type TermPolicies struct {
	PrivacyPolicy  map[string]interface{} `json:"privacy_policy,omitempty"`
	TermsOfService map[string]interface{} `json:"terms_of_service,omitempty"`
}

type TermsLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (*Logic) GetTerms() Terms {
	terms := Terms{}

	privacyVersion := viper.GetString(config.Keys.PrivacyVersion)
	if privacyVersion != "" {
		privacyURLs := viper.GetStringMap(config.Keys.PrivacyURLs)
		privacyURLs["version"] = privacyVersion
		terms.Policies.PrivacyPolicy = privacyURLs
	}

	termsVersion := viper.GetString(config.Keys.TermsVersion)
	if termsVersion != "" {
		termsURLs := viper.GetStringMap(config.Keys.TermsURLs)
		termsURLs["version"] = termsVersion
		terms.Policies.TermsOfService = termsURLs
	}

	return terms
}
