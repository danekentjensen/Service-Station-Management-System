package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Define structs for customers and service history
type Customer struct {
	ID      int
	Name    string
	Car     string
	LastOil time.Time
}

type ServiceHistory struct {
	CustomerID int
	Service    string
	Date       time.Time
}

// Helper function to parse dates
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func main() {
	// Load customers
	customers, err := loadCustomers("customers.csv")
	if err != nil {
		fmt.Println("Error loading customers:", err)
		return
	}

	// Load service history
	services, err := loadServiceHistory("service_history.csv")
	if err != nil {
		fmt.Println("Error loading service history:", err)
		return
	}

	// Generate report
	fmt.Printf("Report for %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("Customer ID\tName\tCar\t\tService\t\tLast Service Date\n")
	for _, service := range services {
		customer := findCustomerByID(customers, service.CustomerID)
		if customer != nil {
			fmt.Printf("%d\t%s\t%s\t%s\t%s\n",
				customer.ID, customer.Name, customer.Car, service.Service,
				service.Date.Format("2006-01-02"))
		}
	}

	// Reminder for customers needing service
	for _, customer := range customers {
		if customer.LastOil.AddDate(0, 6, 0).Before(time.Now()) {
			fmt.Printf("Customer %s needs a service reminder\n", customer.Name)
		}
	}
}

// Load customers from the CSV file
func loadCustomers(filePath string) ([]Customer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var customers []Customer
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		id, _ := strconv.Atoi(record[0])
		name := record[1]
		car := record[2]
		lastOil, err := parseDate(record[3])
		if err != nil {
			fmt.Printf("Error parsing date for customer %s: %v\n", name, err)
			continue
		}

		customers = append(customers, Customer{
			ID:      id,
			Name:    name,
			Car:     car,
			LastOil: lastOil,
		})
	}

	return customers, nil
}

// Load service history from the CSV file
func loadServiceHistory(filePath string) ([]ServiceHistory, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var services []ServiceHistory
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		customerID, _ := strconv.Atoi(record[0])
		service := record[1]
		date, err := parseDate(record[2])
		if err != nil {
			fmt.Printf("Error parsing service date for customer ID %d: %v\n", customerID, err)
			continue
		}

		services = append(services, ServiceHistory{
			CustomerID: customerID,
			Service:    service,
			Date:       date,
		})
	}

	return services, nil
}

// Find a customer by ID
func findCustomerByID(customers []Customer, id int) *Customer {
	for _, customer := range customers {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}
