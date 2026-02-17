package funciones

type ErrorMsg struct {
	Msg string
}

func (e *ErrorMsg) SetErrorMsg(msg string) {

	e.Msg = msg

}
