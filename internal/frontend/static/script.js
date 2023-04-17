const TOTAL_STICKERS = 24;
// const TOTAL_COLORS = 6;

// get info about all current colors
colorsIndexes = {};
const colorButtons = document.getElementsByName("colorSelection");
for (let index = 0; index < colorButtons.length; index++) {
    let color = getComputedStyle(colorButtons[index]).backgroundColor;
    colorsIndexes[color] = index
}

// collect all cube's stickers into array and add event listener for each
currentStickers = Array(TOTAL_STICKERS).fill(-1);
const stickers = document.getElementsByClassName("sticker");
for (let index = 0; index < stickers.length; index++) {
    stickers[index].onclick = function() {
        const selectedButton = document.querySelector('input[name="colorSelection"]:checked');
        const newColor = getComputedStyle(selectedButton).backgroundColor;
        stickers[index].style.backgroundColor = newColor;
        currentStickers[index] = colorsIndexes[newColor];
    };
}

const solution = document.getElementById("solution")

const solveButton = document.getElementById("solve-button");
solveButton.onclick = function() {
    console.log(currentStickers);

    const URL = "http://localhost:8080";
    const options = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({cube: currentStickers}),
    };

    fetch(URL, options)
        .then(response=>response.text())
        .then(data=>{ solution.innerText = data; })
};


