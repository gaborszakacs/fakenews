package main

import (
	"fmt"
	"os"

	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/newsroom"
	"github.com/gaborszakacs/fakenews/real"
)

func main() {
	editor := newsroom.Editor{}
	tag := news.Tag("#climate")
	store := &real.ReportStore{}
	err := editor.CreateReport(tag, &real.NewsFeed{}, store)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	report := store.First()
	for _, story := range report.Stories {
		fmt.Println(story.Title)
	}
}
