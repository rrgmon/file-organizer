import os
import shutil

def organize_files(target_folder):
    extensions = {
        ".jpg": "Images",
        ".png": "Images",
        ".txt": "TextFiles",
        ".pdf": "PDFs",
        ".docx": "WordDocuments",
        ".xlsx": "ExcelFiles"
    }

    for filename in os.listdir(target_folder):
        file_path = os.path.join(target_folder, filename)
        if os.path.isfile(file_path):
            _, extension = os.path.splitext(filename)
            extension = extension.lower()

            folder_name = extensions.get(extension, "OtherFiles")
            destination_folder = os.path.join(target_folder, folder_name)

            os.makedirs(destination_folder, exist_ok = True)

            shutil.move(file_path, os.path.join(destination_folder, filename))

    print("Files organized successfully")


def main():
    target_folder = input("Enter the target folder path: ")

    if not os.path.isdir(target_folder):
        print("Invalid  folder path.")
        return

    organize_files(target_folder)

if __name__ == "__main__":
    main()
