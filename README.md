# File Organizer CLI Tool (Go)

A simple command-line tool written in Go that organizes files in a folder by their extensions. Perfect for cleaning up messy directories automatically.

---

## Features

- Organizes files into folders based on file extension.
- Creates folders automatically if they don't exist.
- Supports `--dry-run` mode to preview changes without moving files.
- Simple CLI flags for ease of use.

---

## Installation

1. Clone the repository:
```bash
git clone https://github.com/saitama-op/file-organizer.git
cd file-organizer
```

2. Initialize Go module:
```bash
go mod init file-organizer
```

3. Run the tool:
```bash
go run main.go --path <directory_path> [--dry-run]
```

---

## Usage

### Organize current directory:
```bash
go run main.go --path .
```

### Dry-run (preview changes without moving files):
```bash
go run main.go --path ./Downloads --dry-run
```

### CLI Flags

| Flag       | Description                                      | Default |
|------------|--------------------------------------------------|---------|
| `--path`   | Path of the directory to organize               | `.`     |
| `--dry-run`| Show what would be done without moving files   | false   |

---

## Example Output

```text
Moved ./Downloads/file1.pdf → ./Downloads/pdf/file1.pdf
Moved ./Downloads/image1.jpg → ./Downloads/jpg/image1.jpg
[DRY-RUN] Move ./Downloads/file2.docx → ./Downloads/docx/file2.docx
```

---

## Contributing

Feel free to fork this repository and submit pull requests for improvements, such as:
- Adding concurrency for faster processing of large directories.
- Custom folder mappings for specific file types.
- Logging and reporting moved files.

---

## License

This project is licensed under the MIT License.

