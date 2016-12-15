package model

/* ===================================
		RESPONSE MODEL
====================================== */
type Response struct {
	Code    int
	Status  string
	Results interface{}
}

/* ===================================
		TABLE MODEL
====================================== */

// CharacterTable will be used as a model to save the respone result
type CharacterTable struct {
	CharacterID int    `gorm:"column:CharacterID;primary_key;default:nil" json:"characterId" `
	Name        string `gorm:"column:Name;" json:"name" `
	Description string `gorm:"column:Description;" json:"description" `
	Thumbnail   string `gorm:"column:Thumbnail;" json:"thumbnail" `
	GivenID     int    `gorm:"column:GivenID;" json:"id" `
}

func (t CharacterTable) TableName() string {
	return "Code2Hire.Marvel_Characters"
}

type Characters struct {
	items []CharacterTable
}

func (cs *Characters) Populate(responseCharacters []marvelCharacter) []CharacterTable {
	var tablecharacter CharacterTable

	for _, respCharacter := range responseCharacters {
		tablecharacter = CharacterTable{}

		tablecharacter.Name = respCharacter.Name
		tablecharacter.Description = respCharacter.Description
		tablecharacter.GivenID = respCharacter.ID
		tablecharacter.Thumbnail = respCharacter.Thumbnail.Path + "." + respCharacter.Thumbnail.Extension

		cs.items = append(cs.items, tablecharacter)
	}

	return cs.items
}

// ComicTable will be used as a model to save the respone result
type ComicTable struct {
	ComicID     int    `gorm:"column:ComicID;primary_key;default:nil" json:"comicId" `
	Title       string `gorm:"column:Title;" json:"title" `
	Description string `gorm:"column:Description;" json:"description" `
	Thumbnail   string `gorm:"column:Thumbnail;" json:"thumbnail" `
	PageCount   int    `gorm:"column:PageCount;" json:"pageCount" `
	GivenID     int    `gorm:"column:GivenID;" json:"id" `
}

func (t ComicTable) TableName() string {
	return "Code2Hire.Marvel_Comics"
}

type Comics struct {
	items []ComicTable
}

func (cs *Comics) Populate(responseComics []marvelComic) []ComicTable {
	var tableComics ComicTable

	for _, respComic := range responseComics {
		tableComics = ComicTable{}

		tableComics.Title = respComic.Title
		tableComics.Description = respComic.Description
		tableComics.GivenID = respComic.ID
		tableComics.PageCount = respComic.PageCount
		tableComics.Thumbnail = respComic.Thumbnail.Path + "." + respComic.Thumbnail.Extension

		cs.items = append(cs.items, tableComics)
	}

	return cs.items
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
