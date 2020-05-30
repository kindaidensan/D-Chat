package usecase

import (
	"log"

	"github.com/kindaidensan/D-Chat/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(userRepository UserRepository) *UserInteractor {
	userInteractor := UserInteractor{
		UserRepository: userRepository,
	}
	return &userInteractor
}

func (interactor *UserInteractor) Create(user domain.User) (domain.User, error) {
	u, err := interactor.UserRepository.Create(user)
	if err != nil {
		log.Printf("[Faild] create user { ID: %s, Name: %s, MailAddress: %s, Profile: %s, Status: %s, StatusMessage: %s }",
			u.ID,
			u.Name,
			u.MailAddress,
			u.Profile,
			u.Status,
			u.StatusMessage)
		return domain.User{}, err
	}
	return u, nil
}
