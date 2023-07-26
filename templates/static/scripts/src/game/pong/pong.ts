import * as PIXI from 'pixi.js';

export default () => {
    const app = new PIXI.Application({
        background: '#1099bb',
        resizeTo: window,
    });
    
    document.body.appendChild(app.view as any);
    
    let ball = new PIXI.Graphics();
    ball.beginFill(0xFFFFFF);
    ball.drawCircle(0, 0, 10);
    ball.endFill();
    ball.x = app.view.width / 2;
    ball.y = app.view.height / 2;
    app.stage.addChild(ball);
    
    let paddleWidth = 15;
    let paddleHeight = 100;
    
    let leftPaddle = new PIXI.Graphics();
    leftPaddle.beginFill(0xFFFFFF);
    leftPaddle.drawRect(0, 0, paddleWidth, paddleHeight);
    leftPaddle.endFill();
    leftPaddle.x = 50;
    leftPaddle.y = app.view.height / 2 - paddleHeight / 2;
    app.stage.addChild(leftPaddle);
    
    let rightPaddle = new PIXI.Graphics();
    rightPaddle.beginFill(0xFFFFFF);
    rightPaddle.drawRect(0, 0, paddleWidth, paddleHeight);
    rightPaddle.endFill();
    rightPaddle.x = app.view.width - 50;
    rightPaddle.y = app.view.height / 2 - paddleHeight / 2;
    app.stage.addChild(rightPaddle);
    
    let ballDirection = { x: 3, y: 3 };
    
    app.ticker.add(delta => {
        ball.x += ballDirection.x * delta;
        ball.y += ballDirection.y * delta;
    
        // reverse direction when ball hits the wall
        if (ball.y < 0 || ball.y > app.view.height) {
            ballDirection.y *= -1;
        }
    });
    
    let keys: { [key: string]: boolean } = {};
    
    window.addEventListener('keydown', (key) => {
        keys[key.key] = true;
    });
    
    window.addEventListener('keyup', (key) => {
        keys[key.key] = false;
    });
    
    app.ticker.add(() => {
        if (keys['w'] && leftPaddle.y > 0) {
            leftPaddle.y -= 7;
        } else if (keys['s'] && leftPaddle.y < app.view.height - paddleHeight) {
            leftPaddle.y += 7;
        }
    
        if (keys['ArrowUp'] && rightPaddle.y > 0) {
            rightPaddle.y -= 7;
        } else if (keys['ArrowDown'] && rightPaddle.y < app.view.height - paddleHeight) {
            rightPaddle.y += 7;
        }
    });
    app.ticker.add(() => {
        if (ball.x < leftPaddle.x + paddleWidth
            && ball.y > leftPaddle.y
            && ball.y < leftPaddle.y + paddleHeight) {
            ballDirection.x *= -1;
        } else if (ball.x > rightPaddle.x
            && ball.y > rightPaddle.y
            && ball.y < rightPaddle.y + paddleHeight) {
            ballDirection.x *= -1;
        }
    
        // Reset when ball hits left or right wall
        if (ball.x < 0) {
            ball.x = app.view.width / 2;
            ball.y = app.view.height / 2;
        } else if (ball.x > app.view.width) {
            ball.x = app.view.width / 2;
            ball.y = app.view.height / 2;
        }
    });
}