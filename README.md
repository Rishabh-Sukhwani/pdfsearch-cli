<h1 align="center">pdfsearch-cli</h1>

This is a command-line tool that can search for a specific string in PDF files in a given directory and its subdirectories.

---

## Installation

1. Clone this repository or download the ZIP file and extract it.
2. Make sure you have Go installed on your system. You can download it from https://golang.org/dl/.
3. Open a terminal or command prompt and navigate to the root directory of this project.
4. Run go install to build and install the tool.

---

## Usage

```bash
./pdfsearch [flags]
```

### Commangs

* (default): Searches for a specific string in PDF files in a given directory and its subdirectories.

### Flags

* --dir: The directory to search in. Defaults to the current working directory.
* --search: The string to search for.

### Examples 

* Search for the string 'example' in the directory 'mydir':

```bash
./pdfsearch --search="example" --dir="mydir"
```

* the help command

```bash
$ ./pdfsearch-cli --help
```

this returns:

```bash
pdfsearch is a CLI tool that searches for a given string in all PDF files in a directory
and its subdirectories.

Usage:
  pdfsearch [flags]

Flags:
  -d, --dir string      the directory to search in (default ".")
  -h, --help            help for pdfsearch
  -s, --search string   the string to search for
```

---

## Contributing

If you find a bug or have a feature request, please open an issue on the GitHub repository. Pull requests are welcome!