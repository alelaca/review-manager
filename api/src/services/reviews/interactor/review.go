// Package interactor is in charge of business logic for reviews
package interactor

import (
	"fmt"
	"net/http"
	"review-manager/api/src/customerror"
	"review-manager/api/src/entities"
	"review-manager/api/src/services/reviews/interfaces"
	"time"
)

type ReviewsService struct {
	Repository interfaces.ReviewsRepository
}

// CreateReview creates a new review in repository
func (reviewService ReviewsService) CreateReview(review entities.Review) (*int64, error) {

	err := reviewService.isValidReview(review)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	review.DateCreated = &now
	review.DateLastUpdated = &now

	return reviewService.Repository.CreateReview(review)
}

// GetReviewForOrder gets a review given order id
func (reviewService ReviewsService) GetReviewForOrder(orderID int64) (*entities.Review, error) {
	review, err := reviewService.Repository.GetReviewForOrder(orderID)
	if err != nil {
		return nil, customerror.WrapWithStatusCode(err, http.StatusInternalServerError, fmt.Sprintf("error finding review by order id '%v'", orderID))
	}

	if review == nil {
		return nil, customerror.WrapWithStatusCode(nil, http.StatusNotFound, fmt.Sprintf("review by order id '%v' not found", orderID))
	}

	return review, nil
}

func (reviewService ReviewsService) isValidReview(review entities.Review) error {
	if review.OrderID == nil {
		return customerror.WrapWithStatusCode(nil, http.StatusBadRequest, fmt.Sprintf("Order id cant be null"))
	}

	orderExists, err := reviewService.existsReviewForOrder(*review.OrderID)
	if err != nil {
		return customerror.WrapWithStatusCode(nil, http.StatusInternalServerError, fmt.Sprintf("Error checking if order exists in repository"))
	}

	if orderExists {
		return customerror.WrapWithStatusCode(nil, http.StatusBadRequest, fmt.Sprintf("The order with id '%v' already has a review associated", *review.OrderID))
	}

	return nil
}

func (reviewService ReviewsService) existsReviewForOrder(orderID int64) (bool, error) {
	return reviewService.Repository.ExistsReviewForOrder(orderID)
}
