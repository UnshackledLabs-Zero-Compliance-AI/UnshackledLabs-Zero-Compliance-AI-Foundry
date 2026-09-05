import os
import time
import json
import requests
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from pydantic import BaseModel
import uvicorn

app = FastAPI()
RELAY_URL = os.getenv("RELAY_URL", "http://relay:8080")

class Prompt(BaseModel):
    text: str

@app.post("/predict")
async def predict(prompt: Prompt):
    # Simulate inference (replace with actual model loading)
    time.sleep(1.5)  # fake compute
    result = f"Echo: {prompt.text[::-1]}"  # reverse string as mock
    
    # Send result to the relay
    try:
        requests.post(f"{RELAY_URL}/publish", json={"message": result})
    except:
        pass
    return {"result": result}

@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await websocket.accept()
    try:
        while True:
            data = await websocket.receive_text()
            # Echo back with timestamp
            await websocket.send_text(json.dumps({"echo": data, "ts": time.time()}))
    except WebSocketDisconnect:
        pass

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)