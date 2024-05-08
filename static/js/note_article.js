function getPathSegments() {
  const path = window.location.pathname;
  const segments = path.split('/').filter(segment => segment);
  return segments;
}

const segments = getPathSegments();
const tags = segments[1];
const slug = segments[2];
const section = document.querySelector('article[name="article"]');

htmx.on('htmx:responseError', function (event) {
  if (event.detail.xhr.status === 404) {
    const target = event.detail.target;
    target.innerHTML = 'Error: The requested article could not be found.';
  }
});

if (section && tags && slug) {
  htmx.ajax('GET', '/_notes/contents/' + tags + '/' + slug, { target: '#article', swap: 'innerHTML', setErrorHandler: true });
}