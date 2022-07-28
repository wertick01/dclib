package models

type Booking struct {
	Id             int64       `json:"id"`
	BookId         int64       `json:"book_id"`
	UserId         int64       `json:"user_id"`
	DateOfIssue    string      `json:"date_of_issue"`
	DateOfDelivery interface{} `json:"date_of_delivery"`
	IsConfirm      bool        `json:"is_confirm"`
}
