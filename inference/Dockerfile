FROM python:3.10-slim
WORKDIR /app
RUN pip install fastapi uvicorn requests pydantic
COPY server.py .
EXPOSE 8000
CMD ["uvicorn", "server:app", "--host", "0.0.0.0", "--port", "8000"]