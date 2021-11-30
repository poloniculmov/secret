package santa

import (
	"math/rand"
	"time"
)

type Santa struct {
	firstName   string
	lastName    string
	email       string
	address     string
	phoneNumber string
	mentions 	string
	wishes      string
	sizing      string
}

func (s Santa) FirstName() string {
	return s.firstName
}

func (s Santa) LastName() string {
	return s.lastName
}

func (s Santa) FullName() string {
	return s.firstName + " " + s.lastName
}

func (s Santa) Email() string {
	return s.email
}

func (s Santa) Address() string {
	return s.address
}

func (s Santa) PhoneNumber() string {
	return s.phoneNumber
}

func (s Santa) Mentions() string {
	return s.mentions
}

func (s Santa) Wishes() string {
	return s.wishes
}

func (s Santa) Sizing() string {
	return s.sizing
}

func New(firstName string, lastName string, email string, address string, phoneNumber string, mentions string,wishes string, sizing string) Santa {
	s := Santa{firstName, lastName, email, address, phoneNumber, mentions, wishes, sizing}
	return s
}

func NewFromCSV(line []string) Santa {
	address := line[4] + " " + line[3] + " " + line[2]
	return New(line[1], line[0], line[5], address, line[6], line[7], line[8], line[9])
}

func Draw(santas []Santa) map[int]int {
	result := make(map[int]int)
	candidates := shuffleSlice(makeRange(len(santas)))
	for index, _ := range santas {
		result[index] = candidates[index]
	}

	return result
}

func shuffleSlice(slice []int) []int {
	shuffled := false
	rand.Seed(time.Now().UnixNano())
	for !shuffled {
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
		shuffled = true
		for i, e := range slice {
			if i == e {
				shuffled = false
			}
		}
	}
	return slice
}

func makeRange(size int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = i
	}
	return a
}