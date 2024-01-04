package cmd

import (
	controller "go-jwt/internal/controller"
	"go-jwt/internal/infrastructure/driver"
	repo "go-jwt/internal/infrastructure/repository"
	usecase "go-jwt/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s server) SetupControllers() {

	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	db := driver.ConnectMongoDB()

	// init collection of mongodb to use
	userCollection := db.Client.Database("Library-Management-Database").Collection("User")
	bookCollection := db.Client.Database("Library-Management-Database").Collection("Book")

	// // create constraint unqiue username for User collection
	// indexOptions := options.Index().SetUnique(true)
	// indexKeys := bson.M{"username": 1}
	// result, err := userCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: indexKeys, Options: indexOptions})

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// init repository
	userRepo := repo.NewUserRepo(userCollection)
	bookRepo := repo.NewBookRepo(bookCollection)

	// init usecase
	userUsecase := usecase.NewUserUsecase(userRepo, bookRepo)
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	// init controller
	controller.SetupUserRoutes(s.router, userUsecase)
	controller.SetupBookRoutes(s.router, bookUsecase)
}
