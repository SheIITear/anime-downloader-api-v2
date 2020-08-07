package unmarshal

import "encoding/json"

/*

	Used to unmarshal the output from a request so the main file won't be filled with these.

*/

func UnmarshalAnimedown(data []byte) (Animedown, error) {
	var r Animedown
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Animedown) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Animedown struct {
	Data  []Datum `json:"data"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
	Anime Anime   `json:"anime"`
}

type Anime struct {
	ID              int64    `json:"id"`
	DynamicID       int64    `json:"dynamic_id"`
	Title           string   `json:"title"`
	EnglishTitle    string   `json:"english_title"`
	Slug            string   `json:"slug"`
	Status          string   `json:"status"`
	Description     string   `json:"description"`
	Year            string   `json:"year"`
	Season          string   `json:"season"`
	Type            string   `json:"type"`
	CoverPhoto      string   `json:"cover_photo"`
	AlternateTitles []string `json:"alternate_titles"`
	Duration        string   `json:"duration"`
	BroadcastDay    string   `json:"broadcast_day"`
	BroadcastTime   string   `json:"broadcast_time"`
	Rating          string   `json:"rating"`
	RatingScores    float64  `json:"rating_scores"`
	GwaRating       float64  `json:"gwa_rating"`
}

type Datum struct {
	ID         int    `json:"id"`
	DynamicID  int64  `json:"dynamic_id"`
	Title      string `json:"title"`
	EpisodeNum string `json:"episode_num"`
	AiringDate string `json:"airing_date"`
	Views      int64  `json:"views"`
	Sub        int    `json:"sub"`
	Dub        int    `json:"dub"`
	Thumbnail  string `json:"thumbnail"`
}

type Links struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  interface{} `json:"next"`
}

type Meta struct {
	CurrentPage int64  `json:"current_page"`
	From        int64  `json:"from"`
	LastPage    int64  `json:"last_page"`
	Path        string `json:"path"`
	PerPage     string `json:"per_page"`
	To          int64  `json:"to"`
	Total       int64  `json:"total"`
}
