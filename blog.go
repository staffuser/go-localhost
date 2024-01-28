package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	titles, err := GetTexts("https://wcademy.ru", "h3 > a")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Последние заголовки:")
	for _, t := range titles {
		fmt.Println("-", t)
	}
}

// GetTexts возвращает текстовое представление элементов
// по selector на странице c переданным url
func GetTexts(url, selector string) ([]string, error) {

	// Скачиваем html-страницу
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Считываем страницу в goquery-документ
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Находим все элементы с переданным селектором
	// и сохраняем их содержимое в список
	var texts []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})

	return texts, nil
