# Programming Project Tracker CLI

[![Go Tests](https://github.com/otg996/ptrack-cli/actions/workflows/validation.yml/badge.svg)](https://github.com/otg996/ptrack-cli/actions/workflows/validation.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/otg996/ptrack-cli)](https://goreportcard.com/report/github.com/otg996/ptrack-cli)

> A command-line interface for finding and listing local software development projects.

This client is an official implementation of the [Programming Project Tracker Specification](https://github.com/otg996/ptrack-spec).

## Installation

```bash
go install github.com/otg996/ptrack-cli@latest
```

## Usage

```bash
# Scan the current directory
ptrack-cli

# Scan a specific directory
ptrack-cli /path/to/your/projects
```

## Contributing

This project follows the same contribution process as the main specification. Please see the [main contributing guide](https://github.com/otg996/programming-project-tracker-spec/blob/main/CONTRIBUTING.md) for details.

## License

This project is licensed under the [GPL-2.0](LICENSE).
