package db_test

import (
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/backstopmedia/gRPC-book-example/server/db"
)

type ErrorsTest struct {
	suite.Suite
}

func TestErrors(t *testing.T) {
	suite.Run(t, new(ErrorsTest))
}

func (assert *ErrorsTest) TestNotFound() {
	err := db.NotFound("thing")
	assert.EqualError(err, "could not find thing")
}
