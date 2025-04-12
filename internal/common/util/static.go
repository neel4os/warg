package util

import (
	"embed"
	"sync"
)

var (
	instance *StaticFileLocation
	once     sync.Once
)

type StaticFileLocation struct {
	staticFiles embed.FS
}

func NewStaticFileLocation(s *embed.FS) *StaticFileLocation {
	once.Do(func() {
		instance = &StaticFileLocation{
			staticFiles: *s,
		}
	})
	return instance
}

func (s *StaticFileLocation) GetStaticFiles() embed.FS {
	return s.staticFiles
}
