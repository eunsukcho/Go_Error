package errlog

import "runtime"

var Unknown_Error string = "Unknown_Error"
var Uri_Binding_Error string = "URI Data Binding Error."
var Json_Binding_Error string = "Json Data Biding Error."
var Database_Connection_Error string = "Database Coneection Error."
var Database_Value_Error string = "Database Error"

/* Keycloak Error Content */
var Get_Token_Error string = "Get Token Error"
var Marshal_Error string = "Marshal_Error"
var Unmarshal_Error string = "Unmarshal Error"
var Access_Denied_Error string = "Access Denied Error"
var ReadIO_Error string = "Read Body Io Buffer Error"
var Client_Result_Error string = "Client Result Error"

/* k8s Error */
var Duplicate_Error string = "Already Exists"
var Create_Error string = "Not Created"

type ErrorContent struct {
	Method       string
	ErrorContent string
	CodeLine     int
}

func (ErrorContent) CustomError(errContent string) *ErrorContent {
	var err *ErrorContent
	function, _, line, _ := runtime.Caller(1)
	err = &ErrorContent{
		Method:       runtime.FuncForPC(function).Name(),
		ErrorContent: errContent,
		CodeLine:     line,
	}
	return err
}

func (e ErrorContent) Error(debugMode bool) *ErrorContent {
	if debugMode {
		errContent := &ErrorContent{Method: e.Method, ErrorContent: e.ErrorContent, CodeLine: e.CodeLine}
		return errContent
	} else {
		errContent := &ErrorContent{ErrorContent: e.ErrorContent}
		return errContent
	}
}

func (e ErrorContent) ModalError(debugMode bool) *ErrorContent {
	if debugMode {
		return &ErrorContent{
			Method:       e.Method,
			ErrorContent: e.ErrorContent,
			CodeLine:     e.CodeLine,
		}
	} else {
		return &ErrorContent{
			ErrorContent: e.ErrorContent,
		}
	}
}
