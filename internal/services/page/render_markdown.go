package page

import "io"

func (s *service) RenderMarkdown(wr io.Writer, domain string, name string) error {
	data, err := s.markdown.ReadFile(domain, name)
	if err != nil {
		s.log.Err(err).E("unable to render markdown")
	}

	if err := s.markdown.Convert(data, wr); err != nil {
		s.log.Err(err).E("unable to render markdown")
	}

	return nil
}
