function getPathSegments() {
  const path = window.location.pathname;
  const segments = path.split('/').filter(segment => segment);
  return segments;
}

const segments = getPathSegments();
const slug = segments[1];
const section = document.querySelector('article[name="article"]');

htmx.on('htmx:responseError', function (event) {
  if (event.detail.xhr.status === 404) {
    const target = event.detail.target;
    target.innerHTML = 'Error: The requested article could not be found.';
  }
});

if (section && slug) {
  htmx.ajax('GET', '/_vacations/contents/' + slug, { target: '#article', swap: 'innerHTML', setErrorHandler: true });
}