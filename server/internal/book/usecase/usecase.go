package usecase

import (
	"context"
	"github.com/scul0405/my-shop/server/config"
	dbmodels "github.com/scul0405/my-shop/server/db/models"
	"github.com/scul0405/my-shop/server/internal/book"
	"github.com/scul0405/my-shop/server/internal/dto"
	"github.com/scul0405/my-shop/server/pkg/logger"
	"github.com/scul0405/my-shop/server/pkg/utils"
)

type bookUseCase struct {
	cfg      *config.Config
	bookRepo book.Repository
	logger   logger.Logger
}

func NewBookUseCase(
	cfg *config.Config,
	bookRepo book.Repository,
	logger logger.Logger) book.UseCase {
	return &bookUseCase{cfg: cfg, bookRepo: bookRepo, logger: logger}
}

func (u *bookUseCase) Create(ctx context.Context, book *dto.BookDTO) (*dto.BookDTO, error) {
	createdBook, err := u.bookRepo.Create(ctx, book.ToModel())
	if err != nil {
		return nil, err
	}

	return bookModelToDto(createdBook), nil
}

func (u *bookUseCase) GetByID(ctx context.Context, id uint64) (*dto.BookDTO, error) {
	bookModel, err := u.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return bookModelToDto(bookModel), nil
}

func (u *bookUseCase) Update(ctx context.Context, book *dto.BookDTO) error {
	whiteList := []string{"category_id", "name", "sku", "desc", "image", "price", "total_sold", "quantity", "status"}
	err := u.bookRepo.Update(ctx, book.ToModel(), whiteList...)
	if err != nil {
		return err
	}

	return nil
}

func (u *bookUseCase) Delete(ctx context.Context, id uint64) error {
	err := u.bookRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *bookUseCase) List(ctx context.Context, pq *utils.PaginationQuery) (*utils.PaginationList, error) {
	paginationList, err := u.bookRepo.List(ctx, pq)
	if err != nil {
		return nil, err
	}

	var booksDTO []*dto.BookDTO
	for _, bookModel := range paginationList.List.(dbmodels.BookSlice) {
		booksDTO = append(booksDTO, bookModelToDto(bookModel))
	}

	paginationList.List = booksDTO

	return paginationList, nil
}

func bookModelToDto(book *dbmodels.Book) *dto.BookDTO {
	return &dto.BookDTO{
		ID:         uint64(book.ID),
		CategoryID: uint64(book.CategoryID),
		Name:       book.Name,
		SKU:        book.Sku,
		Desc:       book.Desc.String,
		Image:      book.Image.String,
		Price:      book.Price,
		TotalSold:  book.TotalSold,
		Quantity:   book.Quantity,
		Status:     book.Status,
	}
}