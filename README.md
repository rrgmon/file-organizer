# Sortzilla: Efficient File Sorting by Type

Sortzilla is a powerful command-line tool designed to bring order to your files. With Sortzilla, you can effortlessly organize your files based on their types, saving you time and keeping your digital life well-structured.

## Features

- **Effortless File Sorting**: Sortzilla simplifies the task of organizing files by automatically categorizing them into dedicated folders based on their file types.
- **Supported File Types**: Sortzilla supports a wide range of file types, including popular formats such as `.jpg`, `.png`, `.txt`, `.pdf`, `.docx`, and `.xlsx`.
- **Saves Time and Reduces Clutter**: By automating the file sorting process, Sortzilla helps you regain control over your files, ensuring a clutter-free and organized file system.

## Usage

1. Download the latest release of Sortzilla from the [releases](https://github.com/rrgmon/sortzilla/releases) page.

2. Run Sortzilla from the command line in one of the following ways:

   - To organize files in the current working directory (PWD), simply execute:

     ```
     ./sortzilla
     ```

   - To organize files in a specific target folder, use the `--target` option followed by the target folder path. For example:

     ```
     ./sortzilla --target /path/to/target/folder
     ```

   If the `--target` option is not provided, Sortzilla will default to organizing files in the current working directory.

3. Sit back and watch as Sortzilla efficiently organizes your files into designated folders based on their file types.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributions and Feedback

Contributions and suggestions are welcome! If you have any ideas or encounter any issues, please feel free to submit a pull request or open an issue on the [GitHub repository](https://github.com/rrgmon/sortzilla).

## Building from Source

If you prefer to build Sortzilla from the source code, follow these steps:

1. Make sure you have Go installed on your system.
2. Clone the repository or download the source code.
3. Navigate to the project's root directory.
4. Build the executable by running the following command:

   ```
   go build -o sortzilla
   ```

5. Run Sortzilla by executing the `sortzilla` binary, either without any arguments to organize files in the current working directory, or with the `--target` option followed by the target folder path.

## Note

Please remember to exercise caution when using Sortzilla and ensure that you have proper backups of your files. Although Sortzilla is designed to organize files safely, it's always a good practice to have backups in place.

Enjoy the benefits of a well-structured file system with Sortzilla!

---
