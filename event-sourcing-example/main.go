package main

import (
	"encoding/json"
	"fmt"
)

type EventType string

const (
	DepositEvent  EventType = "deposit"
	WithdrawEvent EventType = "withdraw"
)

type Event struct {
	Type   EventType
	Amount float64
}

type EventStore struct {
	Events []Event
}

type BankAccount struct {
	balance float64
}

func (s *EventStore) SaveEvent(bankEvent Event) {
	s.Events = append(s.Events, bankEvent)
}

func (s *EventStore) GetEvents() []Event {
	return s.Events
}

func (a *BankAccount) ApplyEvent(event Event) {
	switch event.Type {
	case DepositEvent:
		a.balance += event.Amount
	case WithdrawEvent:
		a.balance -= event.Amount
	}
}

func (a *BankAccount) HandleDeposit(store *EventStore, amount float64) {
	event := Event{
		Type:   "deposit",
		Amount: amount,
	}
	store.SaveEvent(event)
	a.ApplyEvent(event)
}

func (a *BankAccount) HandleWithdraw(store *EventStore, amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("insufficient funds: current balance %.2f, attempted withdrawal %.2f",
			a.balance, amount)
	}

	event := Event{
		Type:   "withdraw",
		Amount: amount,
	}
	store.SaveEvent(event)
	a.ApplyEvent(event)
	return nil
}

func (a *BankAccount) RebuildFromEvents(events []Event) {
	a.balance = 0
	for _, event := range events {
		a.ApplyEvent(event)
	}
}

func main() {
	// Create our event store and bank account
	store := &EventStore{}
	account := &BankAccount{}

	// Perform some transactions
	fmt.Println("Making transactions...")
	account.HandleDeposit(store, 100)
	fmt.Printf("Balance after deposit: $%.2f\n", account.balance)

	account.HandleDeposit(store, 50)
	fmt.Printf("Balance after deposit: $%.2f\n", account.balance)

	err := account.HandleWithdraw(store, 30)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Balance after withdrawal: $%.2f\n", account.balance)

	// Show all stored events
	fmt.Println("\nStored Events:")
	eventsJSON, _ := json.MarshalIndent(store.GetEvents(), "", "  ")
	fmt.Println(string(eventsJSON))

	// Demonstrate rebuilding state from events
	fmt.Println("\nRebuilding account from events...")
	newAccount := &BankAccount{}
	newAccount.RebuildFromEvents(store.GetEvents())
	fmt.Printf("Rebuilt account balance: $%.2f\n", newAccount.balance)
}
