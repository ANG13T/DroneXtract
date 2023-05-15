type NotDATFileError struct {
	filename string
}

func (e NotDATFileError) Error() string {
	return fmt.Sprintf("*** ERROR: %s is not a recognized DJI DAT file. ***", e.filename)
}