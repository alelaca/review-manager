package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"review-manager/api/src/entities"
	"review-manager/api/src/environment"
	reviews "review-manager/api/src/services/reviews/interactor"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ping success")
}

func findReviewByID(env environment.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		reviewService := reviews.ReviewsService{
			Repository: env.LocalRepository,
		}

		idStr := c.Param("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			abortWithCustomError(c, http.StatusBadRequest, fmt.Errorf("invalid review id '%v', it needs to be a number", idStr))
			return
		}

		review, err := reviewService.FindReviewByID(id)
		if err != nil {
			abortWithCustomError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, review)
	}
}

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

func validateReviewParameters(review entities.Review) error {

	if review.Rate == nil || review.OrderID == nil || review.ShopID == nil || review.UserID == nil {
		return fmt.Errorf("review params 'rate', 'order_id', 'shop_id', 'user_id' are required")
	}

	return nil
}
