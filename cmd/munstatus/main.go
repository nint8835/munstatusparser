package main

import(
	"net/http"
	"fmt"
	"strings"
	"github.com/nint8835/munstatusparser"
)

func handler(w http.ResponseWriter, r *http.Request){
	feed := munstatusparser.GetFeed()

	var items = []string{}
	for _, item := range feed.FeedItems{
		items = append(items, item.CleanText())
	}

	fmt.Fprintf(w, strings.Join(items, "\n\n"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
