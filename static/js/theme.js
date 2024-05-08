const theme = document.getElementById('theme-toggle');

theme.addEventListener('click', function () {
  document.documentElement.classList.toggle('dark');
  const isDark = document.documentElement.classList.contains('dark');
  localStorage.setItem('darkMode', isDark);
  if (isDark) {
    theme.classList.add('bg-transparent')
    theme.classList.remove('bg-orange-300')
    theme.style.boxShadow = "inset -9px -8px 1px 1px #ddd"
  } else {
    theme.classList.add('bg-orange-300')
    theme.classList.remove('bg-transparent')
    theme.style.boxShadow = ""
  }
});

// Set initial theme from local storage
if (localStorage.getItem('darkMode') === 'true') {
  document.documentElement.classList.add('dark');
  theme.classList.add('bg-transparent')
  theme.classList.remove('bg-orange-300')
  theme.style.boxShadow = "inset -9px -8px 1px 1px #ddd"
} else {
  document.documentElement.classList.remove('dark');
  theme.classList.add('bg-orange-300')
  theme.classList.remove('bg-transparent')
  theme.style.boxShadow = ""
}