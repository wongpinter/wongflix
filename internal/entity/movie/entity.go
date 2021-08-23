package movie

type Ratings struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

type Movie struct {
	ImdbID     string `json:"imdb_id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Rated      string `json:"rated"`
	Ratings    []Ratings
	Released   string `json:"released"`
	Runtime    string `json:"runtime"`
	Genre      string `json:"genre"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Awards     string `json:"awards"`
	Poster     string `json:"poster"`
	Metascore  string `json:"metascore"`
	ImdbRating string `json:"imdb_rating"`
	ImdbVotes  string `json:"imdb_votes"`
	Type       string `json:"type"`
	DVD        string `json:"dvd"`
	BoxOffice  string `json:"box_office"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"response"`
}

type Search struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchResult struct {
	Search       []Search
	TotalResults string `json:"totalResults"`
	Response     string `json:"response"`
}
