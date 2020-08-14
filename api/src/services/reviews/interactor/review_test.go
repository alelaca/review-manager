package interactor

import (
	"errors"
	"review-manager/api/src/entities"
	"review-manager/api/src/io/mocks"
	"testing"
)

func TestUnitCreateReview_Ok(t *testing.T) {
	comment := "review comment"
	rate := 4
	shopID := int64(1)
	userID := int64(1)
	orderID := int64(1)

	reviewMock := &entities.Review{
		Comment: &comment,
		Rate:    &rate,
		ShopID:  &shopID,
		UserID:  &userID,
		OrderID: &orderID,
	}

	repositoryMock := mocks.RepositoryMock{
		NextReviewID: int64(10),
		NextError:    nil,
	}

	reviewService := ReviewsService{
		Repository: repositoryMock,
	}

	_, err := reviewService.CreateReview(*reviewMock)

	if err != nil {
		t.Errorf("Fail: Test shouldnt return error")
	}
}

func TestUnitCreateReview_Fail(t *testing.T) {
	comment := "review comment"
	rate := 4
	shopID := int64(1)
	userID := int64(1)
	orderID := int64(1)

	reviewMock := &entities.Review{
		Comment: &comment,
		Rate:    &rate,
		ShopID:  &shopID,
		UserID:  &userID,
		OrderID: &orderID,
	}

	repositoryMock := mocks.RepositoryMock{
		NextError: errors.New("error"),
	}

	reviewService := ReviewsService{
		Repository: repositoryMock,
	}

	_, err := reviewService.CreateReview(*reviewMock)

	if err == nil {
		t.Errorf("Fail: Test should have return error")
	}
}
