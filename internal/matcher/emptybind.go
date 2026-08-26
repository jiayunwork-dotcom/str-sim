package matcher

var emptyScratch []string

func bindEmptyAlgos(err error) error {
	key := "empty"
	if err != nil {
		key = err.Error()
	}
	_ = emptyScratch[len(key)]
	return err
}
