function setCubeSize() {
    const aspectRatio = 4 / 3;
    
    const box = document.getElementById("cube-box");
    const boxWidth = box.offsetWidth;   
    const boxHeight = box.offsetHeight;
    
    const cubeSideSize = boxWidth/boxHeight > aspectRatio ? boxHeight/3 : boxWidth/4;
    
    console.log(box);
    console.log(boxWidth);
    console.log(boxHeight);
    console.log(Math.floor(cubeSideSize * 4).toString(), Math.floor(cubeSideSize * 3));
    
    
    const cube = document.getElementById("cube");
    // cube.setAttribute("style","width:" + Math.floor(cubeSideSize * 4).toString() + "px");
    // cube.setAttribute("style","height:" + Math.floor(cubeSideSize * 3).toString() +"px");
    
    cube.style.width  = Math.floor(cubeSideSize * 4).toString() + "px";
    cube.style.height = Math.floor(cubeSideSize * 3).toString() + "px";
}

setCubeSize();
window.onresize = setCubeSize;


const TOTAL_STICKERS = 24;
const TOTAL_COLORS = 6;

// get info about all current cullors
colorsIndexes = {};
const colorButtons = document.getElementsByName("colorSelection");
console.log(colorButtons)
for (let index = 0; index < colorButtons.length; index++) {
    colorsIndexes[colorButtons[index].value] = index
}
console.log(colorsIndexes)


currentStickers = Array(TOTAL_STICKERS).fill(-1);
console.log(currentStickers);

// collect all cube's stickers into array and add event listener for each
for (let stickerIndex = 0; stickerIndex < TOTAL_STICKERS; stickerIndex++) {
    console.log("sticker-" + stickerIndex.toString())

    const stickerButton = document.getElementById("sticker-" + stickerIndex.toString());
    stickerButton.addEventListener('click', function stickerButtonListener() {
        const newColor = document.querySelector('input[name="colorSelection"]:checked').value;
        stickerButton.style.backgroundColor = newColor
        currentStickers[stickerIndex] = colorsIndexes[newColor]
        // console.log(currentStickers)
    });
}
// console.log(currentStickers)
