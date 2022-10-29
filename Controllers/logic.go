package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zxc10110/mvc_11/Models"
)

func GetPlagiatism(c *gin.Context) {
	student01, er := Models.GetPlagiarism01()
	if er != nil {
		c.JSON(http.StatusOK, "ไม่พบข้อมูล")
		return
	}

	student02, err := Models.GetPlagiarism02()
	if err != nil {
		c.JSON(http.StatusOK, "ไม่พบข้อมูล")
		return
	}

	result := Models.GetPlagiatism{
		Student01: student01,
		Student02: student02,
	}

	c.JSON(http.StatusOK, result)
	return
}
