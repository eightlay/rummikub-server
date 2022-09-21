package game

type pack []*Piece

func createInitialPack() pack {
	b := pack{}

	for d := 0; d < decksNumber; d++ {
		for i := minNumber; i <= maxNumber; i++ {
			for _, c := range colors {
				p := createPiece(i, c, false)
				b = append(b, p)
			}
		}

		b = append(b, createPiece(jokerNumber, jokerColor, true))
	}

	return b
}