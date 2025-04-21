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

