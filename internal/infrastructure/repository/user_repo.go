package repository

import (
	"context"
	"fmt"
	entity "go-jwt/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepo(userCollection *mongo.Collection) UserRepository {
	return &userRepository{
		userCollection: userCollection,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, data *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	userCollection *mongo.Collection
}

func (userRepo *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	bbytes, _ := bson.Marshal(user)

	result, err := userRepo.userCollection.InsertOne(context.Background(), bbytes)
	if err != nil {
		return nil, err
	}

	//how to get the id of the inserted document
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (userRepo *userRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {

	user := entity.User{}

	ID, _ := primitive.ObjectIDFromHex(id)

	result := userRepo.userCollection.
		FindOne(context.Background(),
			bson.M{"_id": ID})

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	return &user, nil
}

func (userRepo *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {

	user := entity.User{}

	result := userRepo.userCollection.
		FindOne(context.Background(),
			bson.M{"username": username})

	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	if user.Username == "" {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	return &user, nil
}

func (userRepo *userRepository) UpdateUser(ctx context.Context, id string, data *entity.User) (*entity.User, error) {

	ID, _ := primitive.ObjectIDFromHex(id)

	fmt.Println(data)
	result, err := userRepo.userCollection.
		UpdateOne(context.Background(),
			bson.M{"_id": ID},
			bson.M{
				"$set": data,
			})

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return data, nil
}

func (userRepo *userRepository) DeleteUser(ctx context.Context, id string) error {
	ID, _ := primitive.ObjectIDFromHex(id)
	_, err := userRepo.userCollection.DeleteOne(context.Background(), bson.M{"_id": ID})
	if err != nil {
		return err
	}

	return nil
}

/*--------------------------------------------*/
// Maybe we can reuse two function below ??

// func (userRepo *UserRepoImpl) CheckLoginInfo(email, password string) (models.User, error) {
// 	user := models.User{}

// 	result := userRepo.userCollection.
// 						FindOne(context.Background(),
// 								bson.M{
// 									"email" :email,
// 									"password": password,
// 								})

// 	err := result.Decode(&user)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
// func (userRepo *UserRepoImpl) FindUserByEmail(email string) (models.User, error) {
// 	user := models.User{}

// 	result := userRepo.userCollection.
// 						FindOne(context.Background(),
// 								bson.M{"email" :email})

// 	err := result.Decode(&user)

// 	if user == (models.User{}) {
// 		return user, models.ERR_USER_NOT_FOUND
// 	}

// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
