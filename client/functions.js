let convertKeyPressesToByte = (input) => {
   let buttonsPressed =
        (1 & (input.has("ArrowUp")? 1 : 0) ) |
        (2 & (input.has("ArrowDown")? 2 : 0)) |
        (4 & (input.has("ArrowLeft")? 4 : 0)) |
        (8 & (input.has("ArrowRight")? 8 : 0));
    return new Uint8Array([buttonsPressed]);
}

let getArrayIndex = (levelWidth, tileWidth, x, y) => {
    let xPos = Math.floor(x / tileWidth);
    let yPos = levelWidth * Math.floor(y/tileWidth);
    let res = Math.floor(xPos + yPos);
    if(res < 0) {
        res = 0;
    }
    return 0;
}
