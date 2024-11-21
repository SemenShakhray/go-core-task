package main

type semaphorWait struct {
	sem chan struct{}
}

func newSemaphor(delta int) semaphorWait {
	return semaphorWait{
		sem: make(chan struct{}, delta),
	}
}

func (s *semaphorWait) Add(delta int) {
	for i := 0; i < delta; i++ {
		s.sem <- struct{}{}
	}
}

func (s *semaphorWait) Done(delta int) {
	for i := 0; i < delta; i++ {
		<-s.sem
	}
}

// func main() {
// 	sliceIn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	delta := len(sliceIn)

// 	s := newSemaphor(delta)
// 	ch := make(chan int)

// 	for _, n := range sliceIn {
// 		go func() {
// 			time.Sleep(1000 * time.Millisecond)
// 			ch <- n
// 			s.Add(1)
// 		}()
// 	}

// 	go func() {
// 		s.Done(delta)
// 		close(ch)
// 	}()

// 	var sliceOut []int
// 	for number := range ch {
// 		sliceOut = append(sliceOut, number)
// 	}
// 	fmt.Println(sliceOut)
// }
