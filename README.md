### `README.md`

# goco

`goco` is a CLI tool written in Go that allows you to set the `GOOS` and `GOARCH` environment variables interactively. These variables configure the target operating system and architecture for building Go programs, making cross-compilation easy and intuitive.

## Features

- **Interactive Selection**: Choose the target `GOOS` and `GOARCH` from a prioritized list.
- **Cross-Platform**: Automatically detects the host operating system (Windows or Linux/macOS) and applies the appropriate commands.
- **Immediate Setup**: Sets the environment variables in your terminal for immediate use.
- **Manual Instructions**: Provides equivalent commands for manual setup.

## Prerequisites

- Go 1.20 or higher installed on your system.
- Ensure your terminal supports the commands:
  - `export` for Linux/macOS
  - `$env:` for Windows PowerShell

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/goco.git
   cd goco
   ```

2. Build the executable:

   ```bash
   go build -o goco
   ```

3. (Optional) Move the executable to a directory in your `PATH`:

   ```bash
   mv goco /usr/local/bin/
   ```

   On Windows, move `goco.exe` to a folder in your `PATH`.

## Usage

1. Run the program:

   ```bash
   ./goco
   ```

   On Windows:

   ```powershell
   goco.exe
   ```

2. Follow the prompts to select the desired `GOOS` and `GOARCH`.

3. The tool will:
   - Automatically set the selected environment variables for the current terminal session.
   - Output equivalent manual commands for reference.

### Example

```text
Select a target operating system (GOOS):
1) linux
2) windows
3) darwin
4) android
5) freebsd
...

Enter the number of your choice (1-16): 1

Select a target architecture (GOARCH):
1) amd64
2) 386
3) arm
...

Enter the number of your choice (1-15): 1

Environment variables set in the current shell:
  GOOS   = linux
  GOARCH = amd64

To set these manually, use the following commands:
  export GOOS=linux
  export GOARCH=amd64

Bye.
```

## Limitations

- The tool sets the environment variables only for the current terminal session. For persistent changes, add the displayed commands to your shell configuration file (e.g., `.bashrc`, `.zshrc`, or PowerShell profile).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue if you have ideas for improvements or bug fixes.

---

Happy coding with Go! 🐹
