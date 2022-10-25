package main

type Fibb struct {
	Id         int
	funkyslice []int
	funkyfib   []int
}

func (f *Fibb) CalcFibb() {
	f.funkyfib = make([]int, 21)
	f.funkyslice = make([]int, 10)
	f.funkyfib[0] = 0
	f.funkyfib[1] = 1
	for i := 0; i < 29; i++ {
		if i < 21 && i > 1 {
			prevfib := f.funkyfib[i-1]
			preprefib := f.funkyfib[i-2]
			f.funkyfib[i] = prevfib + preprefib
		}
		if i == 20 {
			f.funkyslice[0] = f.funkyfib[i-1]
			f.funkyslice[1] = f.funkyfib[i]
		}

		if i > 20 {
			prev := f.funkyslice[i-20]
			prevprev := f.funkyslice[i-21]
			f.funkyslice[i-19] = prev + prevprev
		}
	}
}

func (f *Fibb) AnotherFibb() {
	f.funkyslice = make([]int, 30)
	f.funkyslice[0] = 0
	f.funkyslice[1] = 1
	for i := 2; i < 30; i++ {
		prevfib := f.funkyslice[i-1]
		preprefib := f.funkyslice[i-2]
		f.funkyslice[i] = prevfib + preprefib
	}
	f.funkyslice = f.funkyslice[20:30]
}

func (f *Fibb) GetFunky() []int {
	return f.funkyslice
}

func (f *Fibb) GetFibb() []int {
	return f.funkyfib
}
