package models

type Books struct {
	BookId    int64     `json:"book_id"`
	Count     int64     `json:"count"`
	Stars     int64     `json:"stars"`
	BookPhoto string    `json:"book_photo"`
	BookName  string    `json:"book_name"`
	Authors   []Authors `json:"author"`
}

type FinalBooks struct {
	Book    Books     `json:"books"`
	Authors []Authors `json:"authors"`
}
