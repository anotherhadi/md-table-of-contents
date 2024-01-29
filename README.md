# Markdown Table Of Contents

Markdown Table Of Contents is a straightforward command-line tool designed to generate a table of contents for your Markdown files based on titles.

## Features

- **Automatic Generation**: Quickly generate a table of contents for your Markdown files.
- **Title Detection**: Detects headings in your Markdown file and generates links to them in the table of contents.

## How to Use

1. **Installation**:
  ```bash
  go get github.com/anotherhadi/md-table-of-contents@latest
  ```


2. **Usage**: Run the executable with your Markdown file as an argument.
   ```bash
   md-table-of-contents <markdownfile.md>
   ```
3. **Output**: The tool will generate a table of contents based on the headings in your Markdown file.

## Example

Suppose you have a Markdown file named `example.md` with the following headings:

```md
# Introduction
## Getting Started
### Installation
### Usage
## Features
# Conclusion
```

Running the command:

```bash
md-table-of-contents example.md
```

Would generate the following table of contents:

- [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Features](#features)
- [Conclusion](#conclusion)

## License

This project is licensed under the [MIT License](LICENSE).
