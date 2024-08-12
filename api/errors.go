package api

type ErrObjectAlreadyExists struct {
}

func (e ErrObjectAlreadyExists) Error() string {
	return "Object is already presented in (kind of) database."
}

type ErrObjectNotFound struct {
}

func (e ErrObjectNotFound) Error() string {
	return "Object is not presented in (kind of) database."
}
