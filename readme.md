# Rand
A Command-Line tool for generating a cryptographically secure random bytes.

### Install
`
- git clone https://github.com/mukailasam/rand
- cd rand
- go install

`

### Run
`rand`

### Usage

**Write to Console**
- rand -n 16
- rand -n 32
- rand -n 64
- rand -n 128
- rand -n 256

**Write to a File**
- rand -n 16 -o rand.txt
- rand -n 1024 -o rand.txt
- rand -n 2048 -o rand.txt
- rand -n 4096 -o rand.txt

### Note
- By default 32 bytes will be the number of output bytes if the number of bytes is not specify
e.g
rand

- Generated random bytes ouput is in hexadecimal form

- Number of bytes argument can be of any size in this range[16-100001]