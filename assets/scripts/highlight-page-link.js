window.addEventListener('DOMContentLoaded', (event) => {
    const pageLinks = document.querySelectorAll('.page-link');
    pageLinks.forEach((link) => {
        if (window.location.href.includes(link.href)) {
            link.classList.add('active');
        }
    });
});