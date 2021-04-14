package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	nums := []int{3, 1, 2, 4}
	x := 4
	word := "hi my name is riyadlusholihin and i am interested in the opportunity you are giving"

	answerNumberOne := answerNumberOne(nums)
	answerNumberTwo := answerNumberTwo(nums, x)
	answerNumberThree := answerNumberThree(word, x)

	logrus.Info("======================ANSWER NUMBER 1==========================")
	logrus.Infof("I Have Array : %v", nums)
	logrus.Infof("Answer for number one : %v", answerNumberOne)
	logrus.Info("======================ANSWER NUMBER 2==========================")
	logrus.Infof("I Have Array : %v", nums)
	logrus.Infof("Answer for number two : %v", answerNumberTwo)
	logrus.Info("======================ANSWER NUMBER 3==========================")
	logrus.Infof("I Have word : %v", word)
	logrus.Infof("Answer for number three : %v", answerNumberThree)

}

func answerNumberOne(num []int) []int {
	var (
		n int
		l []int
	)

	f := make(map[int]int)
	for _, v := range num {
		for _, k := range num {
			if v != k {
				n = v - k
				if n < 0 {
					f[v] = 0
				} else {
					f[v] = 1
				}
			}
		}

	}
	for i, c := range f {
		if c > 0 {
			l = append(l, i)
		}
	}
	return l
}

func answerNumberTwo(num []int, x int) []int {
	var (
		s int
		k []int
	)
	for _, v := range num {
		s = x / v

		if s != x {
			k = append(k, v)
		}
	}
	return k
}

func answerNumberThree(word string, x int) []string {
	var new []string
	arrayWord := strings.Split(word, " ")

	for _, w := range arrayWord {
		if len(w) == x {
			new = append(new, w)
		}
	}

	return new
}
