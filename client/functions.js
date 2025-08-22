"use strict";
function getArrayIndex(levelWidth, tileWidth, x, y) {
    return Math.floor(x / tileWidth) + Math.floor(y / tileWidth) * levelWidth;
}
