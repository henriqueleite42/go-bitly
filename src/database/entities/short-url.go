package entities

import (
	"errors"
	"go-bitly/src/database"
	"go-bitly/src/utils"

	"github.com/google/uuid"
)

type ShortURLEntity struct {
	ID string `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Code string `json:"code" gorm:"unique,not null"`
	Clicked uint64 `json:"clicked"`
}

func GetAllShortURLs(shortURLs *[]ShortURLEntity) error {
	tx := database.Db.Find(&shortURLs)

	return tx.Error
}

func GetShortURL(shortURL *ShortURLEntity, ID string) error {
	tx := database.Db.Where("id = ?", ID).First(&shortURL)

	return tx.Error
}

func FindShortURLByCode(shortURL *ShortURLEntity, code string) error {
	tx := database.Db.Where("code = ?", code).First(&shortURL)

	return tx.Error
}

func CreateShortURL(shortURL *ShortURLEntity) error {
	shortURL.ID = uuid.New().String()

	if shortURL.Code == "" {
		shortURL.Code = utils.RandomCode()
	}

	tx := database.Db.Create(&shortURL)

	return tx.Error
}

func UpdateShortURL(shortURL *ShortURLEntity) error {
	shortURLToUpdate := ShortURLEntity{}
	database.Db.Where("id = ?", shortURL.ID).First(&shortURLToUpdate)
	if shortURLToUpdate.ID == "" {
		return errors.New("not found")
	}

	if shortURL.Code == "" {
		shortURL.Code = shortURLToUpdate.Code
	}

	if shortURL.Redirect == "" {
		shortURL.Redirect = shortURLToUpdate.Redirect
	}

	if shortURL.Clicked == 0 {
		shortURL.Clicked = shortURLToUpdate.Clicked
	}

	tx := database.Db.Save(&shortURL)

	return tx.Error
}

func DeleteShortURL(ID string) error {
	tx := database.Db.Unscoped().Delete(&ShortURLEntity{
		ID: ID,
	})

	return tx.Error
}
