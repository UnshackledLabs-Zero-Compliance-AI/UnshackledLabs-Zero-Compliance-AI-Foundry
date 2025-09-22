const WebSocket = require('ws');
const wss = new WebSocket.Server({ port: 8765 });
let clients = new Set();

wss.on('connection', ws => {
  clients.add(ws);
  ws.on('close', () => clients.delete(ws));
});

// listen for deltas from inference box
wss.on('connection', ws => {
  ws.on('message', data => {
    try {
      const msg = JSON.parse(data);
      if (msg.secret !== process.env.COUNTER_SECRET) return;
      const pkt = JSON.stringify({ op: 'add', val: msg.delta });
      clients.forEach(c => c.readyState === 1 && c.send(pkt));
    } catch {}
  });
});
