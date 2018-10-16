package database

import (
	"os"
	"testing"
	"time"

	"github.com/avasapollo/users-api/global"
	"github.com/avasapollo/users-api/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	le     *logrus.Entry
	conf   *global.AppConfig
	db     Database
	u      *users.User
	userID string
)

func TestMain(m *testing.M) {
	le = logrus.WithField("service", "testing")
	conf = &global.AppConfig{}

	userID = uuid.New().String()
	u = users.NewUser(
		userID,
		"rayhack",
		"andrea",
		"vasapollo",
		time.Now())

	db = NewMemoryDB(le, conf)
	os.Exit(m.Run())
}

func TestMemoryDB_CreateUser(t *testing.T) {
	err := db.CreateUser(u)
	assert.Nil(t, err)
}

func TestMemoryDB_GetUser(t *testing.T) {
	user, err := db.GetUser(u.ID)
	assert.Nil(t, err)
	assert.Equal(t, user, u)
}

func TestMemoryDB_UpdateUser(t *testing.T) {
	u.Nickname = "rayhack1"
	err := db.UpdateUser(u)
	assert.Nil(t, err)
	user, err := db.GetUser(userID)
	assert.Nil(t, err)
	assert.Equal(t, u, user)
}

func TestMemoryDB_DeleteUser(t *testing.T) {
	err := db.DeleteUser(u.ID)
	assert.Nil(t, err)
	user, err := db.GetUser(userID)
	t.Log(user)
	assert.Equal(t, err, UserNotFound)
}
