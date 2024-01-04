package usecase

import (
	"context"
	"fmt"
	entity "go-jwt/internal/entity"
	repository "go-jwt/internal/infrastructure/repository"
	"time"
)

func NewUserUsecase(userRepo repository.UserRepository, bookRepo repository.BookRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}

type UserUsecase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, data *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	AuthenticateUser(ctx context.Context, username string, password string) (*entity.User, error)
	ReserveBook(ctx context.Context, userID string, bookID string) (*entity.User, error)
	BorrowBook(ctx context.Context, userID string, bookID string) (*entity.User, error)
	ExtendBorrowBook(ctx context.Context, userID string, bookID string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	bookRepo repository.BookRepository
}

func (s *userUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userUsecase) UpdateUser(ctx context.Context, id string, data *entity.User) (*entity.User, error) {
	return s.userRepo.UpdateUser(ctx, id, data)
}

func (s *userUsecase) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func (s *userUsecase) AuthenticateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	if user.Password != password {
		return nil, entity.ERR_USER_PASSWORD_NOT_MATCH
	}

	return user, nil
}

func (s *userUsecase) ReserveBook(ctx context.Context, userID string, bookID string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	if user.CountFine > 2 {
		return nil, entity.ERR_USER_FINE_EXCEED
	}

	book, err := s.bookRepo.GetBookByID(ctx, bookID)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, entity.ERR_BOOK_NOT_FOUND
	}

	if book.Availability != 0 {
		return nil, entity.ERR_BOOK_NOT_AVAILABLE
	}

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 3)
	extendedDate := time.Time{}

	book.Availability = 2 // 2 = reserved
	book.StartDate = startDate
	book.EndDate = endDate
	book.ExtendedDate = extendedDate

	// update book availability
	book, err = s.bookRepo.UpdateBook(ctx, bookID, book)

	if err != nil {
		return nil, err
	}

	user.ReservingList = append(user.ReservingList, entity.UserActivity{
		BookId:       book.ID,
		BookName:     book.Name,
		StartDate:    startDate,
		EndDate:      endDate,
		ExtendedDate: extendedDate,
	})

	// update user reserving book
	user, err = s.userRepo.UpdateUser(ctx, userID, user)

	//print user after
	fmt.Println("user after", user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userUsecase) BorrowBook(ctx context.Context, userID string, bookID string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	if user.CountFine > 2 {
		return nil, entity.ERR_USER_FINE_EXCEED
	}

	book, err := s.bookRepo.GetBookByID(ctx, bookID)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, entity.ERR_BOOK_NOT_FOUND
	}

	if book.Availability != 0 {
		return nil, entity.ERR_BOOK_NOT_AVAILABLE
	}

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	extendedDate := time.Time{}

	book.Availability = 3 // 3 = borrowed
	book.StartDate = startDate
	book.EndDate = endDate
	book.ExtendedDate = extendedDate

	// update book availability
	book, err = s.bookRepo.UpdateBook(ctx, bookID, book)

	if err != nil {
		return nil, err
	}

	user.BorrowingList = append(user.BorrowingList, entity.UserActivity{
		BookId:       book.ID,
		BookName:     book.Name,
		StartDate:    startDate,
		EndDate:      endDate,
		ExtendedDate: extendedDate,
	})

	// update user reserving book
	user, err = s.userRepo.UpdateUser(ctx, userID, user)

	//print user after
	fmt.Println("user after", user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userUsecase) ExtendBorrowBook(ctx context.Context, userID string, bookID string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ERR_USER_NOT_FOUND
	}

	book, err := s.bookRepo.GetBookByID(ctx, bookID)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, entity.ERR_BOOK_NOT_FOUND
	}

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(user.BorrowingList); i++ {
		if user.BorrowingList[i].BookId == book.ID {
			if user.BorrowingList[i].ExtendedDate.IsZero() {
				user.BorrowingList[i].ExtendedDate = user.BorrowingList[i].EndDate.AddDate(0, 0, 7)
			} else {
				return nil, entity.ERR_BOOK_RESERVE_ALREADY_EXTENDED
			}
		}
	}

	// update user reserving book
	user, err = s.userRepo.UpdateUser(ctx, userID, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
