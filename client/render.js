
let renderPlayer = (viewportX, viewportY, x, y, direction, context, image, animationFlag) => {
  context.drawImage(image, 0+20*animationFlag,0+30*direction, 20,30,-viewportX+x-5, -viewportY+y-20,30,30);
}
