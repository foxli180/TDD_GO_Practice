package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrorNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("word exists")
	ErrorWordNotExist = DictionaryErr("word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
		case ErrorNotFound:
			d[word] = definition
		case nil:
			return ErrorWordExists
		default:
			return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
		case ErrorNotFound:
			return ErrorWordNotExist
		case nil:
			d[word] = definition
		default:
			return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
