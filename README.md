# Service Station Management System
A Go program that processes customer and service history from CSV files to generate detailed reports and service reminders. Features include data parsing, normalization, error handling, and identifying customers needing service based on their last service date.

## Features
- Parses and normalizes customer and service data from CSV files.
- Handles inconsistent date formats and missing information gracefully.
- Generates a detailed report with customer details, services performed, and last service dates.
- Provides service reminders based on the last service date.

## Requirements
- Go 1.20 or later
- CSV files with customer and service history data

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/danekentjensen/Service-Station-Management-System.git
2. Navigate to the project directory
   ```bash
   cd service-station-management-system
3. Ensure the CSV files (customers.csv and service_history.csv) are in the project directory.

## Usage
1. Run the program:
   ```bash
   go run service_station.go
2. Ensure the output includes:
   - A report with customer details and services performed
   - Reminders for customers needing service.

## File Format
- customers.csv
  ```csv
  ID,Name,Car,LastOil
  1,Dane,Dodge Charger,2024-11-17
  2,Chelsea,Hundai Ioniq,2024-07-24
  3,Lisa,Chevy Equinox,2024-04-26
  4,Bob,Ford Mustang,2025-01-05

- service_history.csv
  ```csv
  CustomerID,Service,Date
  1,Oil Change,2024-11-17
  2,Oil Change,2024-07-24
  3,Oil Change,2024-04-26
  4,Oil Change,2025-01-05

## Project Structure
- service_station.go
- customers.csv
- service_history.csv

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.

