package types

type PostHogQueryResponse struct {
	IsCached    bool            `json:"is_cached"`
	HasMore     bool            `json:"hasMore"`
	LastRefresh string          `json:"last_refresh"`
	Limit       int             `json:"limit"`
	Offset      int             `json:"offset"`
	Results     [][]interface{} `json:"results"`
}
