package observer

// Observer interface declares the method to notify all subscribers
type Observer interface {
	Notify(string)
}

// Publisher contains a list of observer to spread an message
type Publisher struct {
	ObserversList []Observer
}

// AddObserver adds an observer into ObserversList
func (p *Publisher) AddObserver(o Observer) {}

// RemoveObserver removes an observer from ObserversList
func (p *Publisher) RemoveObserver(o Observer) {}

// NotifyObservers notifies all observers about a message
func (p *Publisher) NotifyObservers(s string) {}
