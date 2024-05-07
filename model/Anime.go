package model

type Anime struct {
	Title           string `json:"title"`
	TitleInJapanese string `json:"title_japanese"`
	Duration        string `json:"duration"`
}

type Data struct {
	Data Anime `json:"data"`
}
