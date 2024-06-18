# File Encryption and Decryption Web Application

This web application provides a simple interface for encrypting and decrypting text using AES encryption. It consists of a Go backend server for handling encryption and decryption operations, and a frontend HTML/CSS/JavaScript interface for user interaction.

## Features

- **Encryption**: Encrypts input text using AES encryption and displays the encrypted result as a hex-encoded string.
- **Decryption**: Decrypts hex-encoded encrypted text back to its original plaintext form.
- **User Interface**: Clean and intuitive interface designed using HTML, CSS, and JavaScript for seamless user interaction.

## Setup and Usage

### Prerequisites

- Go programming language installed ([Download and Install Go](https://golang.org/doc/install))
- Modern web browser (Chrome, Firefox, Safari, etc.)

### Installation and Running the Application

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/Encryption-Decryption-GO.git
   cd Encryption-Decryption-GO
   ```

2. **Start the Go Server**:
   ```bash
   go run main.go
   ```
   This will start the Go server on `http://localhost:8080`.

3. **Access the Web Application**:
   - Open your web browser and go to `http://localhost:8080`.
   - You should see the main page of the application with a textarea for input, radio buttons for selecting encryption or decryption, and a submit button.

### Usage

- **Encrypting Data**:
  1. Enter text into the textarea.
  2. Select the "Encrypt" radio button.
  3. Click the "Submit" button.
  4. The encrypted data will be displayed in the output area.

- **Decrypting Data**:
  1. Paste the hex-encoded encrypted data into the textarea.
  2. Select the "Decrypt" radio button.
  3. Click the "Submit" button.
  4. The decrypted plaintext will be displayed in the output area.

### Development and Customization

- **Adding Styles**:
  - Modify `style.css` to customize the appearance of the web application.

- **Enhancing Functionality**:
  - Expand functionality by modifying `main.go` to handle additional features or improve error handling and security measures.

### Security Considerations

- This application demonstrates basic AES encryption/decryption techniques and is suitable for learning purposes.
- For production use, implement HTTPS to secure data transmission between the client and server.
- Ensure proper handling and storage of encryption keys to maintain data security.

