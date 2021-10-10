package main

import (
	"fmt"
	"os"
	"strconv"

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

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		num, err := strconv.Atoi(args[i])
		if num < 0 || num >= len(rsssites) {
			fmt.Println("not valid argument")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("parse: ", rsssites[num])
		rssObject, err := rss.ParseRSS(rsssites[num])
		if err == nil {

			fmt.Printf("Title           : %s\n", rssObject.Channel.Title)
			fmt.Printf("Generator       : %s\n", rssObject.Channel.Generator)
			fmt.Printf("PubDate         : %s\n", rssObject.Channel.PubDate)
			fmt.Printf("LastBuildDate   : %s\n", rssObject.Channel.LastBuildDate)
			fmt.Printf("Description     : %s\n", rssObject.Channel.Description)

			fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

			for v := range rssObject.Channel.Items {
				item := rssObject.Channel.Items[v]
				fmt.Println()
				fmt.Printf("Item Number : %d\n", v)
				fmt.Printf("Title       : %s\n", item.Title)
				fmt.Printf("Link        : %s\n", item.Link)
				fmt.Printf("Description : %s\n", item.Description)
				fmt.Printf("Guid        : %s\n", item.Guid.Value)
			}
		} else {
			fmt.Println("error: ", err)
		}
	}
}
