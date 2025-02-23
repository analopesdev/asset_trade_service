# Asset Trade Service

Asset Trade Service is a Go-based application designed to facilitate the trading of assets between investors. It handles the creation of orders, processing of transactions, and updating of investor asset positions.

## Project Structure

```
/home/analopes/projects/asset_trade_service
├── internal
│   └── market
│       └── entity
│           ├── asset.go
│           ├── book.go
│           ├── investor.go
│           ├── order.go
│           ├── order_processor.go
│           └── transaction.go
└── README.md
```

### Key Components

- **Asset**: Represents an asset available for trading.
- **Investor**: Represents an investor participating in the market.
- **Order**: Represents a buy or sell order placed by an investor.
- **Transaction**: Represents a completed trade between a buying and selling order.
- **OrderProcessor**: Handles the processing of transactions.
- **Book**: Manages incoming and processed orders, and facilitates trading.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/analopesdev/asset_trade_service.git
    cd asset_trade_service
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application

1. Build the application:
    ```sh
    go build -o asset_trade_service
    ```

2. Run the application:
    ```sh
    ./asset_trade_service
    ```

## Usage

The application processes orders and matches them to create transactions. Orders can be added to the `IncomingOrders` channel, and processed orders will be available in the `ProcessedOrders` channel.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
