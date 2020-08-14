package interactor

import (
	"errors"
	"net/http"
	"review-manager/api/src/customerror"
	"review-manager/api/src/entities"
	"review-manager/api/src/io/mocks"
	"testing"
)

func TestUnitGetReview_Ok(t *testing.T) {
	id := int64(10)
	comment := "review comment"
	rate := 4
	shopID := int64(1)
	userID := int64(1)
	orderID := int64(1)

	reviewMock := &entities.Review{
		ID:      &id,
		Comment: &comment,
		Rate:    &rate,
		ShopID:  &shopID,
		UserID:  &userID,
		OrderID: &orderID,
	}

	repositoryMock := mocks.RepositoryMock{
		NextReview: reviewMock,
		NextError:  nil,
	}

	reviewService := ReviewsService{
		Repository: repositoryMock,
	}

	review, err := reviewService.FindReviewByID(int64(10))

	if err != nil {
		t.Errorf("Fail: Test shouldnt return error")
	}

	if review == nil {
		t.Errorf("Test GetReview cant return nil review")
	}
}

func TestUnitGetReview_NotFound(t *testing.T) {

	repositoryMock := mocks.RepositoryMock{
		NextReview: nil,
		NextError:  nil,
	}

	reviewService := ReviewsService{
		Repository: repositoryMock,
	}

	_, err := reviewService.FindReviewByID(int64(10))

	if err != nil {
		if apiError, ok := err.(*customerror.Error); ok {
			if apiError.StatusCode() != http.StatusNotFound {
				t.Errorf("GetReview Fail. Expected status code %v, got %v", http.StatusNotFound, apiError.StatusCode())
			}
		} else {
			t.Errorf("GetReview Fail. Method returned error without status code")
		}
		return
	}

	t.Errorf("Test should have failed")
}

func TestUnitGetReview_RepositoryError(t *testing.T) {

	repositoryMock := mocks.RepositoryMock{
		NextReview: nil,
		NextError:  errors.New("error"),
	}

	reviewService := ReviewsService{
		Repository: repositoryMock,
	}

	_, err := reviewService.FindReviewByID(int64(10))

	if err == nil {
		t.Errorf("Test should have failed")
	}
}

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
