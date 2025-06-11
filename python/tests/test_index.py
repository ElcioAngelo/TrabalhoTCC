# tests/test_index.py
import os
import tempfile
from PIL import Image

## ? Criação da pasta de imagens
IMAGE_FOLDER: str = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'images')

## ? Extensões de arquivos permitidas
ALLOWED_EXTENSIONS: set[str] = {'png','jpg','jpeg','gif','bmp','webp'}

## ? Verifica se a pasta das imagens existe 
os.makedirs(IMAGE_FOLDER, exist_ok=True)

## ? Verifica se o arquivo enviado é um arquivo de imagem
def is_allowed_file(filename: str) -> bool:
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

## ? Verifica novamente se o arquivo salvo é realmente uma imagem
def is_valid_image(file_path: str) -> bool:
    try:
        with Image.open(file_path) as img:
            img.verify()
        return True
    except Exception:
        return False

def test_is_allowed_file():
    assert is_allowed_file('foto.jpg')
    assert is_allowed_file('imagem.webp')
    assert not is_allowed_file('documento.pdf')

def test_is_valid_image():
    with tempfile.NamedTemporaryFile(suffix='.png', delete=False) as tmp:
        img = Image.new('RGB', (10, 10), color='red')
        img.save(tmp.name)
        tmp_path = tmp.name

    assert is_valid_image(tmp_path) is True

    with tempfile.NamedTemporaryFile(suffix='.txt', delete=False) as tmp:
        tmp.write(b'this is not an image')
        tmp_path = tmp.name

    assert is_valid_image(tmp_path) is False
