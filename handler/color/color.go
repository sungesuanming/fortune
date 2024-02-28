package color

import (
	"fortune/handler"
	"fortune/pkg/log"
	"github.com/gin-gonic/gin"
)

type TestReq struct {
	BestColor      string `json:"best_color"`
	AlternateColor string `json:"alternate_color"`
	WorstColor     string `json:"worst_color"`
}

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

type TodayReq struct {
	UserDay    string `json:"user_day"`
	CurrentDay string `json:"current_day"`
}

type TodayResp struct {
	BestColor      string `json:"best_color"`
	AlternateColor string `json:"alternate_color"`
	WorstColor     string `json:"worst_color"`
}

func TodayColor(c *gin.Context) {
	req := &TodayReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("TodayColor ShouldBindJSON error:%v", err)
		return
	}

	colorResult, err := colorHandler.GetColorByUserAndDay(c, req.UserDay, req.CurrentDay)
	if err != nil {
		log.Errorf("TodayColor GetColorByUserAndDay error:%v", err)
		return
	}

	res := TodayResp{
		BestColor:      colorResult.Optimum,
		AlternateColor: colorResult.Alternative,
		WorstColor:     colorResult.NoRecommend,
	}
	handler.SendResponse(c, nil, res)
}
