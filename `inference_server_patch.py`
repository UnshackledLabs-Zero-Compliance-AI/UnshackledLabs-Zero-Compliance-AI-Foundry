import asyncio, json, os
import websockets

TOKEN_SOCKET = os.getenv("TOKEN_SOCKET", "ws://localhost:8765")
COUNTER_SECRET = os.getenv("COUNTER_SECRET", "unshackled")  # simple auth

async def _report(count: int):
    try:
        async with websockets.connect(TOKEN_SOCKET) as ws:
            await ws.send(json.dumps({"secret": COUNTER_SECRET, "delta": count}))
    except:
        pass  # fail silently – we never block generation

def report_tokens_generated(n: int):
    asyncio.create_task(_report(n))
