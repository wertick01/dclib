package processors

import (
	"github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/models"
)

type ReserveProcessor struct {
	storage *db.BooksStorage
}

func NewReserveProcessor(storage *db.BooksStorage) *ReserveProcessor {
	processor := new(ReserveProcessor)
	processor.storage = storage
	return processor
}

func (processor *ReserveProcessor) GetList() ([]*models.Booking, error) {
	return processor.storage.GetBookingList()
}

func (processor *ReserveProcessor) ReserveBook(booking *models.Booking) (*models.Books, error) {
	reserved, err := processor.storage.ReserveBookById(booking)

	if err != nil {
		return nil, err
	}

	return reserved, nil
}

func (processor *ReserveProcessor) ReturnBook(booking *models.Booking) (*models.Booking, error) {
	returned, err := processor.storage.ReturnReversedBook(booking)

	if err != nil {
		return nil, err
	}

	return returned, nil
}

func (processor *ReserveProcessor) ConfirmBook(booking *models.Booking) (*models.Booking, error) {
	confirmed, err := processor.storage.ConfirmBookRefund(booking)

	if err != nil {
		return nil, err
	}

	return confirmed, nil
}
