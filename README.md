# URL Shortener

A simple URL shortener application built with Go. This application allows users to shorten long URLs, making them easier to share.

## Features

- Shorten long URLs
- Redirect to the original URL using the shortened link
- Simple web interface

## Prerequisites

- Go 1.16 or later

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/bajkacper/URL-shortener.git
    cd URL-shortener
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## Usage

1. Open your web browser and go to `http://localhost:8080`.
2. Enter a URL you want to shorten and click the "Shorten" button.
3. Copy the shortened URL and use it to redirect to the original URL.

## License

This project is licensed under the MIT License. 


## Acknowledgements

- Inspired by various URL shortener services available online.
