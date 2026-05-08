package internet

type SearchResult struct {
	Title string
	URL   string
}

func Search(query string) []SearchResult {

	results := []SearchResult{

		{
			Title: query,
			URL:   "https://en.wikipedia.org/wiki/" + query,
		},
	}

	return results
}
