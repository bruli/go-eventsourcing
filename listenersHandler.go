package eventSourcing

import "sync"

type listenersHandler struct {
	listeners map[string][]Listener
}

func (lh *listenersHandler) handle(ev Event) error {
	list := lh.getListeners(ev)
	var wg sync.WaitGroup
	wg.Add(len(list))
	chErr := make(chan error, len(list))
	defer close(chErr)
	for _, l := range list {
		go func(li Listener) {
			defer wg.Done()
			chErr <- li.Handle(ev)
		}(l)

		select {
		case err := <-chErr:
			if err != nil {
				return err
			}
		}
	}

	wg.Wait()

	return nil
}

func (lh *listenersHandler) getListeners(ev Event) []Listener {
	return lh.listeners[ev.Name()]
}
