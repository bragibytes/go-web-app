import * as PIXI from 'pixi.js';
import { TILE_SIZE as ts } from './game/config';

let socket = new WebSocket("ws://localhost:10000/ws");
let keys: { [key: string]: boolean } = {};
window.addEventListener('keydown', (key) => {
    keys[key.key] = true;
});

socket.onopen = function(e) {
  alert("[open] Connection established");
  alert("Sending to server");
  socket.send("Test message from client");
};

socket.onmessage = function(event) {
  alert(`[message] Data received from server: ${event.data}`);
};

socket.onclose = function(event) {
  if (event.wasClean) {
    alert(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
  } else {
    // e.g. server process killed or network down
    alert('[close] Connection died');
  }
};

socket.onerror = function(error) {
  alert(`[error] ${error}`);
};

let emptyTile = new PIXI.Graphics();
emptyTile.beginFill(0xFFFFFF);
emptyTile.drawRect(0, 0, ts, ts);
emptyTile.endFill();

let solidTile = new PIXI.Graphics();
solidTile.beginFill(0x000000);
solidTile.drawRect(0, 0, ts, ts);
solidTile.endFill();

class World {
    app: PIXI.Application;
    map = [
        [0, 1, 0, 0, 0],
        [0, 1, 0, 1, 0],
        [0, 0, 0, 1, 0],
        [0, 1, 0, 0, 0],
        [0, 0, 0, 1, 0]
    ];
    enteties = []
    constructor(app: PIXI.Application) {
        for(let y = 0; y < this.map.length; y++) {
            for(let x = 0; x < this.map[y].length; x++) {
                let tile;
                if(this.map[y][x] === 0) {
                    tile = emptyTile.clone();
                } else {
                    tile = solidTile.clone();
                }
                tile.x = x * ts;
                tile.y = y * ts;
                app.stage.addChild(tile);

                app.ticker.add(() => {
                    switch(true){
                        case keys["w"]:
                            

                    }
                })
            }
        }
        this.app = app;
    }
    draw() {
        // for(let i = 0; i < this.enteties.length; i++) {
        //     this.enteties[i].draw(this.app);
        // }
    }
    update() {
        // for(let i = 0; i < this.enteties.length; i++) {
        //     this.enteties[i].update(this.app);
        // }
    }
}
class Player {
    app: PIXI.Application;
    body: PIXI.Graphics;
    x = 0
    y = 0
    color = 0xFF0000
    constructor(app: PIXI.Application) {
        this.body = new PIXI.Graphics();
        this.app = app
    }
    draw(){
        this.body.beginFill(this.color);
        this.body.drawRect(0, 0, 64, 64);
        this.body.endFill();
        this.body.x = this.x;
        this.body.y = this.y;

        this.app.stage.addChild(this.body);
    }
}







