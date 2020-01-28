package real

import (
	"time"

	"github.com/gaborszakacs/fakenews/news"
)

type NewsFeed struct{}

func (n *NewsFeed) TaggedWith(news.Tag) []news.Story {
	time.Sleep(3 * time.Second)
	return []news.Story{
		{Title: "Climate change helped spawn East Africa’s locust crisis"},
		{Title: "Hundreds of Amazon workers criticize the company's climate policy"},
	}
}

type ReportStore struct {
	reports []news.Report
}

func (rs *ReportStore) Add(report news.Report) {
	rs.reports = append(rs.reports, report)
}

func (rs *ReportStore) First() news.Report {
	return rs.reports[0]
}
