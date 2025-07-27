// https://www.researchgate.net/publication/220632147_FABRIK_A_fast_iterative_solver_for_the_Inverse_Kinematics_problem

let armLength = 50;

let getDistancesBetweenPoints = (points) => {
  let distances = []
  for(let i = 0; i < points.length-1; i++) {
    distances.push(distance(points[i], points[i+1]));
  }
  return distances;
}

let easeOutQuad = (progress) => {
    return 1 - (1 - progress) * (1 - progress);
}


let lerp = (startValue, endValue, progress) => {
    return startValue + (endValue - startValue) * progress
}

let fabrik = (positions, distancesBetweenJoints, target) => {
  let distanceToTarget = distance(positions[0], target);
  let jointDistanceSum = distancesBetweenJoints.reduce( (a,b) => {return a+b}, 0);
  if (distanceToTarget > jointDistanceSum) { // Target unreachable
    for(let i = 0; i < positions.length-1; i++) {
      let r = distance(target, positions[i]);
      let delta = distancesBetweenJoints[i]/r;
      positions[i+1].x = (1-delta)*positions[i].x+delta*target.x;
      positions[i+1].y = (1-delta)*positions[i].y+delta*target.y;
    }
  } else {
    let bx = positions[0].x;
    let by = positions[0].y;
    let diff = distance(positions[positions.length-1], target);
    for(let iter = 0; iter < 3; iter++) {
      // Stage 1: Forward reaching
      positions[positions.length-1].x = target.x;
      positions[positions.length-1].y = target.y;

      for(let i = positions.length-2; i >=0; i--) {
        let r = distance(positions[i+1], positions[i]);
        let delta = distancesBetweenJoints[i]/r;
        positions[i].x = (1-delta)*positions[i+1].x+delta*positions[i].x;
        positions[i].y = (1-delta)*positions[i+1].y+delta*positions[i].y;
      }

      // Stage 2: Backward reaching
      positions[0].x = bx;
      positions[0].y = by;
    
      for(let i = 0; i < positions.length-2; i++) {
        let r = distance(positions[i+1], positions[i]);
        let delta = distancesBetweenJoints[i]/r;
        positions[i+1].x = (1-delta)*positions[i].x + delta*positions[i+1].x;
        positions[i+1].y = (1-delta)*positions[i].y + delta*positions[i+1].y;
      }
      diff = distance(positions[positions.length-1], target);
    }
  }

}

class Spider2 {
  constructor(startX, startY) {
    this.arms = [];
    this.distances = [];
    let numArms = 3
    let arm_length = 30;
    for(let i = 0; i < numArms; i++) {
      let arm = [];
      for(let j =0; j<5;j++) {
        arm.push({x:startX+40*i+j*arm_length,y:startY+j*arm_length});
      }
      this.arms.push(arm);
      this.distances.push(getDistancesBetweenPoints(arm))
    }
    this.body = this.arms[0][0];
  }
}

class Spider {
  constructor(startX, startY) {
    this.arms = [
      new Arm(startX, startY),
      new Arm(startX+50, startY),
      new Arm(startX+100, startY),
      new Arm(startX+150, startY),
    ]

    for(let i = 0; i < this.arms.length; i++) {
      this.arms[i].segments[0].x = 200;
      this.arms[i].segments[0].y = 300;
    }

  }

  x() {
    return this.arms[0].segments[0].x;
  }

  y() {
    return this.arms[0].segments[0].y;
  }

  move(x, y) {
    for(let i = 0; i < this.arms.length; i++) {
      let arm = this.arms[i];
      for(let j = 0; j < arm.segments.length; j++) {
        arm.segments[j].x+=x;
        arm.segments[j].y+=y;
      }
    }
  }

  move2(x, y) {
    for(let i = 0; i < this.arms.length; i++) {
      let arm = this.arms[i];
        let oldX = arm.segments[0].x;
        let oldY = arm.segments[0].y;
        arm.segments[0].x=x;
        arm.segments[0].y=y;
      for(let j = 1; j < arm.segments.length; j++) {
        arm.segments[j].x+=arm.segments[0].x-oldX;
        arm.segments[j].y+=arm.segments[0].y-oldY;
      }
      }
  }

