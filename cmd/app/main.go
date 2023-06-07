package main

import (
	"github.com/johnsoncwb/userApi/internal/adapters/api"
	"github.com/johnsoncwb/userApi/internal/initializer"
)

func main() {

	initializer.StartApp()
	defer initializer.FinishApp()

	api.LoadRoutes()

}
