package models

import "time"

type AttributesStruct struct {
	Title struct { En string } `json:"attributes.title"`
	AltTitles []struct { En string } `json:"attributes.altTitles"`
	Description struct { En string } `json:"attributes.description"`
	OriginalLanguage string `json:"attributes.originalLanguage"`
	Status string `json:"attributes.status"`
	Tags []struct {
		Type string `json:"attributes.tags.type"`
		Attributes struct {
			Name struct { En string } `json:"attributes.tags.attributes.name"`
			Group string `json:"attributes.tags.attributes.Group"`
		} `json:"attributes.tags.attributes"`
	}
	UpdatedAt time.Time `json:"attributes.updatedAt"`
	LatestUploadedChapter string `json:"attributes.latestUploadedChapter"`
}

type MangaDataStruct struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes AttributesStruct `json:"attributes"`
}
