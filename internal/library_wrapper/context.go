package library_wrapper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type MyGinContext interface {
	ShouldBindJSON(obj interface{}) error
	JSON(code int, obj interface{})
}

type MyGinContextImpl struct {
	*gin.Context
}

func (m *MyGinContextImpl) ShouldBindJSON(obj interface{}) error {
	fmt.Println("my own ShouldBindJSON")
	return m.Context.ShouldBindJSON(obj) // or entirely alternative implementation
}

func (m *MyGinContextImpl) JSON(code int, obj interface{}) {
	m.Context.JSON(code, obj)
}
