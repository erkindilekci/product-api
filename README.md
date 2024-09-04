# Product API
Product API is a web application designed to manage and serve product-related data. It leverages several modern technologies to provide a robust and efficient backend service.

## Technologies

- `Go`
- `Echo`
- `PostgreSQL`
- `Testify`

## Setting Up Product API Locally

To run the Product API application on your local machine, follow these steps:

1. **Clone the Repository:**
   Begin by cloning the Product API repository to your local environment using Git:
   ```bash
   git clone https://github.com/erkindilekci/product-api.git
   ```

2. **Install Dependencies:**
   Navigate to the project directory and install the required dependencies:
   ```bash
   cd product-api
   go mod tidy
   ```

3. **Run the Application:**
   Start the application using the Go command:
   ```bash
   go run ./cmd/productapi
   ```

4. **Access the Application:**
   Open your web browser and navigate to `http://localhost:8080` to access the Product API.

**Note:** Ensure you have Go installed on your system before proceeding with these steps.

## Major Dependencies

- **Echo:** A high performance, extensible, minimalist web framework for Go.
- **pgx:** A PostgreSQL driver and toolkit for Go.
- **Testify:** A toolkit with common assertions and mocks that plays nicely with the standard library.

## Additional Dependencies

The project also includes several indirect dependencies that support various functionalities, such as:

- **go-spew:** Implements a deep pretty printer for Go data structures.
- **pgconn, pgio, pgproto3, pgtype:** Various packages from the pgx family for PostgreSQL connection handling and protocol management.
- **go-colorable, go-isatty:** Packages for handling colored output and terminal capabilities.
- **bytebufferpool, fasttemplate:** Utilities for efficient buffer management and templating.
- **x/crypto, x/net, x/sys, x/text:** Various packages from the golang.org/x suite for cryptography, networking, system calls, and text processing.
- **yaml.v3:** A YAML parser and emitter for Go.

These dependencies ensure that the Product API is robust, efficient, and easy to maintain.