initFilter();

document.querySelector('input[name="search"]').addEventListener('keyup', updateFilter);
document.querySelector('input[name="search"]').addEventListener('change', updateFilter);
document.querySelector('select[name="catalog"]').addEventListener('change', updateFilter);

function initFilter() {
  const urlParams = new URLSearchParams(location.search);

  document.querySelector('input[name="search"]').value = urlParams.get('search') || '';
  document.querySelector('select[name="catalog"]').value = urlParams.get('catalog') || '';
}

function updateFilter() {
  const filters = {
    search: document.querySelector('input[name="search"]').value,
    catalog: document.querySelector('select[name="catalog"]').value,
  };
  const cleanFilters = Object.fromEntries(Object.entries(filters).filter(([_, v]) => v != null && v !== ''));
  const queryParams = new URLSearchParams(cleanFilters).toString();

  history.pushState(null, '', queryParams ? '/books?' + queryParams : '/books');
  htmx.ajax('GET', queryParams ? '/_books/list?' + queryParams : '/_books/list', { target: '#book-list', swap: 'innerHTML' });
}
