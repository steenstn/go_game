let convertKeyPressesToByte = (input) => {
   let buttonsPressed =
        (1 & (input.has("ArrowUp")? 1 : 0) ) |
        (2 & (input.has("ArrowDown")? 2 : 0)) |
        (4 & (input.has("ArrowLeft")? 4 : 0)) |
        (8 & (input.has("ArrowRight")? 8 : 0));
    return new Uint8Array([buttonsPressed]);
}

let sendKeys = () =>  {
    let input = convertKeyPressesToByte(keysDown);
    let data = {Num: requestNumber, Input: input[0]};
    inputHistory.push(data);
    requestNumber++;

    socket.send(JSON.stringify(data));
}

let pressKey = (key) => {
    keysDown.add(key);
    sendKeys();
}

let releaseKey = (key) => {
    keysDown.delete(key);
    sendKeys();
};

let setupKeyListeners = () => {
    addEventListener("keydown", (e) => {
        pressKey(e.key);
    }, false);

    addEventListener("keyup", (e) => {
        releaseKey(e.key);
    }, false)
};