  update() {
    for(let armIndex = 0; armIndex < this.arms.length; armIndex++) {
        let arm = this.arms[armIndex];
        let arms = arm.segments;
        arm.target2 = arms[arms.length-1];
        arm.target = arms[arms.length-1];
        if(Math.abs(arms[0].x - arm.p2.x) > 50) {
            arm.target2.x = arms[0].x -80;
            arm.p0 = {x: arm.target.x, y: arm.target.y};
            arm.p1 = {x: arm.p0.x-arm.reach/2, y: arm.p0.y-arm.rise};
            arm.p2 = {x: arm.p0.x-arm.reach, y: arm.target2.y};
            arm.t=0;
        }
        if(Math.abs(arms[0].x - arm.p2.x) > 50) {
            arm.target2.x = arms[0].x - 80;
            arm.p0 = {x: arm.target.x, y: arm.target.y};
            arm.p1 = {x: arm.p0.x+arm.reach/2, y: arm.p0.y-arm.rise};
            arm.p2 = {x: arm.p0.x+arm.reach, y: arm.target2.y};
            arm.t=0;
        }

        ctx.fillStyle = "#000";

        let bx = bezier(arm.t, arm.p0.x, arm.p1.x, arm.p2.x);
        let by = bezier(arm.t, arm.p0.y, arm.p1.y, arm.p2.y);

        arm.target.x = bx;
        arm.target.y = by;
        if(arm.t <1) {
            arm.t+=0.1;
        }

        let dist = distance(arm.target, arms[0]);
        if(dist > (arms.length)*armLength) {
            for(let i = 0; i < arms.length - 1; i++) {
                let r = distance(arm.target, arms[i]);
                let delta = armLength/r;
                arms[i+1].x = (1-delta) * arms[i].x + delta*arm.target.x;
                arms[i+1].y = (1-delta) * arms[i].y + delta*arm.target.y;
            }
        } else {
            let bx = arms[0].x;
            let by = arms[0].y;
            let max = 500;
            let iterations = 0;
            let diff = distance(arms[arms.length-1], arm.target);
            while (diff > 2 && iterations < max) {
                iterations++;
                arms[arms.length-1].x = arm.target.x;
                arms[arms.length-1].y = arm.target.y;
                for(let i = arms.length-2; i >=0; i--) {
                    let r = distance(arms[i+1], arms[i]);
                    let delta = armLength/r;
                    arms[i].x = (1-delta)*arms[i+1].x + delta*arms[i].x;
                    arms[i].y = (1-delta)*arms[i+1].y + delta*arms[i].y;
                }

                arms[0].x = bx;
                arms[0].y = by;
                for(let i = 0; i < arms.length-1; i++) {
                    let r = distance(arms[i+1], arms[i]);
                    let delta = armLength/r;
                    arms[i+1].x = (1-delta)*arms[i].x + delta*arms[i+1].x;
                    arms[i+1].y = (1-delta)*arms[i].y + delta*arms[i+1].y;
                }
                diff = distance(arms[arms.length-1], arm.target);
            }

        }

        // Rotation check
/*
        for(let i=0; i < arms.length-2; i++) {
            let angle = Math.atan2(arms[i+1].y-arms[i].y, arms[i+1].x-arms[i].x);
            let angle2 = Math.atan2(arms[i+2].y-arms[i].y, arms[i+2].x-arms[i].x);
            ctx.fillText(""+ (angle*(180/Math.PI)), arms[i].x+20, arms[i].y )
            ctx.fillText(""+ (angle2*(180/Math.PI)), arms[i].x+20, arms[i].y+20 )
        }
    */


        if(true) {

            ctx.fillStyle= "#0f0";
            ctx.fillRect(arm.target.x, arm.target.y, 10,10);
            ctx.fillStyle= "#060";
            ctx.fillRect(arm.target2.x, arm.target2.y, 10,10);
            ctx.fillStyle= "#06f";
            ctx.fillText("arm.target2", arm.target2.x, arm.target2.y);
        }
    }
  }

  render(ctx, offsetX, offsetY) {
    for(let armIndex = 0; armIndex < this.arms.length; armIndex++) {
        let arm = this.arms[armIndex];
        let arms = arm.segments;
        for(let i=0; i < arms.length-1; i++) {
            ctx.fillStyle = "#a00";
            ctx.fillRect(offsetX+arms[i].x,offsetY+arms[i].y,5,5);
            ctx.fillStyle= "#000";
            ctx.fillText(""+ i, offsetX+arms[i].x, offsetY+arms[i].y )
            ctx.beginPath();
            ctx.moveTo(offsetX+arms[i].x, offsetY+arms[i].y);
            ctx.lineTo(offsetX+arms[i+1].x, offsetY+arms[i+1].y);
            ctx.stroke();
        }
    }
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


let distance = (a, b) => {
    return Math.sqrt((a.x - b.x)*(a.x-b.x) + (a.y - b.y)*(a.y-b.y));
};

/*
 Interpolate between p0, p1, p2 with with t = 0-1
 */
let bezier = (t, p0, p1, p2) => {
  return (1-t)*((1-t)*p0 + t*p1) + t*((1-t)*p1 +t*p2);
}


