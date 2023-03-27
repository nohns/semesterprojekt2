package domain




type domain struct {

}

func new() *domain {
	return &domain{}
}

func (d domain) Register() (string, error) {
	return "", nil
}



func (d domain) GetLock() (bool, error) {

	return false, nil
}

func (d domain) SetLock() (bool, error) {

	return false, nil
}