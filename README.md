# CaesarCipher
Golang CLI tool for encrytping and decrypting strings using the [Caesar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher) with the shifing value increasing by 1 per character in the string(including spaces)


## Usage

`./caesarCipher "hello there"`

`jgnnq vjgtg` 

-d flag to decrypt the passed string
example: `./caesarCipher -d "wtaad iwtgt"`

-s flag for shift value, default is 1.
example: `./caesarCipher -s=14 "hello there"`
