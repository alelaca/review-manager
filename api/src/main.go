package main

import (
	"review-manager/api/src/environment"
	"review-manager/api/src/io/repository"
	"review-manager/api/src/io/router"
)

func main() {

	localRepository := repository.InitializeLocalRepository()

	env := environment.Environment{
		LocalRepository: localRepository,
	}

	router.InitializeRouter(env)
}
