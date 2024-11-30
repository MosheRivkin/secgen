package help

const helpText = `Secret Generator (secgen)

Usage: secgen [length] [flags]

Arguments:
  length      Number of characters in the secret (default: 32)
  
Flags:
  n          Include numbers (0-9)
  l          Include lowercase letters (a-z) 
  u          Include uppercase letters (A-Z)
  s          Include special characters (!@#$%^&*()_+-=[]{}|;:,.<>?)
  !          Include everything except following flags
  -h         Show this help message

Examples:
  secgen 10              Generate 10 character secret (default flags: !s)
  secgen 16 nlu         16 chars with numbers, lower and uppercase
  secgen 8 nlus         8 chars with all character types
  secgen 12 !n          12 chars excluding numbers
  secgen 20 !           20 chars with all character types

Default: secgen 32 !s (32 characters excluding special chars)`
