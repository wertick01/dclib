package db

import (
	"time"

	"github.com/wertick01/dclib/internals/app/models"
)

func (m *BooksStorage) GetBookingList() ([]*models.Booking, error) {
	stmt := `SELECT book_id, userid, date_of_issue, date_of_delivery, is_confirm FROM dclib_test.booking`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var booking []*models.Booking

	for rows.Next() {
		book := &models.Booking{}
		err := rows.Scan(&book.BookId, &book.UserId, &book.DateOfIssue, &book.DateOfDelivery, &book.IsConfirm)
		if err != nil {
			return nil, err
		}

		booking = append(booking, book)
	}
	return booking, nil
}

func (m *BooksStorage) ReserveBookById(booking *models.Booking) (*models.Books, error) {
	stmt := `INSERT INTO dclib_test.booking (userid, book_id, date_of_issue, is_confirm) VALUES(?, ?, ?, ?)`

	currentTime := time.Now()
	booking.DateOfIssue = currentTime.Format("2006.01.02 15:04:05")

	result, err := m.DB.Exec(stmt, booking.UserId, booking.BookId, booking.DateOfIssue, false)
	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	book, err := m.GetBookById(booking.BookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *BooksStorage) ReturnReversedBook(booking *models.Booking) (*models.Booking, error) {
	stmt := `UPDATE dclib_test.booking SET date_of_delivery = ? WHERE book_id = ? AND userid = ?`
	sdmd := `SELECT id, date_of_issue, is_confirm FROM dclib_test.booking WHERE book_id = ? AND userid = ?`

	currentTime := time.Now()
	booking.DateOfDelivery = currentTime.Format("2006-01-02 15:04:05")
	result, err := m.DB.Exec(stmt, booking.DateOfDelivery, booking.BookId, booking.UserId)
	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := m.DB.QueryRow(sdmd, booking.BookId, booking.UserId)

	err = row.Scan(&booking.Id, &booking.DateOfIssue, &booking.IsConfirm)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *BooksStorage) ConfirmBookRefund(booking *models.Booking) (*models.Booking, error) {
	stmt := `UPDATE dclib_test.booking SET is_confirm = ? WHERE book_id = ? AND userid = ?`
	sdmd := `SELECT id, date_of_issue, date_of_delivery FROM dclib_test.booking WHERE book_id = ? AND userid = ?`

	result, err := m.DB.Exec(stmt, booking.IsConfirm, booking.BookId, booking.UserId)
	if err != nil {
		return nil, err
	}

	row := m.DB.QueryRow(sdmd, booking.BookId, booking.UserId)

	err = row.Scan(&booking.Id, &booking.DateOfIssue, &booking.DateOfDelivery)
	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return booking, nil
}
