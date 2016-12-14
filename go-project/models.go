package main

// Character will be used as a model to save the respone result
type Character struct {
	CharacterID int    `gorm:"column:CharacterID;primary_key" json:"characterId" `
	Name        string `gorm:"column:Name" json:"name" `
	Description string `gorm:"column:Description" json:"description" `
	Thumbnail   string `gorm:"column:Thumbnail" json:"thumbnail" `
	GivenID     int    `gorm:"column:GivenID" json:"id" `
}

// Comic will be used as a model to save the respone result
type Comic struct {
	ComicID     int    `gorm:"column:ComicID;primary_key" json:"comicId" `
	Title       string `gorm:"column:Title" json:"title" `
	Description string
	Thumbnail   string
	PageCount   int
	GivenID     int
}

type MarvelResponse struct {
	Code   int
	Status string
	Etag   string
	Data   struct {
		Offset  int
		Limit   int
		Total   int
		Count   int
		Results []interface{}
	}
}
