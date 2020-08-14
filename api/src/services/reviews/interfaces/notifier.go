package interfaces

import (
	"review-manager/api/src/entities"
)

type ReviewsNotifier interface {
	Publish(message entities.Review) error
}
