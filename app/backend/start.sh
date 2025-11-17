#!/bin/bash

echo "Starting Ollama server..."
OLLAMA_HOST=0.0.0.0:11434 /usr/local/bin/ollama serve &

echo "Waiting for Ollama to start..."
sleep 15

echo "Pulling AI model..."
/usr/local/bin/ollama pull qwen2.5:7b

echo "Starting backend application..."
./main