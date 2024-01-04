package usecase

import (
	"context"
	entity "go-jwt/internal/entity"
	repository "go-jwt/internal/infrastructure/repository"
)

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: bookRepo,
	}
}

type BookUsecase interface {
	CreateBook(ctx context.Context, Book *entity.Book) (*entity.Book, error)
	GetBookByID(ctx context.Context, id string) (*entity.Book, error)
	UpdateBookByID(ctx context.Context, id string, data *entity.Book) (*entity.Book, error)
	DeleteBookByID(ctx context.Context, id string) error
	SearchBookByName(ctx context.Context, name string) ([]*entity.Book, error)
	GetAllBooks(ctx context.Context) ([]*entity.Book, error)
	UpdateBookAvailability(ctx context.Context, id string, availability int) (*entity.Book, error)
}

type bookUsecase struct {
	repo repository.BookRepository
}

func (s *bookUsecase) CreateBook(ctx context.Context, Book *entity.Book) (*entity.Book, error) {
	return s.repo.CreateBook(ctx, Book)
}

func (s *bookUsecase) GetBookByID(ctx context.Context, id string) (*entity.Book, error) {
	return s.repo.GetBookByID(ctx, id)
}

func (s *bookUsecase) UpdateBookByID(ctx context.Context, id string, data *entity.Book) (*entity.Book, error) {
	return s.repo.UpdateBook(ctx, id, data)
}

func (s *bookUsecase) DeleteBookByID(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}

func (s *bookUsecase) SearchBookByName(ctx context.Context, name string) ([]*entity.Book, error) {
	return s.repo.GetBooksByName(ctx, name)
}

func (s *bookUsecase) GetAllBooks(ctx context.Context) ([]*entity.Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *bookUsecase) UpdateBookAvailability(ctx context.Context, id string, availability int) (*entity.Book, error) {
	data := &entity.Book{
		Availability: availability,
	}
	return s.repo.UpdateBook(ctx, id, data)
}
