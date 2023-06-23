# drinkablesa-api

[View Live](https://drinkablesa-api-b93993490e91.herokuapp.com)

The `drinkablesa-api` is a Go-based web application that provides access to data on various water municipalities in South Africa. The API allows users to retrieve information such as municipal water data, disinfectant levels, non-health aesthetic levels, and operational levels for each municipality.

## Installation

To set up the project,follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/your-username/drinkablesa-api.git
   ```

2. Change to the project directory:

   ```
   cd drinkablesa-api
   ```

3. Install Go dependencies:

   ```
   go mod download
   ```

## Configuration

The API requires a configuration file named `.env` in the root directory. Create the file and define the following environment variables:

```
PORT=8080
```

You can adjust the `PORT` value to your desired port number.

## Usage

To run the API, execute the following command:

```
go build main.go
go run main.go
```

The API will start running on the specified port (default: 8080).

## API Endpoints

The API provides the following endpoints:

- `GET /municipalities`: Retrieve information on all municipalities.
- `GET /municipalities/{name}`: Retrieve information on a specific municipality.


## Contributing

Contributions to the South African Municipal Data API are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
