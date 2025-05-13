package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Contact holds individual contact information
type Contact struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// ContactBook manages contacts using a map
type ContactBook struct {
	contacts map[string]Contact
}

// Create a new contact book
func NewContactBook() *ContactBook {
	return &ContactBook{
		contacts: make(map[string]Contact),
	}
}

// Add a contact to the book
func (cb *ContactBook) Add() {
	name := readInput("Enter Name: ")
	if _, exists := cb.contacts[strings.ToLower(name)]; exists {
		fmt.Println("Contact already exists.")
		return
	}

	email := readInput("Enter Email: ")
	phone := readInput("Enter Phone: ")
	address := readInput("Enter Address: ")

	cb.contacts[strings.ToLower(name)] = Contact{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	}

	fmt.Println("Contact added successfully.")
}

// Update an existing contact selectively
func (cb *ContactBook) Update() {
	name := readInput("Enter the name of the contact to update: ")
	key := strings.ToLower(name)
	contact, exists := cb.contacts[key]

	if !exists {
		fmt.Println("Contact not found.")
		return
	}

	fmt.Println("Press Enter to keep current value.")

	newEmail := readInput(fmt.Sprintf("Enter new Email [%s]: ", contact.Email))
	newPhone := readInput(fmt.Sprintf("Enter new Phone [%s]: ", contact.Phone))
	newAddress := readInput(fmt.Sprintf("Enter new Address [%s]: ", contact.Address))

	if newEmail != "" {
		contact.Email = newEmail
	}
	if newPhone != "" {
		contact.Phone = newPhone
	}
	if newAddress != "" {
		contact.Address = newAddress
	}

	cb.contacts[key] = contact
	fmt.Println("Contact updated successfully.")
}

// Delete a contact by name
func (cb *ContactBook) Delete() {
	name := readInput("Enter the name of the contact to delete: ")
	key := strings.ToLower(name)

	if _, exists := cb.contacts[key]; !exists {
		fmt.Println("Contact not found.")
		return
	}

	delete(cb.contacts, key)
	fmt.Println("Contact deleted successfully.")
}

// View all contacts
func (cb *ContactBook) View() {
	if len(cb.contacts) == 0 {
		fmt.Println("No contacts to show.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, c := range cb.contacts {
		fmt.Printf("Name: %s\nEmail: %s\nPhone: %s\nAddress: %s\n\n", c.Name, c.Email, c.Phone, c.Address)
	}
}

// Helper to read user input
func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Main menu
func main() {
	contactBook := NewContactBook()

	for {
		fmt.Println("\nContact Book Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Update Contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. View Contacts")
		fmt.Println("5. Exit")

		choice := readInput("Enter your choice (1-5): ")

		switch choice {
		case "1":
			contactBook.Add()
		case "2":
			contactBook.Update()
		case "3":
			contactBook.Delete()
		case "4":
			contactBook.View()
		case "5":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
