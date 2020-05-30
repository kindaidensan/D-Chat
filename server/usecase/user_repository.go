package usecase

import "github.com/kindaidensan/D-Chat/domain"

type UserRepository interface {
	Create(domain.User) (domain.User, error)
}
