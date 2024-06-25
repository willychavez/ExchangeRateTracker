# ExchangeRateTracker

## Overview

This project consists of two Go services: a server (`server.go`) that fetches the current exchange rate from USD to BRL and stores it in a SQLite database, and a client (`client.go`) that requests the exchange rate from the server and saves it to a text file.

## Project Structure

- `server.go`: Service that fetches the exchange rate from an external API, stores it in a database, and exposes an HTTP endpoint.
- `client.go`: Client that requests the exchange rate from the server and saves it to a text file.

## Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- [SQLite3](https://www.sqlite.org/download.html)

## Setup

1. Clone this repository:

    ```sh
    git clone https://github.com/your-username/your-repo.git
    cd your-repo
    ```

2. Install Go dependencies:

    ```sh
    go mod tidy
    ```

3. Ensure you have SQLite3 installed. To install on Ubuntu, use:

    ```sh
    sudo apt-get install sqlite3
    ```

## Usage

### Running the Server

1. Create and start the database:

    ```sh
    sqlite3 quotes.db
    ```

    In the SQLite console, run:

    ```sql
    CREATE TABLE IF NOT EXISTS quotes (
        id INTEGER PRIMARY KEY,
        bid TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    .quit
    ```

2. Start the server:

    ```sh
    cd server/
    go run server.go
    ```

    The server will listen on port `8080` and has an endpoint `/cotacao` to fetch the exchange rate.

### Running the Client

1. In another terminal, run the client:

    ```sh
    cd client/
    go run client.go
    ```

    The client will request the exchange rate from the server and save the result to a `cotacao.txt` file.

## Code Explanation

### Server (`server.go`)

- Opens a connection to the SQLite database.
- Creates a `quotes` table if it does not exist.
- Defines an HTTP endpoint `/cotacao` that:
  - Fetches the exchange rate from the `economia.awesomeapi.com.br` API.
  - Saves the exchange rate to the database.
  - Returns the exchange rate as a JSON response.

### Client (`client.go`)

- Defines a context with a timeout for the request.
- Requests the exchange rate from the server.
- Saves the response to a `cotacao.txt` file.

## Examples

### API Response

A typical response from the server's `/cotacao` API will be:

```json
{
  "bid": "5.42"
}
```

This value will be saved to the database by the server and to a `cotacao.txt` file by the client, in the format 'Dólar:{valor}'.

## Common Issues

- **Server not running**: Ensure the server is running before starting the client.
- **Network issues**: Verify the client can reach the server on the correct port.
- **SQLite3 not installed**: Ensure SQLite3 is installed and properly configured.

## Contribution

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin feature/fooBar`).
5. Create a new Pull Request.

---

Feel free to adjust or expand this README as needed to suit the specifics of your project.
