# Go Text Transformer

A simple command-line tool written in Go that reads an input text file, processes each line, and writes the modified result to an output file.

---

## 📌 Features

* Reads a text file line by line
* Applies transformations based on content:

  * Converts fully uppercase lines to lowercase
  * Converts fully lowercase lines to uppercase
  * Replaces `TODO:` (or `todo`) with `ACTION:`
* Writes the processed lines to a new output file

---

## ⚙️ Requirements

* Go installed (version 1.18 or later recommended)

---

## 🚀 Usage

Run the program using:

```bash
go run main.go input.txt output.txt
```

### Arguments

* `input.txt` → The source file to read from
* `output.txt` → The destination file where results are written

---

## 🧠 How It Works

The program processes the file line by line:

1. Reads each line using a scanner
2. Applies conditional transformations:

   * If the entire line is uppercase → convert to lowercase
   * If the entire line is lowercase → convert to uppercase
   * If the line contains `TODO:` or `todo` → replace with `ACTION:`
3. Writes the modified line to the output file

---

## 📝 Example

### Input (`input.txt`)

```
I AM OF GOD ALMIGHTY
jesus is his name
TODO: we are ready
```

### Output (`output.txt`)

```
i am of god almighty
JESUS IS HIS NAME
ACTION: we are ready
```

---

## ⚠️ Notes

* String comparisons are case-sensitive
* Only lines that are fully uppercase or fully lowercase are transformed
* Mixed-case lines remain unchanged unless they contain `TODO:` or `todo`
* The program does not modify the original file; it creates/overwrites the output file

---

## 🔧 Possible Improvements

* Support mixed-case transformations (e.g., title case)
* Add flags instead of positional arguments
* Handle punctuation more robustly
* Add unit tests
* Improve error handling and logging

---

## 📄 License

This project is open-source and free to use for learning purposes.
