package news

type Tag string

type Story struct {
	Title string
}

type Report struct {
	Stories []Story
}
