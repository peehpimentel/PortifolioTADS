import os
import subprocess
import csv

# Using exiftool path as literal string.
exiftool = r"C:\Windows\exiftool.exe"

# Function to extract metadata from files.
def extract_metadata(filepath):
    try:
        result = subprocess.run([exiftool, '-json', filepath], capture_output=True, text=True, check=True)
        metadata = result.stdout
        return metadata
    except subprocess.CalledProcessError as err:
        # This f before the string is a formatted string literal, it helps incorporate expressions and variables into strings.
        print(f"Error trying to execute exiftool: {err}")
        return None

# Main function.
def main():
    # Get user's folder.
    folder_path = input("Enter the folder path: ")

    # Verify if the folder's path is valid.
    if not os.path.isdir(folder_path):
        print("Wrong path. Please, try entering a valid path.")
        return
    
    # Define the supported extensions.
    supported_extensions = ['.docx', '.txt', '.pdf', '.jpg', '.jpeg', '.png', '.gif', '.wav', '.mp3', '.ogg', '.mp4', '.mov', '.avi', '.mkv']

    # Creating a CSV file to save metadata.
    output_file = os.path.join(folder_path,"metadados.csv")
    with open(output_file, "w", newline="", encoding='utf-8') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(['File', 'Metadata'])

        # Extracting and saving every file's metadata into the folder.
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

