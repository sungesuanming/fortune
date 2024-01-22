package color

import (
	"fortune/handler"
	"github.com/gin-gonic/gin"
)

type TestResp struct {
	BestColor      string `json:"best_color"`
	AlternateColor string `json:"alternate_color"`
	WorstColor     string `json:"worst_color"`
}

func ColorTest(c *gin.Context) {
	res := TestResp{
		BestColor:      "红色",
		AlternateColor: "黄色",
		WorstColor:     "蓝色",
	}
	handler.SendResponse(c, nil, res)
}
