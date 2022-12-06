package controller

import (
	"demo/service"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context){
	service.TestA()
}

func TestB(c *gin.Context)  {
	service.TestB()
}

func TestC(c *gin.Context)  {
	service.TestC()
}

func TestD(c *gin.Context)  {
	service.TestD()
}