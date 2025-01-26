# SHA-256 Hash Generator

A simple Go program to generate SHA-256 hashes for files.

## Usage

### Command-line arguments

* `file_path`: The path to the file for which you want to generate the SHA-256 hash.
* `--save` or `-s`: Optional flag to save the hash to a file with the same name as the input file but with a `.sha256` extension.

### Example

```bash
# On Windows
sha256.exe example.txt

# On Linux
./sha256 example.txt
```
This will generate the SHA-256 hash for the `example.txt` file and print it out to the console.

```bash
# On Windows
sha256.exe example.txt --save # or -s

# On Linux
./sha256 example.txt --save # or -s
```
##### _Replace sha256 with your executable name._

This will generate the SHA-256 hash for the `example.txt` file and save it to a file named `example.txt.sha256`.

## Building and Running

To build this program, you'll need to have Go installed on your system. You can then build the program using the following command:
```bash
go build main.go
```

## Current Issues
- Eats a lot of memory if processing large files _(and by a lot, I mean, a LOT)_.

## License

This program is licensed under the [MIT License](https://opensource.org/licenses/MIT).
