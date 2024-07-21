const inputBox = document.getElementById('inputBox');
const outputBox = document.getElementById('outputBox');
const copyButton = document.getElementById('copyButton');

inputBox.addEventListener('input', () => {
    outputBox.innerHTML = inputBox.innerHTML;
});

copyButton.addEventListener('click', () => {
    const range = document.createRange();
    range.selectNode(outputBox);
    window.getSelection().removeAllRanges();
    window.getSelection().addRange(range);
    try {
        document.execCommand('copy');
    } catch (err) {
        alert('Erro ao copiar texto');
    }
    window.getSelection().removeAllRanges();
});