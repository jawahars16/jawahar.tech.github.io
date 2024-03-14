window.addEventListener('DOMContentLoaded', (event) => {
    const copyLinkButtons = document.querySelectorAll('.copy-link');
    copyLinkButtons.forEach((button) => {
        button.addEventListener('click', () => {
            const href = window.location.href.split('#')[0];
            const link = href + "#" + button.id
            navigator.clipboard.writeText(link);
        });
    });
});