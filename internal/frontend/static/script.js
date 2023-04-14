function setCubeSize() {
    const aspectRatio = 4 / 3;
    
    const box = document.getElementById("cube-box");
    const boxWidth = box.offsetWidth;   
    const boxHeight = box.offsetHeight;
    
    const cubeSideSize = boxWidth/boxHeight > aspectRatio ? boxHeight/3 : boxWidth/4;

    const cube = document.getElementById("cube");
    cube.style.width  = Math.floor(cubeSideSize * 4).toString() + "px";
    cube.style.height = Math.floor(cubeSideSize * 3).toString() + "px";
}

setCubeSize();
window.onresize = setCubeSize;


const TOTAL_STICKERS = 24;
// const TOTAL_COLORS = 6;

// get info about all current colors
colorsIndexes = {};
const colorButtons = document.getElementsByName("colorSelection");
console.log(colorButtons)
for (let index = 0; index < colorButtons.length; index++) {
    let color = getComputedStyle(colorButtons[index]).backgroundColor;
    colorsIndexes[color] = index
}
console.log(colorsIndexes)


currentStickers = Array(TOTAL_STICKERS).fill(-1);
console.log(currentStickers);

// collect all cube's stickers into array and add event listener for each
const stickers = document.getElementsByClassName("sticker");
for (let index = 0; index < stickers.length; index++) {
    stickers[index].addEventListener('click', function() {
        const selectedButton = document.querySelector('input[class="color-selection"]:checked');
        const newColor = getComputedStyle(selectedButton).backgroundColor;
        stickers[index].style.backgroundColor = newColor;
        currentStickers[index] = colorsIndexes[newColor];
        console.log(currentStickers);
    });
}

// for (let stickerIndex = 0; stickerIndex < TOTAL_STICKERS; stickerIndex++) {
//     console.log("sticker-" + stickerIndex.toString())
//
//     const stickerButton = document.getElementById("sticker-" + stickerIndex.toString());
//     stickerButton.addEventListener('click', function() {
//         const selectedButton = document.querySelector('input[name="colorSelection"]:checked');
//         const newColor = getComputedStyle(selectedButton).backgroundColor;
//         stickerButton.style.backgroundColor = newColor;
//         currentStickers[stickerIndex] = colorsIndexes[newColor];
//         console.log(currentStickers);
//     });
// }
// console.log(currentStickers)

const solution = document.getElementById("solution")

const solveButton = document.getElementById("solve-me-button");
solveButton.addEventListener('click', function() {
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
});


