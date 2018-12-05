package gopool

// task interface ,your task need implements Consume func;
// then use it to handle your business logic
type Task interface {
	Consume() error
}
