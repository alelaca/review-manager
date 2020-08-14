package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"review-manager/api/src/customerror"
	"review-manager/api/src/entities"
	"review-manager/api/src/environment"
	reviews "review-manager/api/src/services/reviews/interactor"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func createReview(env environment.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		reviewService := reviews.ReviewsService{
			Repository: env.LocalRepository,
		}

		requestContent := c.Request.Body
		body, err := ioutil.ReadAll(requestContent)

		var review entities.Review
		if err := json.Unmarshal(body, &review); err != nil {
			abortWithCustomError(c, http.StatusBadRequest, fmt.Errorf("invalid review data"))
			return
		}

		err = validateReviewParameters(review)
		if err != nil {
			abortWithCustomError(c, http.StatusBadRequest, err)
			return
		}

		reviewID, err := reviewService.CreateReview(review)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": reviewID})
	}
}

func deleteReview(env environment.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		reviewService := reviews.ReviewsService{
			Repository: env.LocalRepository,
		}

		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			abortWithCustomError(c, http.StatusBadRequest, fmt.Errorf("invalid review id '%v', it needs to be a number", idParam))
			return
		}

		err = reviewService.DeleteReview(id)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("review with id '%v' deleted", id)})
	}
}

func findReviewByOrderID(env environment.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		reviewService := reviews.ReviewsService{
			Repository: env.LocalRepository,
		}

		orderIDParam := c.Param("orderID")

		orderID, err := strconv.ParseInt(orderIDParam, 10, 64)
		if err != nil {
			abortWithCustomError(c, http.StatusBadRequest, fmt.Errorf("invalid order id '%v', it needs to be a number", orderIDParam))
			return
		}

		review, err := reviewService.GetReviewForOrder(orderID)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, review)
	}
}

func findReviewByShopID(env environment.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		reviewService := reviews.ReviewsService{
			Repository: env.LocalRepository,
		}

		shopIDParam := c.Param("shopID")

		shopID, err := strconv.ParseInt(shopIDParam, 10, 64)
		if err != nil {
			abortWithCustomError(c, http.StatusBadRequest, fmt.Errorf("invalid order id '%v', it needs to be a number", shopIDParam))
			return
		}

		dateFromStr := c.Query("date_from")
		dateFrom, err := getDate(dateFromStr)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		dateToStr := c.Query("date_to")
		dateTo, err := getDate(dateToStr)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		reviews, err := reviewService.GetReviewsForStore(shopID, *dateFrom, *dateTo)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, reviews)
	}
}

func validateReviewParameters(review entities.Review) error {

	if review.Rate == nil || review.OrderID == nil || review.ShopID == nil || review.UserID == nil {
		return fmt.Errorf("review params 'rate', 'order_id', 'shop_id', 'user_id' are required")
	}

	return nil
}

func getDate(dateStr string) (*time.Time, error) {
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		customerror.WrapWithStatusCode(err, http.StatusBadRequest, fmt.Sprintf("invalid date format, it needs to follow %v, got %v", time.RFC3339, dateStr))
	}

	return &date, nil

}
