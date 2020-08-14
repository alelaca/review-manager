package mocks

import "review-manager/api/src/entities"

type RepositoryMock struct {
	NextReview   *entities.Review
	NextReviewID int64

	NextError error
}

func (mock RepositoryMock) GetReview(id int64) (*entities.Review, error) {
	return mock.NextReview, mock.NextError
}

func (mock RepositoryMock) CreateReview(review entities.Review) (int64, error) {
	return mock.NextReviewID, mock.NextError
}
