package workchain

func WorkChain(works ...*Work) *Work {
	main := NewWork(nil)
	count := len(works)
	for i, w := range works {
		if i != 0 {
			w.prev = works[i-1]
		}
		if i != count-1 {
			w.next = works[i+1]
		}
	}
	if count > 0 {
		main.next = works[0]
	}
	return main
}

func WorkRing(works ...*Work) *Work {
	w := WorkChain(works...)
	if len(works) > 0 {
		w.prev = works[len(works)-1]
		w.prev.next = w
	}
	return w
}
