package service

import (
	"github.com/gosimple/slug"
)

type ISlugGenerator interface {
	Generate(value string) string
}

type SlugGenerator struct {
}

func (SlugGenerator) Generate(value string) string {
	return slug.Make(value)
}

func NewSlugGenerator() ISlugGenerator {
	return &SlugGenerator{}
}
