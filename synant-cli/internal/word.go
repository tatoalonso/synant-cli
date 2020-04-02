package wordscli

//ListOfWords representation of a list of words into data struct
type ListOfWords struct {
	Sinonimos []Sinonimo
	Antonimos []Antonimo
}

//Sinonimo representation as individual word
type Sinonimo struct {
	Sinonimo string
}

//Antonimo representation as individual word
type Antonimo struct {
	Antonimo string
}

// WordsRepo definiton of methods to access word
type WordsRepo interface {
	GetWords(typeOfWord string, word string) (ListOfWords, error)
}

// WordsRepoStorage definition if method to persist List of words
type WordsRepoStorage interface {
	SaveWords(typeOfWord string, word string, list ListOfWords) error
}
