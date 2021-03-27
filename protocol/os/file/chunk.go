package file

type chunk struct {
	File *File

 	Index int

	Size int
	Data []byte
}

