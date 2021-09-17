package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	// пакет для работы с  UTF-8 строками
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>")
	r.ParseForm()
	arg := r.URL.Query().Get("arg")
	if arg == "" {
		arg = r.FormValue("arg")
	}

	fmt.Fprintf(w, "<form method=\"POST\" action=\"\"><input type=\"text\" name=\"arg\" /><br><br><input type=\"submit\" value=\"Отправить\" /></form></body>")

	fmt.Fprintf(w, "<a href=\"%s\"</a> Вы хотели найти %s в google?<br>", "https://www.google.ru/search?q="+arg, arg)
	fmt.Fprintf(w, "<a href=\"%s\"</a> Вы хотели найти %s в duckduckgo?<br>", "https://duckduckgo.com/?q="+arg, arg)
	fmt.Fprintf(w, "<a href=\"%s\"</a> Вы хотели найти %s в yandex?<br>", "https://yandex.ru/search/?text="+arg, arg)

	//fmt.Fprintf(w, "<frameset rows=\"30%, 10%, 60%\" ><frame><frame><frame></frameset>")
	fmt.Fprintf(w, "</div>")
}
func main() {
	http.HandleFunc("/", HomeRouterHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}
