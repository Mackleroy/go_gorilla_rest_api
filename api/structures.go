package api

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateCategoryRequest struct {
	CategoryName string
}

type CreateGoodRequest struct {
	GoodName     string
	CategoryName string
	Price        uint32
	Width        uint16
	Length       uint16
	Tags         []string
}

func (r CreateGoodRequest) Validate() error {
	if !ProjectStorageInstance.HasCategory(r.CategoryName) {
		return errors.New("category does not exist")
	}
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.GoodName, validation.Required, validation.Length(2, 256)),
		validation.Field(&r.CategoryName, validation.Required, validation.Length(2, 256)),
		validation.Field(&r.Price, validation.Required, validation.Min(uint32(1))),
		validation.Field(&r.Width, validation.Required, validation.Min(uint16(1)), validation.Max(uint16(1000))),
		validation.Field(&r.Length, validation.Required, validation.Min(uint16(1)), validation.Max(uint16(1000))),
		validation.Field(&r.Tags, validation.Length(2, 10)),
	)
}
