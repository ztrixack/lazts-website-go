package watermark

import (
	"image"
)

func (s *service) LoadImage(filePath string) (image.Image, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if img, found := s.cache[filePath]; found {
		return img, nil
	}

	watermark, err := s.imager.Open(s.config.Path)
	if err != nil {
		s.logger.Err(err).Fields("path", s.config.Path).E("unable to open watermark")
		return nil, err
	}
	watermark = s.imager.Resize(watermark, s.config.Size, 0)

	original, err := s.imager.Open(filePath)
	if err != nil {
		s.logger.Err(err).Fields("path", filePath).E("unable to open image")
		return nil, err
	}
	offset := image.Pt(original.Bounds().Dx()-watermark.Bounds().Dx()-10, original.Bounds().Dy()-watermark.Bounds().Dy()-10)
	dst := s.imager.Overlay(original, watermark, offset, 1.0)

	s.cache[filePath] = dst
	return dst, nil
}
