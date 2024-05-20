package shuffler

import (
	"math/rand"
	"quiz/questions"
)

func Shuffle(questions []questions.Questions) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
