FROM python:3.11-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    python3-dev \
    libpq-dev \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt .
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install -r requirements.txt

COPY python_server/ ./python_server/
COPY generated/python/ ./generated/python/

ENV PYTHONPATH=/app

EXPOSE 5060

CMD ["python", "python_server/server/grpc_server.py"]