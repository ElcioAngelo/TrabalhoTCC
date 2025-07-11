from flask import Flask, render_template, request, send_from_directory, abort, jsonify, Response
from werkzeug.utils import secure_filename
from werkzeug.datastructures import FileStorage
from typing import Optional
from PIL import Image
import requests
import os
import uuid

app = Flask(__name__)

## ? Localiza a pasta de imagens
IMAGE_FOLDER: str = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'images')

## ? Extensões de arquivos permitidas
ALLOWED_EXTENSIONS: set[str] = {'png','jpg','jpeg','gif','bmp','webp'}

## ? Verifica se a pasta das imagens existe 
os.makedirs(IMAGE_FOLDER, exist_ok=True)

## ? Verifica se o arquivo enviado é de uma extensão permitida
def is_allowed_file(filename: str) -> bool:
    return '.' in filename and filename.split('.', 1)[1].lower() in ALLOWED_EXTENSIONS

def is_valid_image(file_path: str) -> bool:
    try:
        with Image.open(file_path) as img:
            img.verify()
        return True
    except Exception:
        return False

## ? Limita o tamanho máximo de arquivos enviados para 10MB
app.config['MAX_CONTENT_LENGTH'] = 10 * 1024 * 1024  
    

## ** PÁGINA INICIAL
@app.route('/',methods=['GET'])
def index():
    return render_template('index.html')

## ** ROTA PARA VISUALIZAÇÃO DAS IMAGENS 
@app.route('/images',methods=['GET'])
def return_images():
    files = []
    for filename in os.listdir(IMAGE_FOLDER):
        if filename.lower().endswith(['.png', '.jpg', '.jpeg', '.gif', '.webp']):
            files.append(filename)
    return jsonify(files)

## ** ROTA PARA RETORNAR IMAGENS 
@app.route('/images/<filename>')
def serve_image(filename: str) -> Response:
    try:
        return send_from_directory(IMAGE_FOLDER, filename)
    except FileNotFoundError:
        abort(404)

## * Rota para a galeria de imagens.
@app.route('/imageGallery')
def image_galleri():
    return render_template('ImageGallery.html')

## ** ROTA PARA FAZER O UPLOAD DE IMAGENS
@app.route('/upload', methods=['POST'])
def upload_image() -> Response:

    ## ? Recebe o arquivo e o nome personalizado
    file: Optional[FileStorage] = request.files.get('image')
    custom_name = request.form.get('imageName')

    ## ? Verifica se o arquivo foi enviado e se o nome é válido.
    if not file or not custom_name:
        return jsonify({'error': 'Nenhum arquivo foi enviado'}), 204
    
    ## ? Caso o arquivo enviado não seja do tipo aceitado, retorne erro.
    if not is_allowed_file(file.filename):
        return jsonify({'error': 'Tipo de arquivo invalido'}), 415

    ## ? Gera um novo nome de arquivo único, mantendo a extensão original
    ext: str = secure_filename(file.filename).split('.',1)[1].lower()
    safe_name = secure_filename(custom_name) + '.' + ext
    
    ## ? Armazena a imagem renomeada
    file_path: str = os.path.join(IMAGE_FOLDER, safe_name)

    ## ? Caso uma imagem com nome similar existir, retorne erro
    if os.path.exists(file_path):
        return jsonify({'erro': 'Uma imagem come esse nome já existe.'}), 409
    else:
        ## ? Salva a imagem
        file.save(file_path)

    ## ? Verificação se o arquivo é uma imagem válida
    if not is_valid_image(file_path):
        os.remove(file_path)
        return jsonify({'error': 'O arquivo não é uma imagem valida'}), 415
    
    return jsonify({'message': 'Imagem enviada com sucesso'}), 201


## * Rota necessária para o script perl funcionar.
@app.route('/check', methods=['GET'])
def status_checking():
    return jsonify({'status': 'ok'}), 200

if __name__ == '__main__':
    app.run(debug=True)
