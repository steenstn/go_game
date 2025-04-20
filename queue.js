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
