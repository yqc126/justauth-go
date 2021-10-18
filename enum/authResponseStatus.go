package enum

type AuthResponseStatus struct {
	Code int
	Msg  string
}

var (
	/**
	 * 2000：正常；
	 * other：调用异常，具体异常内容见{@code msg}
	 */
	SUCCESS                = AuthResponseStatus{Code: 2000, Msg: "Success"}
	FAILURE                = AuthResponseStatus{Code: 5000, Msg: "Failure"}
	NOT_IMPLEMENTED        = AuthResponseStatus{Code: 5001, Msg: "Not Implemented"}
	PARAMETER_INCOMPLETE   = AuthResponseStatus{Code: 5002, Msg: "Parameter incomplete"}
	UNSUPPORTED            = AuthResponseStatus{Code: 5003, Msg: "Unsupported operation"}
	NO_AUTH_SOURCE         = AuthResponseStatus{Code: 5004, Msg: "AuthDefaultSource cannot be null"}
	UNIDENTIFIED_PLATFORM  = AuthResponseStatus{Code: 5005, Msg: "Unidentified platform"}
	ILLEGAL_REDIRECT_URI   = AuthResponseStatus{Code: 5006, Msg: "Illegal redirect uri"}
	ILLEGAL_REQUEST        = AuthResponseStatus{Code: 5007, Msg: "Illegal request"}
	ILLEGAL_CODE           = AuthResponseStatus{Code: 5008, Msg: "Illegal code"}
	ILLEGAL_STATUS         = AuthResponseStatus{Code: 5009, Msg: "Illegal state"}
	REQUIRED_REFRESH_TOKEN = AuthResponseStatus{Code: 5010, Msg: "The refresh token is required; it must not be null"}
	ILLEGAL_TOKEN          = AuthResponseStatus{Code: 5011, Msg: "Invalid token"}
)
