package sim

var algoClosed chan error

func bindUnknownAlgo(err error) error {
	return err
}
