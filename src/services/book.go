package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/shuheishintani/quote-memo-api/src/models"
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

type GetBooksQuery struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   string `json:"page"`
}

func (service *Service) GetExternalBooks(getBooksInput GetBooksQuery) ([]models.Book, error) {
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
		return []models.Book{}, errors.New("parameters is not valid")
	}

	resp, err := http.Get(url)
	if err != nil {
		return []models.Book{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []models.Book{}, err
	}

	var data ApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return []models.Book{}, err
	}

	books := []models.Book{}

	for _, item := range data.Items {
		book := models.Book{
			ISBN:          item.Item.Isbn,
			Title:         item.Item.Title,
			Author:        item.Item.Author,
			Publisher:     item.Item.Publishername,
			CoverImageUrl: item.Item.Largeimageurl,
		}
		books = append(books, book)
	}

	return books, nil
}

func (service *Service) GetBooks(keyword string) ([]models.Book, error) {
	books := []models.Book{}

	if keyword == "" {
		if result := service.db.Preload("Quotes", "published IS true").Find(&books); result.Error != nil {
			return []models.Book{}, result.Error
		}
		return books, nil
	}

	if result := service.db.
		Preload("Quotes", "published IS true").
		Where("title like ?", "%"+keyword+"%").
		Or("author like ?", "%"+keyword+"%").
		Find(&books); result.Error != nil {
		return []models.Book{}, result.Error
	}
	return books, nil
}

func (service *Service) GetBookById(id string) (models.Book, error) {
	book := models.Book{}
	if result := service.db.
		Preload("Quotes", "published IS true").
		Preload("Quotes.User").
		Preload("Quotes.Tags").
		First(&book, id); result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}
