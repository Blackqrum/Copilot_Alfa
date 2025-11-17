FROM golang:1.21-bullseye  

# Установка зависимостей для Ollama
RUN apt-get update && apt-get install -y \
    curl \
    bash \
    && rm -rf /var/lib/apt/lists/*

# Установка Ollama
RUN curl -fsSL https://ollama.ai/install.sh | sh

WORKDIR /app

COPY go.mod ./
RUN go mod tidy && go mod download

COPY . .

RUN go build -o main .

# Копируем скрипт запуска
COPY start.sh /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]