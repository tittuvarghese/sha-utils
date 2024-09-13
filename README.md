# Hashing Utility

This Go utility program hashes each line from an input file using a specified hashing algorithm and writes the resulting hashes to an output file. Supported hashing algorithms include SHA-1, SHA-256, and SHA-512. The program also allows you to specify the bit size of the hash.

## Features

- Hashes input lines using SHA-1, SHA-256, or SHA-512 algorithms.
- Supports customizable bit sizes (128, 256, 512).
- Reads from an input file and writes hashed results to an output file.

## Example Usage

```bash
# To create SHA512 with bit size 512
./shautils -input data.txt -output hashed_output.txt -algorithm SHA512 -bits 512

# To create SHA128 with bit size 128
./shautils -input data.txt -output hashed_output.txt -algorithm SHA128 -bits 128

```

## Building from Source

1. **Clone the repository:**

   ```bash
   git clone https://github.com/tittuvarghese/sha-utils.git
   ```

2. **Navigate to the directory**

    ```bash
    cd your-repository
    go build -ldflags="-s -w" -o shautils
    ```
3. **Usage**
    ```bash
    ./shautils -input data.txt -output hashed_output.txt -algorithm SHA512 -bits 512
    ```