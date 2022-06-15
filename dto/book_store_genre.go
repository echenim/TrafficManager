package dto

import (
	"github.com/echenim/openbookstore/models"
	errors "github.com/echenim/openbookstore/utils/errors"
)

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

func (g *GenreDto) Build() (models.Genre, error) {
	logs := errors.ErrorBuilder{}
	if g.Name == "" {
		logs.Affixed("Condition is required")
	}

	return models.Genre{
		Name: g.Name,
	}, logs.Print()
}
