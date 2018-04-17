const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', function connection(ws) {

  console.log(new Date() + ' - Connected');

  ws.on('message', function incoming(message) {

    console.log(new Date() + ' - Message - ' + message);

    var input = JSON.parse(message);

    var command = {
      "movement": {
        "yaw": input.yaw,
        "roll": input.roll,
        "pitch": input.pitch,
        "gaz": input.gaz
      }
    };

    ws.send(JSON.stringify(command));

  });
  
});