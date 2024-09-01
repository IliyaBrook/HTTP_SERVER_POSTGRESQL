
# Dog Breeds Search Application

This is a Go application that allows users to search for dog breeds using the Dog API. The application serves an HTML page with a form where users can search for breeds, and the results are displayed dynamically.

## Features

- Search for dog breeds by name.
- Display detailed information about each breed, including images, breed group, life span, temperament, weight, and height.
- Interact with the Dog API to fetch breed data.

## Prerequisites

Before running the application, ensure you have the following:

- Go installed on your machine.
- An API key from the Dog API. You can obtain one by signing up at [The Dog API](https://thedogapi.com/).

## Installation

1. Clone the repository:
    ```
    git clone dogBreedsSearch
    cd dogBreedsSearch
    ```

2. Create a `.env` file in the root directory with the following content:
    ```
    API_KEY=your-dog-api-key
    PORT=8080
    ```

    - `API_KEY`: Your API key from the Dog API.
    - `PORT`: The port on which the application will run (default is `8080`).

3. Install the required Go packages:
    ```
    go mod tidy
    ```

4. Run the application:
    ```
    go run main.go
    ```

5. Open your browser and navigate to `http://localhost:8080` to use the application.

## Project Structure

- `main.go`: The main entry point of the application.
- `pages/`: Contains HTML templates and scripts used in the application.

## Usage

- Enter a breed name in the search field.
- Click "Submit" to see details about the breed.
- Click "Clear" to reset the search.

## License

This project is part of the "Golang Ninja" course on Udemy and is for educational purposes.
