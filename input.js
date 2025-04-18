let pressKey = (key) => {
    keysDown.add(key);
    socket.send(convertKeyPressesToByte(keysDown));
}

let releaseKey = (key) => {
    keysDown.delete(key);
    var input = convertKeyPressesToByte(keysDown);
    socket.send(input);
};

let setupKeyListeners = () => {
    addEventListener("keydown", (e) => {
        pressKey(e.key);
    }, false);

    addEventListener("keyup", (e) => {
        releaseKey(e.key);
    }, false)
};

