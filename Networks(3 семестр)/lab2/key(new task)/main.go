package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Item struct {
	time, data string
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func downloadNews() []*Item {
	log.Println("sending request to www.maineantiquedigest.com")
	if response, err := http.Get("https://www.maineantiquedigest.com/all-shows?start=0&end=34"); err != nil {
		log.Println("request to www.maineantiquedigest.com failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Println("got response from www.maineantiquedigest.com", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Println("invalid HTML from www.maineantiquedigest.com", "error", err)
			} else {
				log.Println("HTML from www.maineantiquedigest.com parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}
func curclass(node *html.Node, str string) bool {
	return strings.Contains(fmt.Sprint(node), str)
}
func search(node *html.Node) []*Item {
	var res []*Item
	//fmt.Println(node)
	if curclass(node, "interior") {
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {

			if curclass(c, "wrapper") {

				for c1 := c.FirstChild; c1 != nil; c1 = c1.NextSibling {

					if curclass(c1, "page") {
						for c2 := c1.FirstChild; c2 != nil; c2 = c2.NextSibling {
							if curclass(c2, "live") {
								for c3 := c2.FirstChild; c3 != nil; c3 = c3.NextSibling {
									if curclass(c3, "left-column") {
										for c4 := c3.FirstChild; c4 != nil; c4 = c4.NextSibling {
											if curclass(c4, "search-results") {
												for c5 := c4.FirstChild; c5 != nil; c5 = c5.NextSibling {

													for c6 := c5.FirstChild; c6 != nil; c6 = c6.NextSibling {
														var cur Item
														if curclass(c6, "column date") {
															cur.time = c6.FirstChild.Data
														} else {
															if curclass(c6, "column") {
																cur.data = c6.FirstChild.FirstChild.Data
															}
														}
														items = append(items, &cur)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		res = append(res, items...)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			res = append(res, items...)
		}
	}

	return res
}

func HomeRouterHandler(w http.ResponseWriter, _ *http.Request) {

	items := downloadNews()
	fmt.Fprintf(w, "<h1>maineantiquedigest</h1><br>")
	for i := 0; i < len(items); i++ {
		fmt.Fprintf(w, "%s %s <br>", items[i].time, items[i].data)
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
