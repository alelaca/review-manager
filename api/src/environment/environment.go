package environment

import (
	"review-manager/api/src/io/repository"
)

type Environment struct {
	LocalRepository repository.LocalRepository
}
