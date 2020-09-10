package computer

type unknownOpcodeError struct {
	errString string
}

func NewUnknownOpcodeError(s string) error {
	return &unknownOpcodeError{errString: s}
}

func (err *unknownOpcodeError) Error() string {
	return err.errString
}

type positionOutOfRangeError struct {
	errString string
}

func NewPositionOutOfRangeError(s string) error {
	return &positionOutOfRangeError{errString: s}
}

func (err *positionOutOfRangeError) Error() string {
	return err.errString
}