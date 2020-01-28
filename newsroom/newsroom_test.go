package newsroom_test

import (
	"testing"

	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/newsroom"
)

type StubNewsFeed struct {
	Stories []news.Story
}

func (s *StubNewsFeed) TaggedWith(news.Tag) []news.Story {
	return s.Stories
}

func TestCreateReport(t *testing.T) {
	t.Run("when there are news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("climate")
		feed := &StubNewsFeed{Stories: []news.Story{
			{Title: "Story1"}}}
		err := e.CreateReport(tag, feed)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
	})

	t.Run("when there are no news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("climate")
		feed := &StubNewsFeed{Stories: []news.Story{}}
		err := e.CreateReport(tag, feed)
		if err == nil {
			t.Errorf("expected error, got none")
		}
	})
}
