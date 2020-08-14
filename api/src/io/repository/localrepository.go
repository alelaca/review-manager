package repository

import (
	"log"
	"review-manager/api/src/entities"
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

func (repo LocalRepository) DeleteReview(id int64) error {
	delete(repo.mapRepository, id)
	return nil
}

func (repo LocalRepository) ExistsReviewForOrder(orderID int64) (bool, error) {
	for _, review := range repo.mapRepository {
		if *review.OrderID == orderID {
			return true, nil
		}
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
