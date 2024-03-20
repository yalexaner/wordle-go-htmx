document.body.addEventListener('htmx:load', function (evt) {
    console.log('load');
    makeFirstLetterFocused();
    attachLetterListeners();
});

function makeFirstLetterFocused() {
    const firstLetter = document.querySelector('.letter[data-letter="0"][required]');
    firstLetter.focus();
}

function attachLetterListeners() {
    document.getElementById('board')
        .querySelectorAll('.letter[required]')
        .forEach(letter => {
            letter.addEventListener('input', inputEventListener);
            letter.onkeydown = onKeyDownListener(letter);
        })

    document.querySelector('.submit-btn')
        .onkeydown = function (e) {
        if (e.key === 'Backspace') {
            const nextLetter = document.querySelector(`.letter[required][data-letter="4"]`);
            nextLetter.focus();
        }
    }
}

function inputEventListener() {
    const letterIndex = parseInt(this.dataset.letter);
    if (this.value !== "" && letterIndex < 4) {
        const nextLetter = document.querySelector(`.letter[required][data-letter="${letterIndex + 1}"]`);
        nextLetter.focus();
    } else if (this.value !== "" && letterIndex === 4) {
        document.querySelector('.submit-btn').focus();
    }
}

function onKeyDownListener(letter) {
    return function (e) {
        if (e.key === 'Backspace' && letter.value === "") {
            const letterIndex = parseInt(this.dataset.letter);
            if (letterIndex > 0) {
                const nextLetter = document.querySelector(`.letter[required][data-letter="${letterIndex - 1}"]`);
                nextLetter.focus();
            }
        }
    }
}
