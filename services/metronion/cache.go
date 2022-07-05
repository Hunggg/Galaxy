package metronion

import "go.uber.org/zap"

type MetronionCache struct {
	l            *zap.SugaredLogger
	persistentDb PersistentDb
}

func NewMetronionCache(persistentDb PersistentDb) *MetronionCache {
	l := zap.S()
	return &MetronionCache{
		l:            l,
		persistentDb: persistentDb,
	}
}