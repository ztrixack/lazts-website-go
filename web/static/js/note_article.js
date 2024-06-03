function getPathSegments() {
  const path = window.location.pathname;
  const segments = path.split('/').filter(segment => segment);
  return segments;
}

const segments = getPathSegments();
const tags = segments[1];
const slug = segments[2];
const section = document.querySelector('article[name="article"]');

if (section && tags && slug) {
  setTimeout(function(){
    htmx.ajax('GET', '/_notes/headers/' + tags + '/' + slug, {
      target: '#title',
      swap: 'innerHTML',
      setErrorHandler: true,
    });
}, 100);
  htmx.ajax('GET', '/_notes/contents/' + tags + '/' + slug, {
    target: '#article',
    swap: 'innerHTML',
    setErrorHandler: true,
  });
} else {
  section.innerHTML = 'Error: The requested article could not be found.';
}
