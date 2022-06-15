package dto

import (
	"github.com/echenim/openbookstore/models/domain"
)

// errors "github.com/echenim/openbookstore/utils/errors"

type GenreDto struct {
	ID   int
	Name string
}

func NewGenreDto() *GenreDto {
	return &GenreDto{}
}

func (g *GenreDto) SetName(name string) {
	g.Name = name
}

func (g *GenreDto) Build() (domain.Genre, error) {
	// logs := errors.ErrorBuilder{}
	if g.Name == "" {
		//	logs.Affixed("Condition is required")
	}

	return domain.Genre{
		Name: g.Name,
	}, nil // logs.Print()
}
