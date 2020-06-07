package mydic

import (
	"errors"
)
//Dictionary is
type Dictionary map[string]string



var errNotFound = errors.New("Not Found")
var errExist = errors.New("Already Exist")
var errNotUpdate = errors.New("Not Update")

// Search is
func (d Dictionary) Search(word string)(string, error){
	value, exist := d[word]
	if(exist){
		return value, nil
	}
	return "", errNotFound
}

//Add is..
func (d Dictionary) Add(word , def string)error{
	_, err := d.Search(word)
	if err == errNotFound{
		d[word] = def
	}else if err == nil{
		return errExist
	}
	return nil
}

// Update is...
func (d Dictionary) Update(word, def string) error{
	_, err := d.Search(word)
		switch err{
		case nil:
			d[word] = def
		case errNotFound:
			return errNotUpdate
	}
	return nil
}