package newsroom

import (
	"errors"

	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/real"
)

type Editor struct{}

type NewsFeed interface {
	TaggedWith(news.Tag) []news.Story
}

func (n *Editor) CreateReport(tag news.Tag, feed NewsFeed) error {
	stories := feed.TaggedWith(tag)
	if len(stories) == 0 {
		return errors.New("No story found.")
	}

	store := &real.ReportStore{}
	store.Add(news.Report{Stories: stories})

	return nil
}
