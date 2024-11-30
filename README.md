# secgen - Secret Generator

A command-line tool to generate secure random secrets with configurable character sets.

## Installation
just download and extract the binary from the [releases](https://github.com/MosheRivkin/secgen/releases) page and add it to your path.

## Usage

```sh
secgen [length] [flags]
```

### Arguments

- `length`: Number of characters in the secret (default: 32)

### Flags

- `n` - Include numbers (0-9)
- `l` - Include lowercase letters (a-z)
- `u` - Include uppercase letters (A-Z) 
- `s` - Include special characters (!@#$%^&*()_+-=[]{}|;:,.<>?)
- `!` - Include everything except following flags
- `-h` - Show help message

### Examples

```sh
# Generate 10 character secret without special characters (default flags: !s)
secgen 10
output: yY12yK8w11

# 16 chars with numbers, lower and uppercase (equivalent to secgen 16 !s)
secgen 16 nlu
output: SaC0kpmZe93d3zWF

# 8 chars with all character types
secgen 8 nlus
output: =D-IB.+=

# 20 chars without nothing i.e. with all character types (equivalent to secgen 20 nlus)
secgen 20 !
output: 2@l6).)dLu,5az[Pzt_W

# 12 chars excluding numbers 
secgen 12 !n
output: OvJ>}OPsiDU

# default 32 chars with all character types except special characters
secgen
output: jirb8rfXOjwh9RUekmeAPc91lDDe9XDV

```
