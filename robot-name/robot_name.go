package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// Robot is a struct that holds the robot's name
type Robot struct {
	name string
}

// existingNames is used to keep track of created names
var existingNames = map[string]bool{}

// Name is a Robot method that creates a random name if the name is not
// set or if it was reset
func (bot *Robot) Name() (string, error) {
	if bot.name == "" {
		nameExists := true

		for nameExists || bot.name == "" {
			bot.name = ""
			for len(bot.name) < 2 {
				newSeed := rand.NewSource(time.Now().UnixNano())
				newRand := rand.New(newSeed)

				asciiCode := newRand.Intn(26) + 65
				bot.name += string(asciiCode)
			}
			for len(bot.name) < 5 {
				newSeed := rand.NewSource(time.Now().UnixNano())
				newRand := rand.New(newSeed)

				num := newRand.Intn(10)
				bot.name += strconv.Itoa(num)
			}

			_, nameExists = existingNames[bot.name]
			existingNames[bot.name] = true
		}
	}

	return bot.name, nil
}

// Reset is a Robot method that sets the Robot's name to an empty
// string. If Name is called again, the Robot will get a new name.
func (bot *Robot) Reset() {
	bot.name = ""
}
