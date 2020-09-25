package doi

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//
func GetHrefs()  {

	res, err := http.Get("https://www.worldscientific.com/doi/abs/10.1142/S0218348X95000588")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	log.Print(res.Body)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		band, ok := s.Attr("href")
		if ok {
			log.Print(band)
		}
	})

	log.Print(doc.Find("href").Text())

}
