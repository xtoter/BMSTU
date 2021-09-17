package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола

	"github.com/IzeBerg/rss-parser-go"
)

var rsssites []string = []string{
	"http://blagnews.ru/rss_vk.xml",
	"http://www.rssboard.org/files/sample-rss-2.xml",
	"https://lenta.ru/rss",
	"https://news.mail.ru/rss/90/",
	"http://technolog.edu.ru/index.php?option=com_k2&view=itemlist&layout=category&task=category&id=8&lang=ru&format=feed",
	"https://vz.ru/rss.xml",
	"http://news.ap-pa.ru/rss.xml",
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	site := r.URL.Query().Get("site")

	fmt.Println(site)
	if site == "" {
		fmt.Fprintf(w, "<h1>Enter rcc link or select below</h1><br>")

		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=http://blagnews.ru/rss_vk.xml", "blagnews")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=http://www.rssboard.org/files/sample-rss-2.xml", "rssboard")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=https://lenta.ru/rss", "lenta")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=https://news.mail.ru/rss/90/", "news.mail.ru")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=http://technolog.edu.ru/index.php?option=com_k2&view=itemlist&layout=category&task=category&id=8&lang=ru&format=feed", "technolog.edu.ru")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=https://vz.ru/rss.xml", "vz.ru")
		fmt.Fprintf(w, "<a href=\"%s\"</a> %s <br>", "user?site=http://news.ap-pa.ru/rss.xml", "news.ap-pa.ru")

	} else {
		rssObject, err := rss.ParseRSS(site)
		fmt.Fprintf(w, "<div>")
		if err == nil {

			fmt.Fprintf(w, "<h1> %s </h1> <br>", rssObject.Channel.Title)
			for v := range rssObject.Channel.Items {
				item := rssObject.Channel.Items[v]
				fmt.Fprintf(w, "<a href=\"%s\"</a><br>", item.Link)
				fmt.Fprintf(w, "%s<br>", item.Title)

			}
		} else {
			fmt.Fprint(w, "error: ", err)
		}
		fmt.Fprintf(w, "</div>")
	}
}
func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}
