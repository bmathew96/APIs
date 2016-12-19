package main

import (
	"personal-project/APIs/go-project/model"
	"strings"
)

func fetchAllCharacters(search string, limit int) (charactersCollection []model.CharacterTable, err error) {
	search = strings.Replace(search, "\"", "", -1)

	if limit < 1 {
		limit = 100
	}

	if len(search) < 1 {
		db.Limit(limit).Find(&charactersCollection)
		return
	}

	db.Where("Name LIKE ?", "%"+search+"%").Limit(limit).Find(&charactersCollection)
	return

}

func fetchCharacter(ID int) (character model.CharacterTable) {

	db.Where("CharacterID = ?", ID).First(&character)
	return
}

func deleteCharacter(ID int) {
	character := fetchCharacter(ID)
	db.Delete(&character)
}

func addUpdateCharacter(request model.CharacterTable) (response model.CharacterTable) {
	if request.CharacterID > 1 {

		character := fetchCharacter(request.CharacterID)
		db.First(&character)

		// we can use reflect to loop over these
		if len(request.Name) > 0 {
			character.Name = request.Name
		}
		if len(request.Description) > 0 {
			character.Description = request.Description
		}
		if len(request.Thumbnail) > 0 {
			character.Thumbnail = request.Thumbnail
		}

		if request.GivenID > 0 {
			character.GivenID = request.GivenID
		}

		db.Save(&character)

		response = character
	}

	if db.NewRecord(request) {

		db.Create(&request)
		response = request
	}

	return
}

func fetchAllComics(search string, limit int) (comicsCollection []model.ComicTable) {
	search = strings.Replace(search, "\"", "", -1)
	if limit < 1 {
		limit = 100
	}

	if len(search) < 1 {
		db.Limit(limit).Find(&comicsCollection)
		return
	}

	db.Where("Title LIKE ?", "%"+search+"%").Limit(limit).Find(&comicsCollection)
	return

}

func fetchComic(ID int) (comic model.ComicTable) {

	db.Where("ComicID = ?", ID).First(&comic)
	return
}

func deleteComic(ID int) {

	comic := fetchComic(ID)
	db.Delete(&comic)
}

func addUpdateComic(request model.ComicTable) (response model.ComicTable) {
	if request.ComicID > 1 {

		comic := fetchComic(request.ComicID)
		db.First(&comic)

		// we can use reflect to loop over these
		if len(request.Title) > 0 {
			comic.Title = request.Title
		}
		if len(request.Description) > 0 {
			comic.Description = request.Description
		}
		if len(request.Thumbnail) > 0 {
			comic.Thumbnail = request.Thumbnail
		}

		if request.GivenID > 0 {
			comic.GivenID = request.GivenID
		}

		if comic.PageCount != request.PageCount && request.PageCount > 0 {
			comic.PageCount = request.PageCount
		}

		db.Save(&comic)

		response = comic
	}

	if db.NewRecord(request) {

		db.Create(&request)
		response = request
	}

	return
}
