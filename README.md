 A command-line tool written in Go that generates secure, random passwords based on user-specified criteria such as length and character types.

 ## Features
 - Generates cryptographically secure passwords using Go's `crypto/rand`.
 - Customizable options for including lowercase, uppercase, numbers, and special characters.
 - Simple and intuitive CLI interface.

 ## Installation
 1. Ensure [Go](https://golang.org/dl/) is installed.
 2. Clone the repository:
    ```bash
    git clone https://github.com/your-username/advanced-password-generator.git
    cd advanced-password-generator
    ```
 3. Run the program:
    ```bash
    go run main.go
    ```

 ## Usage
 Generate a password with default settings (12 characters, all character types):
 ```bash
 go run main.go
 ```

 Customize the password:
 ```bash
 go run main.go -length 16 -special=false
 ```

 Available flags:
 - `-length int`: Password length (default 12)
 - `-lower bool`: Include lowercase letters (default true)
 - `-upper bool`: Include uppercase letters (default true)
 - `-numbers bool`: Include numbers (default true)
 - `-special bool`: Include special characters (default true)

 ## Example Output
 ```bash
 $ go run main.go -length 8 -special=false
 k7nP9mW2
 ```
