package main

/* ===================================
		TABLE MODEL
====================================== */

// CharacterTable will be used as a model to save the respone result
type CharacterTable struct {
	CharacterID int    `gorm:"column:CharacterID;primary_key" json:"characterId" `
	Name        string `gorm:"column:Name" json:"name" `
	Description string `gorm:"column:Description" json:"description" `
	Thumbnail   string `gorm:"column:Thumbnail" json:"thumbnail" `
	GivenID     int    `gorm:"column:GivenID" json:"id" `
}

// ComicTable will be used as a model to save the respone result
type ComicTable struct {
	ComicID     int    `gorm:"column:ComicID;primary_key" json:"comicId" `
	Title       string `gorm:"column:Title" json:"title" `
	Description string
	Thumbnail   string
	PageCount   int
	GivenID     int
}

/* ===================================
			MARVEL RESPONSE STRUCT
====================================== */

// MarvelResponseCharacter is all the data we want to parse out from the response json
type MarvelResponseCharacter struct {
	Code   int
	Status string
	Etag   string
	Data   struct {
		Offset  int
		Limit   int
		Total   int
		Count   int
		Results []marvelCharacter
	}
}

// MarvelResponseComic is all the data we want to parse out from the response json
type MarvelResponseComic struct {
	Code   int
	Status string
	Etag   string
	Data   struct {
		Offset  int
		Limit   int
		Total   int
		Count   int
		Results []marvelComic
	}
}

/* ===================================
			HELPER STRUCT
====================================== */
type marvelComic struct {
	ID          int
	Title       string
	Description string
	PageCount   int
	Thumbnail   marvelThumbnail
}

type marvelCharacter struct {
	ID          int
	Name        string
	Description string
	Thumbnail   marvelThumbnail
	Comics      struct {
		Available     int
		CollectionURI string
		Items         []marvelCharacterComic
	}
}

type marvelThumbnail struct {
	Path      string
	Extension string
}

type marvelCharacterComic struct {
	ResourceURI string
	Name        string
}
