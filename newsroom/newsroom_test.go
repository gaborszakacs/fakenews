package newsroom_test

import (
	"reflect"
	"testing"

	mock_newsroom "github.com/gaborszakacs/fakenews/mocks"
	"github.com/gaborszakacs/fakenews/news"
	"github.com/gaborszakacs/fakenews/newsroom"
	"github.com/golang/mock/gomock"
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

type SpyStore struct {
	Called  bool
	Reports []news.Report
}

func (s *SpyStore) Add(report news.Report) {
	s.Called = true
	s.Reports = append(s.Reports, report)
}

type MockStore struct {
	AddFn  func(news.Report)
	Called bool
}

func (s *MockStore) Add(report news.Report) {
	s.Called = true
	s.AddFn(report)
}

func TestCreateReport(t *testing.T) {
	t.Run("when there are news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("climate")
		feed := &StubNewsFeed{Stories: []news.Story{
			{Title: "Story1"}}}
		store := &SpyStore{}
		err := e.CreateReport(tag, feed, store)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}
		if !store.Called {
			t.Errorf("Add was expected to be called")
		}
		got := store.Reports
		want := []news.Report{{Stories: []news.Story{{Title: "Story1"}}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Add was called with wrong arguments\n, got:%+v\n, want:%+v\n", got, want)
		}
	})

	t.Run("when there are no news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("climate")
		feed := &StubNewsFeed{Stories: []news.Story{}}
		store := &SpyStore{}
		err := e.CreateReport(tag, feed, store)
		if err == nil {
			t.Errorf("expected error, got none")
		}
	})

	t.Run("when there are no fake news", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("won't have stories")
		feed := &FakeNewsFeed{}
		store := &SpyStore{}
		err := e.CreateReport(tag, feed, store)
		if err == nil {
			t.Errorf("expected error, got none")
		}
	})

	t.Run("with mock", func(t *testing.T) {
		e := newsroom.Editor{}
		tag := news.Tag("won't have stories")
		feed := &StubNewsFeed{Stories: []news.Story{
			{Title: "Story1"},
		},
		}
		store := &MockStore{
			AddFn: func(report news.Report) {
				got := report
				want := news.Report{Stories: []news.Story{{Title: "Story1"}}}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("Add was called with wrong arguments\n, got:%+v\n, want:%+v\n", got, want)
				}
			}}

		e.CreateReport(tag, feed, store)
		// verify if it was called at all.
	})

	t.Run("with generated mocks", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		e := newsroom.Editor{}
		tag := news.Tag("climate")
		stories := []news.Story{
			{Title: "Story1"},
		}
		feed := mock_newsroom.NewMockNewsFeed(mockCtrl)
		feed.EXPECT().TaggedWith(tag).Return(stories).Times(1)
		report := news.Report{Stories: stories}
		store := mock_newsroom.NewMockReportAdder(mockCtrl)
		store.EXPECT().Add(report)
		e.CreateReport(tag, feed, store)
	})
}
