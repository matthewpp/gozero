package exam

import "fmt"

type ErrorState int

func (e ErrorState) Error() string {
	return fmt.Sprintf("error state is %d, mean user error", e)
}

func RetErrState() error {
	e := ErrorState(5)
	return e
}
