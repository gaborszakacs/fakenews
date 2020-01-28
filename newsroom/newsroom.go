package newsroom

import (
	"errors"

	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/real"
)

var RealReportStore *real.ReportStore

type Editor struct{}

func (n *Editor) CreateReport(tag news.Tag) error {
	feed := real.NewsFeed{}
	stories := feed.TaggedWith(tag)
	if len(stories) == 0 {
		return errors.New("No story found.")
	}

	store := RealReportStore
	store.Add(news.Report{Stories: stories})

	return nil
}
