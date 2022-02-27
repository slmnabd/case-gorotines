package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	data, err := getSequential(123)
	if err != nil {
		log.Printf("err %v", err)
	}
	log.Printf("[SEQUENTIAL] data ID %+v", data.Book.ID)
	log.Printf("[SEQUENTIAL] response time %s", time.Since(start).String())

	start = time.Now()
	data, err = getConcurrent(123)
	if err != nil {
		log.Printf("err %v", err)
	}
	log.Printf("[CONCURRENT] data ID %+v", data.Book.ID)
	log.Printf("[CONCURRENT] response time %s", time.Since(start).String())

}

func getConcurrent(bookID int64) (data ResponseBookDetail, err error) {
	var book Book
	var rating Rating
	var bestseller []BookList

	var wg sync.WaitGroup
	var mtx sync.Mutex
	wg.Add(3)

	go func() {
		defer wg.Done()
		var errGo error
		book, errGo = GetBook(bookID)
		if errGo != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	go func() {
		defer wg.Done()
		var errGo error
		rating, errGo = GetBookRating(bookID)
		if err != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	go func() {
		defer wg.Done()
		var errGo error
		bestseller, errGo = GetBestSellerInCategory(bookID)
		if err != nil {
			mtx.Lock()
			err = errGo
			mtx.Unlock()
			return
		}
	}()

	wg.Wait()

	return ResponseBookDetail{
		Book:       book,
		Bestseller: bestseller,
		Rating:     rating,
	}, nil
}

func getSequential(bookID int64) (data ResponseBookDetail, err error) {
	book, err := GetBook(bookID)
	if err != nil {
		return data, err
	}


	bestseller, err := GetBestSellerInCategory(bookID)
	if err != nil {
		return data, err
	}

	rating, err := GetBookRating(bookID)
	if err != nil {
		return data, err
	}

	return ResponseBookDetail{
		Book:       book,
		Bestseller: bestseller,
		Rating:     rating,
	}, nil
}
