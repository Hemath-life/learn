import os
import shutil

def organize_files(path):
    # List all files in the specified directory
    files = os.listdir(path)

    # Remove 'main.py' from the list, as we don't want to move it
    files.remove('main.py')

    # Print the list of files to be organized
    print("Files to be organized:", files)

    # Loop through each file in the list
    for file in files:
        # Split the filename and extension
        filename, extension = os.path.splitext(file)
        extension = extension[1:]  # Remove the leading dot from the extension

        # Create the destination folder path
        dest_folder = os.path.join(path, extension)

        # Check if the destination folder exists, if not, create it
        if not os.path.exists(dest_folder):
            os.makedirs(dest_folder)

        # Construct the source and destination file paths
        src_file = os.path.join(path, file)
        dest_file = os.path.join(dest_folder, file)

        # Move the file to the destination folder
        shutil.move(src_file, dest_file)

if __name__ == "__main__":
    # Change this path to the desired directory that you want to organize
    current_directory = './'
    organize_files(current_directory)
