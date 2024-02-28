package color

import (
	"context"
	"fortune/db"
)

type UCManager struct {
	MysqlDB *db.MysqlDB
}

type DayMatch struct {
	Id          uint32 `gorm:"column:id"`
	UserDay     string `gorm:"column:user_day"`
	CurrentDay  string `gorm:"column:current_day"`
	Optimum     string `gorm:"column:optimum"`
	Alternative string `gorm:"column:alternative"`
	NoRecommend string `gorm:"column:no_recommend"`
}

type ColorConf struct {
	Id          uint32 `gorm:"column:id"`
	ColorSystem string `gorm:"column:color_system"`
	ColorNumber string `gorm:"column:color_number"`
}

func NewModel(ctx context.Context, mysqlDb *db.MysqlDB) (*UCManager, error) {
	return &UCManager{mysqlDb}, nil
}

func (s *UCManager) GetColorByUserAndDay(ctx context.Context, userDay, currentDay string) (*DayMatch, error) {
	var r *DayMatch
	err := s.MysqlDB.DB.WithContext(ctx).Model(&DayMatch{}).Where("user_day = ? and current_day = ?", userDay, currentDay).Find(&r).Error
	return r, err
}

func (s *UCManager) GetColorsConfBySystem(ctx context.Context, colorSystem string) ([]string, error) {
	r := make([]*ColorConf, 0)
	err := s.MysqlDB.DB.WithContext(ctx).Model(&ColorConf{}).Where("color_system = ?", colorSystem).Find(&r).Error
	if err != nil {
		return nil, err
	}

	res := make([]string, len(r))
	for i, conf := range r {
		res[i] = conf.ColorNumber
	}

	return res, err
}

func (s *UCManager) GetAllColorConfs(ctx context.Context) ([]*ColorConf, error) {
	r := make([]*ColorConf, 0)
	err := s.MysqlDB.DB.WithContext(ctx).Find(&r).Error
	if err != nil {
		return nil, err
	}

	return r, err
}
