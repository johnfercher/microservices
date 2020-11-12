package repository

import "context"

type UserRepository interface {
	GetById(ctx context.Context, id string)
}
