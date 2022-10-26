// Generating text of a song 'The Twelve Days of Christmas'
package twelve

var verse = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming",
	"eleventh": "eleven Pipers Piping",
	"tenth":    "ten Lords-a-Leaping",
	"ninth":    "nine Ladies Dancing",
	"eighth":   "eight Maids-a-Milking",
	"seventh":  "seven Swans-a-Swimming",
	"sixth":    "six Geese-a-Laying",
	"fifth":    "five Gold Rings",
	"fourth":   "four Calling Birds",
	"third":    "three French Hens",
	"second":   "two Turtle Doves",
}

var wording = map[int]string{
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

// Verse generates a verse by its number
func Verse(i int) string {
	line := "On the " + wording[i] + " day of Christmas my true love gave to me"
	for j := i; j >= 1; j-- {
		if j == 1 && i != 1 {
			line += ", and " + verse[wording[j]]
		} else {
			line += ", " + verse[wording[j]]
		}

	}
	return line
}

// Song returns all the song as one string
func Song() string {
	var song = ""
	for i := 1; i <= len(wording); i++ {
		song += Verse(i) + "\n"
	}
	return song
}
