#!/bin/bash

LOG_FILE="log_react.txt"
PORT=3000
START_CMD="bun run start"

## ? Inicialização do servidor ao fundo,
## ? e colocando as informações no log
cd frontend
$START_CMD > "$LOG_FILE" 2>&1 & 

REACT_PID=$!
echo "Servidor react incializado com o PID: $REACT_PID"

for i in {1..30}; do 
    if (echo > /dev/tcp/localhost/$PORT) >/dev/null 2>&1; then
        exit 0
    fi 
    sleep 1
done 

echo "Falha ao iniciar o servidor após 30 segundos."
kill $REACT_PID
exit 1 