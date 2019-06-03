package observer

import "fmt"

// Observer interface declares the method to notify all subscribers
type Observer interface {
	Notify(string)
}

// Publisher contains a list of observer to spread an message
type Publisher struct {
	ObserversList []Observer
}

// AddObserver adds an observer into ObserversList
func (p *Publisher) AddObserver(o Observer) {
	p.ObserversList = append(p.ObserversList, o)
}

// RemoveObserver removes an observer from ObserversList
func (p *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range p.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	p.ObserversList = append(p.ObserversList[:indexToRemove],
		p.ObserversList[indexToRemove+1:]...)
}

// NotifyObservers notifies all observers about a message
func (p *Publisher) NotifyObservers(s string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", s)
	for _, observer := range p.ObserversList {
		observer.Notify(s)
	}
}
