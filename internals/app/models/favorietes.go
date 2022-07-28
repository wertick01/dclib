package models

type FavorieteBooks struct {
	UserId         int64 `json:"user_id"`
	FavoriteBookId int64 `json:"favoriete_book"`
}

type FavorieteAuthors struct {
	UserId           int64 `json:"user_id"`
	FavoriteAuthorId int64 `json:"favoriete_author"`
}
