package gopool

// task interface ,your task need imp Consume func
type Task interface {
	Consume() error
}
