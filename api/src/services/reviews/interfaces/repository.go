// Package interfaces defines all review service interactions
package interfaces

import "review-manager/api/src/entities"

// ReviewsRepositoryInterface defines all interactions between review service and IO operations
type ReviewsRepositoryInterface interface {
	GetReview(id int64) (*entities.Review, error)

	CreateReview(review entities.Review) (*int64, error)

	ExistsReviewForOrder(orderID int64) (bool, error)
}
