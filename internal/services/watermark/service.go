package watermark

import (
	"image"
	"lazts/internal/modules/imaging"
	"lazts/internal/modules/log"
	"sync"
)

type Servicer interface {
	LoadImage(filePath string) (image.Image, error)
}

type service struct {
	config *config
	logger log.Moduler
	imager imaging.Moduler
	cache  map[string]image.Image
	mutex  *sync.Mutex
}

var _ Servicer = (*service)(nil)

func New(c *config, logger log.Moduler, imager imaging.Moduler) *service {
	return &service{
		config: c,
		logger: logger,
		imager: imager,
		cache:  make(map[string]image.Image),
		mutex:  &sync.Mutex{},
	}
}
