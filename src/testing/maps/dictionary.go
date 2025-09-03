package dictionary

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryErr("unknown word")
	ErrAlreadyExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		d[word] = def
		return nil
	}

	if err == nil {
		return ErrAlreadyExists
	}

	return nil

}

func (d Dictionary) Update(word, def string) {
	d[word] = def
}
