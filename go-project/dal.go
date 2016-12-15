package main

import (
	"fmt"
	"personal-project/APIs/go-project/model"
)

func fetchAllCharacters(search string) (charactersCollection []model.CharacterTable) {

	if len(search) < 1 {
		db.Find(&charactersCollection)
		return
	}

	db.Where("Name LIKE ?", "%"+search+"%").Find(&charactersCollection)
	return

}

func fetchCharacter(ID int) (character model.CharacterTable) {
	db.Where("CharacterID = ?", ID).First(&character)
	fmt.Println(character)
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

func fetchAllComics(search string) (comicsCollection []model.ComicTable) {

	if len(search) < 1 {
		db.Find(&comicsCollection)
		return
	}

	db.Where("Title LIKE ?", "%"+search+"%").Find(&comicsCollection)
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
