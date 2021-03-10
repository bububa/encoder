package category

const (
	NotFoundError Error = "not found in encoder"
	BoundsError   Error = "out of bounds"
)

type Error string

func (e Error) Error() string {
	return string(e)
}
