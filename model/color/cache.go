package color

import (
	"context"
	"sync"
	"time"
)

var confCache = struct {
	confMap map[string][]string
	sync.RWMutex
}{
	confMap: make(map[string][]string, 0),
}

func (s *UCManager) InitCache() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	confs, err := s.GetAllColorConfs(ctx)
	if err != nil {
		return err
	}

	cmap := make(map[string][]string, 0)
	for _, conf := range confs {
		if len(cmap[conf.ColorSystem]) == 0 {
			cmap[conf.ColorSystem] = make([]string, 0)
		}
		cmap[conf.ColorSystem] = append(cmap[conf.ColorSystem], conf.ColorNumber)
	}

	confCache.Lock()
	confCache.confMap = cmap
	confCache.Unlock()
	return nil
}

func (s *UCManager) GetColorConfByCache(ctx context.Context, colorSystem string) []string {
	confCache.RLock()
	defer confCache.RUnlock()
	res := confCache.confMap[colorSystem]
	return res
}
