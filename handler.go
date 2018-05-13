package eventSourcing

import "sync"

func handle(event Event) error {
	eb := &stdEventBus
	list := eb.getListeners(event)
	var wg sync.WaitGroup
	ch := make(chan error, len(list))
	defer close(ch)
	wg.Add(len(list))
	for _, l := range list {
		go func(li Listener) {
			err := li.Handle(event)
			defer wg.Done()

			if err != nil {
				ch <- err
			}

		}(l)

		select {
		case err := <-ch:
			if err != nil {
				return err
			}

		}
	}
	wg.Wait()

	return nil
}
