

function getArrayIndex(levelWidth: number, tileWidth: number, x: number, y: number): number {
    return Math.floor(x/tileWidth)+Math.floor(y/tileWidth)*levelWidth;
}
