package main

type Organiser struct {
	Path string
}

// NewOrganiser is the constructor function for the organiser struct
func NewOrganiser(path string) *Organiser {
	return &Organiser{
		Path: path,
	}
}
