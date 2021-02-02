package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"jet.com/infrared/logger"
)

func Middleware(r *gin.Engine) {
	r.Use(JSONAppErrorReporter())
	r.Use(ginlogrus.Logger(logger.Logger))
	r.Use(gin.Recovery())
}
