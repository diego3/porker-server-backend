package core

// TODO separar num pacote diferente o Observer Pattern
type Listener struct {
	Event string
}
// qlq objeto que implementar esse método será um listener
// verificado em tempo de compilação(TODO testar)
func (lis *Listener) Listen(eventData interface{}){}

type EventManager struct {
	Listeners []Listener
}
func (e *EventManager) subscribe(event string, l Listener) {
	// armazena o evento no qual o objeto está se inscrevendo
	l.Event = event
	e.Listeners = append(e.Listeners, l)
}
func (e *EventManager) trigger(event string, eventData interface{}) {
	for _, l := range e.Listeners {
		if l.Event == event {
			l.Listen(eventData)
		}
	}
}
