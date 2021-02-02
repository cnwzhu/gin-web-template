package endpoint

import (
	"github.com/gin-gonic/gin"
	"jet.com/infrared/models"
)

// index godoc
// @Summary 健康状态
// @Produce json
// @Success 200 {object} models.Index
// @Router /api/v1/app/index [get]
func Index(c *gin.Context) {
	c.JSON(200, models.Ok(models.Index{Status: "up"}))
}
