package unmarshal

import "encoding/json"

/*

	Used to unmarshal the output from a request so the main file won't be filled with these.

*/

func UnmarshalAnimedown2(data []byte) (Animedown2, error) {
	var r Animedown2
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Animedown2) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Animedown2 []Animedown2Element

type Animedown2Element struct {
	ID         string `json:"id"`
	Provider   string `json:"provider"`
	File       string `json:"file"`
	Lang       string `json:"lang"`
	Type       string `json:"type"`
	Hardsub    bool   `json:"hardsub"`
	Thumbnail  string `json:"thumbnail"`
	Resolution string `json:"resolution"`
}
