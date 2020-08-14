// Package interfaces defines all review service interactions
package interfaces

import "review-manager/api/src/entities"

// ReviewsRepository defines all interactions between review service and IO operations
type ReviewsRepository interface {
	CreateReview(review entities.Review) (*int64, error)

	GetReviewForOrder(orderID int64) (*entities.Review, error)

	ExistsReviewForOrder(orderID int64) (bool, error)
}
