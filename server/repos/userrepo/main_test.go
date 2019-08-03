package userrepo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"outstagram/server/db"
	"outstagram/server/models"
	"outstagram/server/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository UserRepo
	person     *model.Person
}
