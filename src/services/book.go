package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/shuheishintani/quote-manager-api/src/dto"
)

type ApiResponse struct {
	Items []struct {
		Item struct {
			Limitedflag    int    `json:"limitedFlag"`
			Authorkana     string `json:"authorKana"`
			Author         string `json:"author"`
			Subtitle       string `json:"subTitle"`
			Seriesnamekana string `json:"seriesNameKana"`
			Title          string `json:"title"`
			Subtitlekana   string `json:"subTitleKana"`
			Itemcaption    string `json:"itemCaption"`
			Publishername  string `json:"publisherName"`
			Listprice      int    `json:"listPrice"`
			Isbn           string `json:"isbn"`
			Largeimageurl  string `json:"largeImageUrl"`
			Mediumimageurl string `json:"mediumImageUrl"`
			Titlekana      string `json:"titleKana"`
			Availability   string `json:"availability"`
			Postageflag    int    `json:"postageFlag"`
			Salesdate      string `json:"salesDate"`
			Contents       string `json:"contents"`
			Smallimageurl  string `json:"smallImageUrl"`
			Discountprice  int    `json:"discountPrice"`
			Itemprice      int    `json:"itemPrice"`
			Size           string `json:"size"`
			Booksgenreid   string `json:"booksGenreId"`
			Affiliateurl   string `json:"affiliateUrl"`
			Seriesname     string `json:"seriesName"`
			Reviewcount    int    `json:"reviewCount"`
			Reviewaverage  string `json:"reviewAverage"`
			Discountrate   int    `json:"discountRate"`
			Chirayomiurl   string `json:"chirayomiUrl"`
			Itemurl        string `json:"itemUrl"`
		} `json:"Item"`
	} `json:"Items"`
	Pagecount        int           `json:"pageCount"`
	Hits             int           `json:"hits"`
	Last             int           `json:"last"`
	Count            int           `json:"count"`
	Page             int           `json:"page"`
	Carrier          int           `json:"carrier"`
	Genreinformation []interface{} `json:"GenreInformation"`
	First            int           `json:"first"`
}

func (service *Service) GetBooks(getBooksInput dto.GetBooksInput) ([]dto.Book, error) {
	title := getBooksInput.Title
	author := getBooksInput.Author
	page := getBooksInput.Page

	url := fmt.Sprintf("%s?format=json&applicationId=%s", os.Getenv("RAKUTEN_BOOK_API_URL"), os.Getenv("RAKUTEN_APP_ID"))
	if title != "" && author != "" {
		url += fmt.Sprintf("&title=%s&author=%s&page=%s", title, author, page)
	} else if title != "" {
		url += fmt.Sprintf("&title=%s&page=%s", title, page)
	} else if author != "" {
		url += fmt.Sprintf("&&author=%s&page=%s", author, page)
	} else {
		return []dto.Book{}, errors.New("parameters is not valid")
	}

	resp, err := http.Get(url)
	if err != nil {
		return []dto.Book{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []dto.Book{}, err
	}

	var data ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return []dto.Book{}, err
	}

	books := []dto.Book{}

	for _, item := range data.Items {
		book := dto.Book{
			Isbn:          item.Item.Isbn,
			Title:         item.Item.Title,
			Author:        item.Item.Author,
			Publisher:     item.Item.Publishername,
			CoverImageUrl: item.Item.Largeimageurl,
		}
		books = append(books, book)
	}

	return books, nil
}
