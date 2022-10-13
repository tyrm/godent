package models

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
