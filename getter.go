package main

import (
"encoding/json"
"fmt"
"io"
"net/http"
)

var (
	client  = &http.Client{}
	baseURL = "toko-buku-mania.herokuapp.com"
)

// GetBook get book detail
func GetBook(bookID int64) (data Book, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("https://%s/books/%d", baseURL, bookID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}

// GetBestSellerInCategory recommendation
func GetBestSellerInCategory(bookID int64) (data []BookList, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("https://%s/books/%d/bestseller", baseURL, bookID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}

// GetBookRating get rating from third party
func GetBookRating(bookID int64) (data Rating, err error) {
	responseAPI, err := http.Get(fmt.Sprintf("https://%s/books/%d/rating", baseURL, bookID))
	if err != nil {
		return data, err
	}
	defer responseAPI.Body.Close()
	byteBody, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		return data, err
	}
	if responseAPI.StatusCode != http.StatusOK {
		return data, fmt.Errorf("got status %d", responseAPI.StatusCode)
	}

	return data, json.Unmarshal(byteBody, &data)
}