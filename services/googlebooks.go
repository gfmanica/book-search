package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"publishedDate"`
	Description   string   `json:"description"`
	PageCount     int      `json:"pageCount"`
	CoverUrl      string   `json:"coverUrl"`
}

type googleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title         string   `json:"title"`
			Authors       []string `json:"authors"`
			Publisher     string   `json:"publisher"`
			PublishedDate string   `json:"publishedDate"`
			Description   string   `json:"description"`
			PageCount     int      `json:"pageCount"`
			ImageLinks    struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func FetchBookFromGoogle(isbn string) (*Book, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s", isbn)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("erro ao consultar Google Books")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("livro não encontrado na API do Google Books")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("erro ao ler resposta da API do Google Books")
	}

	var gbResp googleBooksResponse

	if err := json.Unmarshal(body, &gbResp); err != nil {
		return nil, errors.New("erro ao decodificar resposta da API do Google Books")
	}

	if len(gbResp.Items) == 0 {
		return nil, errors.New("livro não encontrado")
	}

	info := gbResp.Items[0].VolumeInfo

	book := &Book{
		Title:         info.Title,
		Authors:       info.Authors,
		Publisher:     info.Publisher,
		PublishedDate: info.PublishedDate,
		Description:   info.Description,
		PageCount:     info.PageCount,
		CoverUrl:      info.ImageLinks.Thumbnail,
	}

	return book, nil
}
