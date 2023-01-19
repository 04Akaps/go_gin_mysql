package err

import (
	"errors"

	"github.com/gin-gonic/gin"
)


var (
	ErrNotFoundConfigFile = errors.New("not found config file")
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error" : err.Error()}
}