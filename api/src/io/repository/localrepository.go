package repository

import (
	"log"
	"review-manager/api/src/entities"
	"time"
)

type LocalRepository struct {
	mapRepository map[int64]*entities.Review
	nextID        *int64
}

func InitializeLocalRepository() LocalRepository {
	idInit := int64(1)

	return LocalRepository{
		mapRepository: map[int64]*entities.Review{},
		nextID:        &idInit,
	}
}

func (repo LocalRepository) CreateReview(review entities.Review) (*int64, error) {
	review.ID = new(int64)
	*review.ID = *repo.nextID

	*repo.nextID++
	log.Println(repo.nextID)

	repo.mapRepository[*review.ID] = &review

	return review.ID, nil
}

func (repo LocalRepository) DeleteReview(id int64) (bool, error) {
	review := repo.mapRepository[id]

	if review != nil {
		*review.Deleted = true
		return true, nil
	}

	return false, nil
}

func (repo LocalRepository) GetReviewForOrder(orderID int64) (*entities.Review, error) {
	for _, review := range repo.mapRepository {
		if *review.OrderID == orderID {
			return review, nil
		}
	}

	return nil, nil
}

func (repo LocalRepository) GetReviewsForStore(shopID int64, dateFrom, dateTo time.Time) ([]entities.Review, error) {
	reviews := []entities.Review{}

	for _, review := range repo.mapRepository {
		if *review.ShopID == shopID && !*review.Deleted && isDateInRange(*review.DateCreated, dateFrom, dateTo) {
			reviews = append(reviews, *review)
		}
	}

	return reviews, nil
}

func (repo LocalRepository) ExistsReviewForOrder(orderID int64) (bool, error) {
	for _, review := range repo.mapRepository {
		if *review.OrderID == orderID {
			return true, nil
		}
	}

	return false, nil
}

func isDateInRange(date, dateFrom, dateTo time.Time) bool {
	return date.Equal(dateFrom) || date.Equal(dateTo) || (date.Before(dateTo) && date.After(dateFrom))
}
