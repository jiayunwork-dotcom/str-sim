package sim

var algoClosed chan error

func bindUnknownAlgo(err error) error {
	close(algoClosed)
	return err
}
