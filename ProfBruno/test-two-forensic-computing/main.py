import os
import subprocess
import csv

# Tratando o caminho do exiftool como string literal
exiftool = r"C:\Windows\exiftool.exe"

# Função para extrair metadados de arquivos DOCX
def extract_metadata(filepath):
    try:
        result = subprocess.run([exiftool, '-json', filepath], capture_output=True, text=True, check=True)
        metadata = result.stdout
        return metadata
    except subprocess.CalledProcessError as err:
        # Esse f antes da string é uma forma de formatted string literal que serve para incorporar expressões e variáveis dentro de strings.
        print(f"Error trying to execute exiftool: {err}")
        return None

# Função principal
def main():
    # Obter o caminho da pasta do usuário
    folder_path = input("Enter the folder path: ")

    # Verificar se o caminho da pasta é valido
    if not os.path.isdir(folder_path):
        print("Wrong path. Please, try entering a valid path.")
        return
    
    # Definir os tipos de arquivos suportados
    supported_extensions = ['.docx', '.txt', '.pdf', '.jpg', '.jpeg', '.png', '.gif', '.wav', '.mp3', '.ogg', '.mp4', '.mov', '.avi', '.mkv']

    # Criar arquivo CSV para salvar os metadados
    output_file = os.path.join(folder_path,"metadados.csv")
    with open(output_file, "w", newline="", encoding='utf-8') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(['File', 'Metadata'])

        # Extrair e salvar metadados de cada arquivo na pasta
        for root, dirs, files in os.walk(folder_path): 
            for filename in files: 
                filepath = os.path.join(root, filename)
                if any(filename.lower().endswith(ext) for ext in supported_extensions):
                    print(f"Processing file: {filepath}")
                    metadata = extract_metadata(filepath)
                    if metadata:
                        writer.writerow([filepath, metadata])
                        print(f"Metadata extracted e saved into: {filepath}")
                    else:
                        print(f"Failed to extracted metadata into: {filepath}")

if __name__ == "__main__":
    main()

