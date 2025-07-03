cd python
source virtualenv/bin/activate 
pytest

if [ $? -eq 0 ]; then
    echo "Os testes passaram. Iniciando o servidor."
    python index.py
else 
    echo "Os testes falharam."
    exit 1
fi
