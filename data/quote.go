package data

import "math/rand"

type Quote struct {
	Quote  string
	Author string
}

type Quotes []Quote

var QuotesData = Quotes{
	{
		Quote:  "Do I contradict myself? Very well then I contradict myself. (I am large, I contain multitudes)",
		Author: "Walt Whitman",
	},
	{
		Quote:  "Theory is good, but it doesn't prevent things from existing.",
		Author: "Jean-Martin Charcot",
	},
	{
		Quote:  "The safest general characterization of the European philosophical tradition is that it consists of a series of footnotes to Plato.",
		Author: "Alfred North Whitehead",
	},
	{
		Quote:  "What's new? is an interesting and broadening eternal question, but one which, if pursued exclusively, results only in an endless parade of trivia and fashion, the silt of tomorrow. I would like, instead, to be concerned with the question 'What is best?', a question which cuts deeply rather than broadly, a question whose answers tend to move the silt downstream.",
		Author: "Robert M. Pirsig",
	},
	{
		Quote:  "Part of the inhumanity of the computer is that, once it is competently programmed and working smoothly, it is completely honest.",
		Author: "Isaac Asimov",
	},
	{
		Quote:  "Program testing can be a very effective way to show the presence of bugs, but it is hopelessly inadequate for showing their absence.",
		Author: "Edsger W. Dijkstra",
	},
	{
		Quote:  "Programmers have to fight against the two most destructive forces in the universe: entropy and stupidity.",
		Author: "Damian Conway",
	},
	{
		Quote:  "A design is 'simple' if it follows these rules: Runs all the tests; Contains no duplication; Expresses the intent of the programmer; Minimizes the number of classes and methods.",
		Author: "Kent Beck",
	},
	{
		Quote:  "To be an expert in a field that changes from one day to the next is akin to placing your hand in a running river; you can trap a small bit of water for the moment, but once you lift your hand again the river rushes on...",
		Author: "Kelly Ripley Feller",
	},
	{
		Quote:  "The programmer, like the poet, works only slightly removed from pure thought-stuff. He builds his castles in the air, from air, creating by exertion of the imagination. Few media of creation are so flexible, so easy to polish and rework, so readily capable of realizing grand conceptual structures.",
		Author: "Fred Brooks",
	},
	{
		Quote:  "If you wish to make an apple pie from scratch you must first invent the universe.",
		Author: "Carl Sagan",
	},
	{
		Quote:  "There are only two kinds of programming languages: the ones people complain about and the ones nobody uses.",
		Author: "Bjarne Stroustrup",
	},
	{
		Quote:  "Syntactic sugar causes cancer of the semicolon.",
		Author: "Alan Perlis",
	},
	{
		Quote:  "Always plan under the assumption that those who become involved with the project later will have or have developed the experience and insight to improve on the design.",
		Author: "James C. Scott",
	},
	{
		Quote:  "Power is also like love, easier to experience than to define or measure, but no less real for that.",
		Author: "Joseph Nye",
	},
	{
		Quote: "The man who works so moderately as to be able to work constantly not only preserves his health the longest, but, in the course of a year, executes the greatest quantity of work.",
		Author: "Adam Smith",
	},
	{
		Quote:  "In the seen, only the seen. In the heard, only the heard. In the sensed, only the sensed. In the cognized, only the cognized. That is how you should train yourself. When for you, there will be only the seen in the seen...then there is no you in that. That is the end of suffering",
		Author: "Buddha",
	},
	{
		Quote:  "All we have to decide is what to do with the time that is given us.",
		Author: "J.R.R. Tolkien",
	},
	{
		Quote: "You must always work not just within, but below your means. If you can handle three elements, handle only two. If you can handle ten, then handle only five. In that way, the ones you do handle, you handle with more ease, more mastery, and you create a feeling of strength in reserve.",
		Author: "Pablo Picasso",
	}
}

func GetRandomQuote() Quote {
	return QuotesData[rand.Intn(len(QuotesData))]
}
