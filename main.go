package main

import (
	"fmt"
	"os"

	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/newsroom"
	"github.com/gaborszakacs/fakenews/real"
)

var RealReportStore = &real.ReportStore{}

func main() {
	newsroom.RealReportStore = RealReportStore
	editor := newsroom.Editor{}
	tag := news.Tag("#climate")
	err := editor.CreateReport(tag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	report := RealReportStore.First()
	for _, story := range report.Stories {
		fmt.Println(story.Title)
	}
}
