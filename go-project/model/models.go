package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

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

// Resonse generic model
type Resonse struct {
	Code   int
	Status string
	Etag   string
}

//Data generic model
type Data struct {
	Offset int
	limit  int
	total  int
	count  int
}

// HeroData Character specifc data model
type HeroData struct {
	Data
	Results []Character
}

//ComicData Comic specifc data model
type ComicData struct {
	Data
	Results []Comic
}

//HeroResponse Character specifc response model
type HeroResponse struct {
	Resonse
	HeroData
}

//ComicResponse Comic specifc response model
type ComicResponse struct {
	Resonse
	ComicData
}

func (t Character) TableName() string {
	return "Code2Hire.Characters"
}

func (t Character) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("LastUpdatedDateTime", time.Now())
	return nil
}

func (t Character) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("LastUpdatedDateTime", time.Now())
	return nil
}

func (t Comic) TableName() string {
	return "Code2Hire.Comics"
}

func (t Comic) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("LastUpdatedDateTime", time.Now())
	return nil
}

func (t Comic) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("LastUpdatedDateTime", time.Now())
	return nil
}
