package unmarshal

import "encoding/json"

func UnmarshalAnimedown3(data []byte) (Animedown3, error) {
	var r Animedown3
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Animedown3) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Animedown3 struct {
	Data []Datum2 `json:"data"`
}

type Datum2 struct {
	ID              int64        `json:"id"`
	DynamicID       int64        `json:"dynamic_id"`
	Title           string       `json:"title"`
	EnglishTitle    *string      `json:"english_title"`
	Slug            string       `json:"slug"`
	Status          Status       `json:"status"`
	Description     string       `json:"description"`
	Year            string       `json:"year"`
	Season          string       `json:"season"`
	CoverPhoto      string       `json:"cover_photo"`
	AlternateTitles []string     `json:"alternate_titles"`
	Duration        string       `json:"duration"`
	BroadcastDay    BroadcastDay `json:"broadcast_day"`
	BroadcastTime   *string      `json:"broadcast_time"`
	Rating          Rating       `json:"rating"`
	RatingScores    float64      `json:"rating_scores"`
	GwaRating       float64      `json:"gwa_rating"`
}

type BroadcastDay string

const (
	Empty   BroadcastDay = ""
	Mondays BroadcastDay = "Mondays"
	Sundays BroadcastDay = "Sundays"
)

type Rating string

const (
	Nc17 Rating = "NC-17"
	PG13 Rating = "PG-13"
	R    Rating = "R"
)

type Status string

const (
	Finished Status = "Finished"
	OnGoing  Status = "On-going"
)
