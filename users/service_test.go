package users

import (
	"os"
	"testing"
	"time"

	"github.com/avasapollo/users-api/global"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	le   *logrus.Entry
	conf *global.AppConfig
)

func TestMain(m *testing.M) {
	le = logrus.WithField("service", "testing")
	conf = &global.AppConfig{}
	os.Exit(m.Run())
}

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := NewUser(
		uuid.New().String(),
		"rayhack",
		"andrea",
		"vasapollo",
		time.Now())
	db := NewMockDatabase(ctrl)
	db.EXPECT().CreateUser(user).Return(nil)
	svc := NewService(le, conf, db)
	err := svc.CreateUser(user)
	assert.Nil(t, err)
}

func TestService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	u := &User{
		ID:        "user-id-1",
		Nickname:  "rayhack",
		Name:      "andrea",
		Surname:   "vasapollo",
		CreatedAt: time.Now(),
	}
	db := NewMockDatabase(ctrl)
	db.EXPECT().GetUser("user-id-1").Return(u, nil)
	svc := NewService(le, conf, db)
	user, err := svc.GetUser("user-id-1")
	assert.Nil(t, err)
	assert.Equal(t, user, u)
}

func TestService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	u := &User{
		ID:        "user-id-1",
		Nickname:  "rayhack",
		Name:      "andrea",
		Surname:   "vasapollo",
		CreatedAt: time.Now(),
	}
	db := NewMockDatabase(ctrl)
	db.EXPECT().UpdateUser(u).Return(nil)
	svc := NewService(le, conf, db)
	err := svc.UpdateUser(u)
	assert.Nil(t, err)
}

func TestService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := NewMockDatabase(ctrl)
	db.EXPECT().DeleteUser("user-id-1").Return(nil)
	svc := NewService(le, conf, db)
	err := svc.DeleteUser("user-id-1")
	assert.Nil(t, err)
}
