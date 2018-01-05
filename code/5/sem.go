// implementation of semaphore using a buffer channel.
package main

type sem chan struct{}

func newsem(permits int) sem {
	return make(chan struct{}, permits)
}

func (s sem) acquire() {
	s <- struct{}{}
}

func (s sem) release() {
	<-s
}

func (s sem) wait() {
	for i := 0; i < cap(s); i++ {
		s.acquire()
	}
}
