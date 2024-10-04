package models

type Song struct {
	Id          int    `json:"id,omitempty"`
	GroupName   string `json:"group_name"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

// TODO
type SongDetails struct {
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

type SongInput struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

// TODO
type SongResponse struct {
	UUID string
}
