package page

import "io"

func (s *service) RenderMarkdown(wr io.Writer, domain string, name string) error {
	s.logger.Fields("domain", domain, "name", name).I("rendered markdown")

	data, err := s.markdown.ReadFile(domain, name)
	if err != nil {
		s.logger.Err(err).E("unable to render markdown")
	}

	if err := s.markdown.Convert(data, wr); err != nil {
		s.logger.Err(err).E("unable to render markdown")
	}

	return nil
}
