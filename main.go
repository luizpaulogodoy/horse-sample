package main

import "log"

type Horse struct {
	Speed      Attribute
	Resistence Attribute
}

type Attribute struct {
	Gens []Gen
}

type Gen struct {
	Dominant bool
	Value    int8
}

func main() {
	horse1 := &Horse{
		Speed: Attribute{
			Gens: []Gen{
				{
					Dominant: true,
					Value:    63,
				},
				{
					Dominant: true,
					Value:    87,
				},
			},
		},
		Resistence: Attribute{
			Gens: []Gen{
				{
					Dominant: false,
					Value:    48,
				},
				{
					Dominant: true,
					Value:    39,
				},
			},
		},
	}

	horse2 := &Horse{
		Speed: Attribute{
			Gens: []Gen{
				{
					Dominant: false,
					Value:    93,
				},
				{
					Dominant: true,
					Value:    47,
				},
			},
		},
		Resistence: Attribute{
			Gens: []Gen{
				{
					Dominant: true,
					Value:    73,
				},
				{
					Dominant: false,
					Value:    68,
				},
			},
		},
	}

	son := horse1.Breed(horse2)

	log.Println("Son:")
	log.Println("Speed:", son.Speed)
	log.Println("Resistence:", son.Resistence)
}

func (s *Horse) Breed(h *Horse) (son *Horse) {
	speedPos := s.Speed.Possibilities(h.Speed)
	resistencePos := s.Resistence.Possibilities(h.Resistence)
	son = &Horse{
		Speed: Attribute{
			Gens: chooseGens(speedPos),
		},
		Resistence: Attribute{
			Gens: chooseGens(resistencePos),
		},
	}
	return
}

func (s Attribute) Possibilities(a Attribute) [][]Gen {
	// xX | yX
	// xy, xX, Xy, XX

	p1 := []Gen{
		s.Gens[0],
		a.Gens[0],
	}

	p2 := []Gen{
		s.Gens[0],
		a.Gens[1],
	}

	p3 := []Gen{
		s.Gens[1],
		a.Gens[0],
	}

	p4 := []Gen{
		s.Gens[1],
		a.Gens[1],
	}

	return [][]Gen{p1, p2, p3, p4}
}

func chooseGens(pos [][]Gen) []Gen {
	return pos[rand()]
}

func rand() int8 {
	return 3
}
