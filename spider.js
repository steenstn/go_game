// https://www.researchgate.net/publication/220632147_FABRIK_A_fast_iterative_solver_for_the_Inverse_Kinematics_problem

class Spider {
  constructor(startX, startY) {
    this.arms = [new Arm(startX, startY),
      new Arm(startX+50, startY),
      new Arm(startX+100, startY),
    ]
  }
}

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

let distance = (a, b) => {
    return Math.sqrt((a.x - b.x)*(a.x-b.x) + (a.y - b.y)*(a.y-b.y));
};
/*
 Interpolate between p0, p1, p2 with with t = 0-1
 */
let bezier = (t, p0, p1, p2) => {
  return (1-t)*((1-t)*p0 + t*p1) + t*((1-t)*p1 +t*p2);
}


