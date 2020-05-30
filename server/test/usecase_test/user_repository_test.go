package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindaidensan/D-Chat/domain"
	mock "github.com/kindaidensan/D-Chat/test/mock_usecase"
	"github.com/kindaidensan/D-Chat/usecase"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	uMock := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(uMock)

	user := domain.User{
		ID:            "testid",
		Name:          "test name",
		MailAddress:   "test@test.jp",
		Profile:       "test profile",
		Status:        "offline",
		StatusMessage: "test message",
	}
	uMock.EXPECT().Create(user).Return(user, nil)
	u, err := interactor.Create(user)
	if err != nil {
		t.Errorf("[faild] Expected no error")
	}
	if user != u {
		t.Errorf("[faild] Expected input user === return user")
	}

	uMock.EXPECT().Create(user).Return(user, errors.New("faild create user"))
	u, err = interactor.Create(user)
	if err == nil {
		t.Errorf("[faild] Expected error")
	}
}
