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

// Circular array queue
class Queue {
    
    constructor(size) {
        this.size = size;
        this.data = new Array(size);
        this.tail = 0;
        this.head = 0;
        this.numEntries = 0;
    }

    // Push a value. Pushing past the capacity overwrites the oldest value
    push(value) {
        if(this.numEntries >= this.data.length) {
            this.head = (this.head + 1) % (this.data.length);
        }
        this.data[this.tail] = value;
        this.tail = (this.tail + 1) % this.data.length;
        this.numEntries++;
        return true;
    }

    pop() {
        if(this.numEntries <=0) {
            return null;
        }
        let value = this.data[this.head];
        this.head = (this.head + 1) % (this.data.length);
        this.numEntries--;
        return value;
    }
}
