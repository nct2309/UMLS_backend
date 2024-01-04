package request

import (
	entity "go-jwt/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func NewBookRequest() BookRequest {
	return &bookRequest{}
}

type BookRequest interface {
	Bind(c *gin.Context) error
	//Validate() error
	GetID() string
	GetIDFromURL(c *gin.Context) string
	GetISBN() string
	GetName() string
	GetNameFromURL(c *gin.Context) string
	GetCondition() bool
	GetAvailability() int
	GetLocation() string
	GetStartDate() time.Time
	GetEndDate() time.Time
	GetExtendedDate() time.Time
	GetAuthor() string
	GetImageURL() string
	GetGenre() string
	GetTotalPages() int
	GetDescription() string
	GetPublisher() string
	GetPulishDate() time.Time
}

type bookRequest struct {
	book entity.Book
}

func (r *bookRequest) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(&r.book)
}

func (r *bookRequest) GetID() string {
	return r.book.ID.Hex()
}

func (r *bookRequest) GetIDFromURL(c *gin.Context) string {
	return c.Param("id")
}

func (r *bookRequest) GetISBN() string {
	return r.book.ISBN
}

func (r *bookRequest) GetName() string {
	return r.book.Name
}

func (r *bookRequest) GetNameFromURL(c *gin.Context) string {
	return c.Param("name")
}

func (r *bookRequest) GetCondition() bool {
	return r.book.Condition
}

func (r *bookRequest) GetAvailability() int {
	return r.book.Availability
}

func (r *bookRequest) GetLocation() string {
	return r.book.Location
}

func (r *bookRequest) GetStartDate() time.Time {
	return r.book.StartDate
}

func (r *bookRequest) GetEndDate() time.Time {
	return r.book.EndDate
}

func (r *bookRequest) GetExtendedDate() time.Time {
	return r.book.ExtendedDate
}

func (r *bookRequest) GetAuthor() string {
	return r.book.Author
}

func (r *bookRequest) GetImageURL() string {
	return r.book.ImageURL
}

func (r *bookRequest) GetGenre() string {
	return r.book.Genre
}

func (r *bookRequest) GetTotalPages() int {
	return r.book.TotalPages
}

func (r *bookRequest) GetDescription() string {
	return r.book.Description
}

func (r *bookRequest) GetPublisher() string {
	return r.book.Publisher
}

func (r *bookRequest) GetPulishDate() time.Time {
	return r.book.PublishDate
}

// func (r *bookRequest) Validate() error {
// 	if r.ID == 0 {
// 		return errors.New("ID is required")
// 	} else if r.Email == "" {
// 		return errors.New("Email is required")
// 	} else if r.Password == "" {
// 		return errors.New("password is required")
// 	} else if r.Name == "" {
// 		return errors.New("name is required")
// 	}
// 	return nil
// }
