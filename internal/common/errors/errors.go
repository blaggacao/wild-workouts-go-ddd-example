package errors

import "errors"

type SlugError struct {
	err  string
	slug string
}

func (s SlugError) Error() string {
	return s.err
}

func (s SlugError) Slug() string {
	return s.slug
}

func NewSlugError(err string, slug string) SlugError {
	return SlugError{
		err:  err,
		slug: slug,
	}
}


func (e *AuthorizationError) Unwrap() error { return e.slugError }
func (e *AuthorizationError) Error() string { return e.slugError.Error() }

func NewAuthorizationError(err string, slug string) AuthorizationError {
	return AuthorizationError{slugError: NewSlugError(err, slug)}
}

type IncorrectInputError struct{ slugError SlugError }

func (e *IncorrectInputError) Unwrap() error { return e.slugError }
func (e *IncorrectInputError) Error() string { return e.slugError.Error() }

func NewIncorrectInputError(err string, slug string) IncorrectInputError {
	return IncorrectInputError{slugError: NewSlugError(err, slug)}
}
