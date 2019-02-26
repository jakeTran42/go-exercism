package twelve

// Verse returns a verse from the song "The Twelve Days of Christmas".
// It takes in an int (the verse number) and returns a string.
func Verse(line int) string {
	ordinals := map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}

	gifts := [12]string{
		"a Partridge in a Pear Tree.", "two Turtle Doves, ", "three French Hens, ",
		"four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, ",
	}
	if line > 1 {
		gifts[0] = "and a Partridge in a Pear Tree."
	}

	verse := "On the " + ordinals[line] + " day of Christmas my true love gave to me: "
	list := ""
	for i := 0; i < line; i++ {
		list = gifts[i] + list
	}
	verse += list

	return verse
}

// Song returns a string containing the entirety of "The Twelve Days of Christmas".
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}
