package models

import "time"

type AttributesStruct struct {
	Title struct { En string } `json:"title"`
	AltTitles []struct { En string } `json:"altTitles"`
	Description struct { En string } `json:"description"`
	OriginalLanguage string `json:"originalLanguage"`
	Status string `json:"status"`
	Tags []struct {
		Type string `json:"type"`
		Attributes struct {
			Name struct { En string } `json:"name"`
			Group string `json:"Group"`
		} `json:"attributes"`
	}
	UpdatedAt time.Time `json:"updatedAt"`
	LatestUploadedChapter string `json:"latestUploadedChapter"`
}

type MangaDataStruct struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes AttributesStruct `json:"attributes"`
}
