package mocks

import "review-manager/api/src/entities"

type RepositoryMock struct {
	NextReview   *entities.Review
	NextReviewID int64
	NextFound    bool

	NextError error
}

func (mock RepositoryMock) CreateReview(review entities.Review) (int64, error) {
	return mock.NextReviewID, mock.NextError
}

func (mock RepositoryMock) GetReviewForOrder(orderID int64) (*entities.Review, error) {
	return mock.NextReview, mock.NextError
}

func (mock RepositoryMock) ExistsReviewForOrder(orderID int64) (bool, error) {
	return mock.NextFound, mock.NextError
}
