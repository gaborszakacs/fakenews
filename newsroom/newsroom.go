package newsroom

import (
	"errors"

	"github.com/gaborszakacs/fakenews/news"
)

type Editor struct{}

type NewsFeed interface {
	TaggedWith(tag news.Tag) []news.Story
}

type ReportAdder interface {
	Add(report news.Report)
}

func (n *Editor) CreateReport(tag news.Tag, feed NewsFeed, store ReportAdder) error {
	stories := feed.TaggedWith(tag)
	if len(stories) == 0 {
		return errors.New("No story found.")
	}

	store.Add(news.Report{Stories: stories})

	return nil
}
