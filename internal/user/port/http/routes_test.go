package http

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"

	"main/pkg/dbs/mocks"
)

func TestRoutes(t *testing.T) {
	mockDB := mocks.NewIDatabase(t)
	Routes(gin.New().Group("/"), mockDB, validation.New())
}
