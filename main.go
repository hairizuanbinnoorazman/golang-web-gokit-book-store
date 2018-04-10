package main

import (
	"context"
	"errors"
	"strings"
)

type Stringer struct {
	Text string
}

func (s Stringer) Uppercase() (string, error) {
	if s.Text == "" {
		return "", errors.New("Empty String")
	}
	return strings.ToUpper(s.Text), nil
}

func (s Stringer) Count() int {
	if s.Text == "" {
		return 0
	}
	return len(s.Text)
}

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	z := Stringer{Text: s}
	return z.Uppercase()
}

func (stringService) Count(_ context.Context, s string) int {
	z := Stringer{Text: s}
	return z.Count()
}
