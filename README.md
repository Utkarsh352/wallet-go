# Project Overview

The project involves creating a simple cryptocurrency wallet system using Go. The application allows users to create wallets, perform transactions, and maintain a ledger of transactions.

Directory Structure

- `models/transaction.go`: Contains the `Transaction` struct and the `RecordTransaction` function.
- `models/wallet.go`: Contains the `Wallet` struct with methods for sending cryptocurrency and checking balance.
- `utility/helper.go`: Provides the `FormatTimestamp` function for formatting timestamps.
- `main.go`: Contains the main application logic for initializing wallets, performing transactions, and displaying results.

Key Components

1. Wallet Model (`models/wallet.go`)
   - `Wallet`: Struct representing a cryptocurrency wallet with fields for address, balance, and a mutex for thread safety.
   - `NewWallet()`: Function to create a new wallet with a unique address.
   - `Send()`: Method to transfer funds between wallets, ensuring thread safety and updating the ledger.
   - `CheckBalance()`: Method to retrieve the current balance of the wallet.

2. Transaction Model (`models/transaction.go`)
   - `Transaction`: Struct representing a transaction with sender, receiver, amount, and timestamp.
   - `RecordTransaction()`: Function to record a transaction in the ledger, ensuring thread safety using a mutex.

3. Utility Functions (`utility/helper.go`)
   - `FormatTimestamp()`: Function to format timestamps into a human-readable format (hh:mm:ss dd/mm/yyyy).

4. Main Application (`main.go`)
   - Initializes wallets and performs sample transactions.
   - Uses the `Send` method to transfer funds and updates the transaction ledger.
   - Outputs the balances of all wallets and displays the transaction ledger.

How to Run

1. Clone the Repository
   git clone https://github.com/Utkarsh352/wallet-go

2. Navigate to the Project Directory
   cd wallet-go

3. Run the Application
   go run main.go

#  Video Demonstration

I have created a video demonstration to showcase the functionality of the cryptocurrency wallet application. You can view the video using the following link:

[Video Demonstration](https://youtu.be/O985LebqJfY)

Notes

- Ensure that you have Go installed on your system to run the application.
- The video demonstration provides a walkthrough of the application’s features and functionality.

If you have any questions or need further clarification, please feel free to reach out.
