package common

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

/***************/
/* IDGenerator */
/***************/

type IIDGenerator interface {
	Generate() string
}

type UUIDGenerator struct {
}

func (UUIDGenerator) Generate() string {
	return uuid.New().String()
}

func NewUUIDGenerator() IIDGenerator {
	return &UUIDGenerator{}
}

/*****************/
/* SlugGenerator */
/*****************/

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
