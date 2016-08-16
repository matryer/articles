package lineofsight

import "errors"

func BrilliantFunction() (*Thing, error) {
	something, err := GetSomething()
	if err != nil {
		return nil, err
	}
	defer something.Close()
	if !something.OK() {
		return nil, errors.New("something not right")
	}
	another, err := something.Else()
	if err != nil {
		return nil, err
	}
	another.Lock()
	defer another.Unlock()
	err = another.Update(1)
	if err != nil {
		return nil, err
	}
	return another.Thing(), nil
}
