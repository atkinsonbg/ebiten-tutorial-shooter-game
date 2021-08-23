package utils

type Score struct {
	currentScore int
}

func init() {

}

func (s *Score) GetScore() int {
	return s.currentScore
}

func (s *Score) UpdateScore(inc int) {
	s.currentScore += inc
}
