package inits

import (
	"books/models"
	"encoding/json"
	"net/http"
)

func ImportBooks() error {
	var count int64
	result := DB.Model(&models.Book{}).Count(&count)

	if result.Error != nil {
		return result.Error
	} else if count > 0 {
		return nil
	}
	books, err := FetchBooks()
	if err != nil {
		return err
	}
	for _, book := range books {
		result := DB.Create(&book)
		if result.Error != nil {
			return err
		}
	}
	return nil
}

func FetchBooks() ([]models.Book, error) {
	url := "http://gutendex.com/books/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Results []models.Book `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Results, nil
}
