package api

type (
	// PageInfo is a struct representation of the YouTube model.
	PageInfo struct {
		TotalResults   int64 `json:"totalResults"`
		ResultsPerPage int64 `json:"resultsPerPage"`
	}
)
