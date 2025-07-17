class Arm {
    constructor(startX, startY) {
        this.segments = [
            new ArmSegment(50),
            new ArmSegment(50),
            new ArmSegment(50),
            new ArmSegment(50)
        ]
        this.t=0;
        this.target = {x: startX, y: startY};
        this.target2 = {x: this.target.x, y: this.target.y};
        this.rise = 60;
        this.reach = 120;
        this.p0 = {x: this.target2.x, y:this.target2.y};
        this.p1 = {x: this.target2.x, y:this.target2.y};
        this.p2 = {x: this.target2.x, y:this.target2.y};
    }
};

class ArmSegment {
    constructor(length) {
        this.x = 0;
        this.y = 0;
        this.length = length;
    }
};

let moveUp = (arms) => {
  for(let i = 0; i < arms.length; i++) {
      arms[i].y-=2;
  }
}

let moveDown = (arms) => {
  for(let i = 0; i < arms.length; i++) {
      arms[i].y+=2;
  }
}

/*
 Interpolate between p0, p1, p2 with with t = 0-1
 */
let bezier = (t, p0, p1, p2) => {
  return (1-t)*((1-t)*p0 + t*p1) + t*((1-t)*p1 +t*p2);
}


