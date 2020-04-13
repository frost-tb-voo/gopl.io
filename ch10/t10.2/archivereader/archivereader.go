package archivereader

import (
	"errors"
	"sync"
	"sync/atomic"
)

type FileContent struct {
	Name string
	Body []byte
}

var ErrFormat = errors.New("archive: unknown format")
var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

type format struct {
	name    string
	match   func(filePath string) error
	extract func(filePath string) ([]FileContent, error)
}

func RegisterFormat(name string, match func(filePath string) error, extract func(filePath string) ([]FileContent, error)) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, match, extract}))
	formatsMu.Unlock()
}

func Read(filePath string) ([]FileContent, error) {
	ff := sniff(filePath)
	if ff.extract == nil {
		return nil, ErrFormat
	}
	return ff.extract(filePath)
}

func sniff(filePath string) format {
	formats, _ := atomicFormats.Load().([]format)
	for _, f := range formats {
		if f.match(filePath) == nil {
			return f
		}
	}
	return format{}
}
