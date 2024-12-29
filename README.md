# Go Reloaded

## Introduction

This project involves building a simple text completion, editing, and auto-correction tool in Go. The tool will read a text file, apply modifications based on specific patterns or flags, and write the corrected text to a new file. It challenges developers to leverage string manipulation, file handling, and error handling techniques while adhering to Go best practices.

## Problem Statement

The program receives two arguments:
1. The name of the input file containing the text to modify.
2. The name of the output file where the modified text will be saved.

The tool performs the following modifications:
- Converts words based on specific flags:
  - `(hex)` replaces the preceding word with its decimal equivalent (hexadecimal input).
  - `(bin)` replaces the preceding word with its decimal equivalent (binary input).
  - `(up)` converts the preceding word to uppercase.
  - `(low)` converts the preceding word to lowercase.
  - `(cap)` converts the preceding word to capitalized form.
- For `(low, <number>)`, `(up, <number>)`, or `(cap, <number>)`, it applies the transformation to the specified number of preceding words.
- Formats punctuation to be closer to the preceding word and spaced apart from the following word.
- Adjusts `'` marks to enclose words or groups of words correctly.
- Corrects "a" to "an" before words starting with vowels or 'h'.

## Features

- Handles complex text transformations using user-defined flags.
- Formats punctuation correctly, even for edge cases like ellipses (`...`) and compound punctuation (`!?`).
- Ensures proper usage of articles ("a" vs. "an") based on phonetics.

## Error Handling

The program ensures robust error handling:
- **File Reading Errors:** If the input file does not exist or cannot be read, the program gracefully exits with an appropriate error message.
- **Invalid Flags:** Unrecognized flags or invalid numeric values are ignored, and the text is left unchanged.
- **Output File Writing Errors:** If the output file cannot be written, the program displays an error message without altering the existing file.

## Solution Outline

1. **File Handling:**
   - Use `os.ReadFile` to read the input file and check for errors.
   - Use `os.WriteFile` to save the processed text to the output file.

2. **String Manipulation:**
   - Process the text by splitting it into words using `strings.Fields`.
   - Implement transformations for flags like `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` based on their definitions.
   - Ensure proper handling of punctuation and article correction.

3. **Functions:**
   - **Hex and Bin Conversion:** Convert hexadecimal and binary strings to decimal using `strconv.Atoi` or similar methods.
   - **Text Transformations:** Implement functions for uppercase, lowercase, and capitalization.
   - **Punctuation Formatting:** Use conditions to format punctuation correctly.
   - **Article Correction:** Adjust "a" to "an" based on the subsequent word.

4. **Output:**
   - Combine the modified words and save them to the specified output file.

---
This project is an excellent opportunity to deepen your understanding of Go's file system API, string manipulation, and error handling. Happy coding!
