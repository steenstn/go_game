let convertKeyPressesToByte = (input) => {
   let buttonsPressed =
        (1 & (input.has("ArrowUp")? 1 : 0) ) |
        (2 & (input.has("ArrowDown")? 2 : 0)) |
        (4 & (input.has("ArrowLeft")? 4 : 0)) |
        (8 & (input.has("ArrowRight")? 8 : 0));
    return new Uint8Array([buttonsPressed]);
}

let getArrayIndex = (levelWidth, tileWidth, x, y) => {
    return Math.floor(x/tileWidth)+Math.floor(y/tileWidth)*levelWidth;
  }
