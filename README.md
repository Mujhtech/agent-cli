# Agent CLI

An open-source alternative to OpenAI Codex & Claude Code

## Prerequisites

- Go 1.21 or later
- Git

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/agent-cli.git
cd agent-cli
```

2. Install dependencies:

```bash
go mod download
```

## Building

To build the project, run:

```bash
make build
```

This will create an executable named `agent-cli` in the root directory.

## Running

After building, you can run the CLI tool:

```bash
./agent-cli
```

## Development

The project structure is organized as follows:

- `cmd/`: Contains the main application entry points
- `internal/`: Private application code
- `pkg/`: Public library code that can be used by external applications

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
