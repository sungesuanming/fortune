package color

import (
	"context"
	"fortune/db"
	"fortune/model/color"
	"fortune/pkg/log"
)

type ColorHandler struct {
	*color.UCManager
}

var colorHandler *ColorHandler

func NewColorHandler(ctx context.Context, mysqlDB *db.MysqlDB) error {
	model, err := color.NewModel(ctx, mysqlDB)
	if err != nil {
		log.Errorf("NewColorHandler NewModel:%v", err)
		return err
	}
	err = model.InitCache()
	if err != nil {
		log.Errorf("NewColorHandler InitCache:%v", err)
		return err
	}
	colorHandler = &ColorHandler{model}
	return nil
}
