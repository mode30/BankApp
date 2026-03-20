# BankApp

A simple command-line banking application written in Go. Supports basic account operations — withdrawal and deposit — with structured error handling and a clean package layout.

---

## Project Structure

```
bankApp/
├── go.mod
├── bankApp.go        # entry point, Bank struct definition
├── userService.go    # withdrawal and deposit logic
├── shared/
│   └── shared.go     # shared state (account balance)
└── userInput/
    └── userInput.go  # reads and trims user input from stdin
```

---

## Features

- Withdraw from an account balance with overdraft protection
- Deposit funds with negative value validation
- Modular package structure separating I/O, business logic, and shared state
- Errors propagated cleanly up the call stack to `main()`

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21 or higher

### Run

```bash
# From the project root
go run .
```

### Build

```bash
go build .
```

---

## Usage

The app runs interactively in the terminal. On launch it will prompt you to:

1. Enter an amount to withdraw
2. Enter an amount to deposit

The current balance starts at `10,000.00` and updates after each operation.

```
enter amount to withdraw: 2000
enter amount to deposit: 500
```

---

## Error Handling

Errors are returned up the call stack and handled only in `main()`. Service functions never call `log.Fatal` directly — they return descriptive errors using `fmt.Errorf` with `%w` wrapping for context, and `errors.New` for static validation messages.

---

## Limitations

- Balance is held in memory — no persistence between runs
- Single account only
- No authentication

