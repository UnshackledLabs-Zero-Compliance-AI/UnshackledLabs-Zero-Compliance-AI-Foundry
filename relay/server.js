const express = require('express');
const http = require('http');
const WebSocket = require('ws');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());

const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

// Store all connected clients
const clients = new Set();

wss.on('connection', (ws) => {
    clients.add(ws);
    ws.on('close', () => clients.delete(ws));
});

// HTTP endpoint for inference servers to push results
app.post('/publish', (req, res) => {
    const msg = JSON.stringify(req.body);
    clients.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(msg);
        }
    });
    res.json({ status: 'published', count: clients.size });
});

// Health check
app.get('/health', (req, res) => res.json({ status: 'ok' }));

server.listen(8080, () => console.log('Relay running on 8080'));