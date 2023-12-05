package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

// Estrutura para armazenar os dados de cada noticia
type MengaoNews struct {
	url, image, title, desc, date string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("ID")

	bot, err := tgbotapi.NewBotAPI(telegramToken)

	/*if err != nil {
		log.Fatal(err)
	}*/

	bot.Debug = true

	// Scraping
	var mengaoNews []MengaoNews

	var pagesToScrape []string
	pageToScrape := "https://www.flamengo.com.br/noticias/futebol?page=1"
	pagesDiscovered := []string{pageToScrape}

	// current iteration
	i := 1
	// max pages to scrape
	limit := 2

	// initializing a Colly instance
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("ul.pagination li a", func(e *colly.HTMLElement) {
		// discovering a new page
		newPaginationLink := e.Attr("href")

		// if the page discovered is new
		if !contains(pagesToScrape, newPaginationLink) {
			// if the page discovered should be scraped
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})

	// scraping the news data
	c.OnHTML("ul.list-unstyled li", func(e *colly.HTMLElement) {
		mengoPost := MengaoNews{}

		mengoPost.url = e.ChildAttr("a", "href")
		mengoPost.image = e.ChildAttr("img", "src")
		mengoPost.title = e.ChildText("h4.text-uppercase")
		mengoPost.desc = e.ChildText("p.paragraph")
		mengoPost.date = e.ChildText("span.destaque-span")

		mengaoNews = append(mengaoNews, mengoPost)

	})

	c.OnScraped(func(response *colly.Response) {
		// until there is still a page to scrape
		if len(pagesToScrape) != 0 && i < limit {
			// getting the current page to scrape and removing it from the list
			pageToScrape = pagesToScrape[0]
			pagesToScrape = pagesToScrape[1:]

			// incrementing the iteration counter
			i++

			// visiting a new page
			c.Visit(pageToScrape)
		}
	})

	// visiting the first page
	c.Visit(pageToScrape)

	for i := len(mengaoNews) - 1; i >= 0; i-- {
		mengoPost := mengaoNews[i]

		if mengoPost.url == "" || mengoPost.image == "" || mengoPost.title == "" || mengoPost.desc == "" || mengoPost.date == "" {
			continue
		}

		/*record := []string{
			mengoPost.url,
			mengoPost.image,
			mengoPost.title,
			mengoPost.desc,
			mengoPost.date,
		}

		// writing a new record
		fmt.Println(record)*/
		formatMessage := fmt.Sprintf("<a href='%s'>&#8205;</a><a href='%s'><b>%s</b></a>\n<code>%s</code>\n\n%s", 
			mengoPost.image, 
			mengoPost.url, 
			mengoPost.title, 
			mengoPost.date, 
			mengoPost.desc)

		// log.Printf("Authorized on account %s", bot.Self.UserName)

		message := tgbotapi.NewMessageToChannel(chatID, formatMessage)
		message.ParseMode = "HTML"

		_, err = bot.Send(message)
		if err != nil {
			log.Fatal(err)
		}

	}

}
