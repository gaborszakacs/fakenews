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

type FakeNewsFeed struct{}

func (f *FakeNewsFeed) TaggedWith(tag news.Tag) []news.Story {
	if tag == "tag-with-stories" {
		return []news.Story{{Title: "Story"}}
	} else {
		return []news.Story{}
	}
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

	t.Run("when there are no fake news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("won't have stories")
		feed := &FakeNewsFeed{}
		err := e.CreateReport(tag, feed)
		if err == nil {
			t.Errorf("expected error, got none")
		}
	})
}
