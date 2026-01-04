package color

import (
	"context"
	"sync"
	"time"
)

var confCache sync.Map

func (s *UCManager) InitCache() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	confs, err := s.GetAllColorConfs(ctx)
	if err != nil {
		return err
	}

	cmap := make(map[string][]string)
	for _, conf := range confs {
		if cmap[conf.ColorSystem] == nil {
			cmap[conf.ColorSystem] = make([]string, 0)
		}
		cmap[conf.ColorSystem] = append(cmap[conf.ColorSystem], conf.ColorNumber)
	}

	// Store all color systems in sync.Map
	for colorSystem, numbers := range cmap {
		confCache.Store(colorSystem, numbers)
	}
	return nil
}

func (s *UCManager) GetColorConfByCache(ctx context.Context, colorSystem string) []string {
	if val, ok := confCache.Load(colorSystem); ok {
		return val.([]string)
	}
	return nil
}
