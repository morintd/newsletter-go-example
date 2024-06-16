package service

import (
	"github.com/google/uuid"
)

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
