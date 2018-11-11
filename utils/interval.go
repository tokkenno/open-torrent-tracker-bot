package utils

import "time"

func SetInterval(someFunc func(), interval time.Duration, async bool) chan bool {
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					// This won't block
					go someFunc()
				} else {
					// This will block
					someFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}

		}
	}()

	return clear
}
