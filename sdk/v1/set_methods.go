package v1

import "go-ws/sdk/types"

func (s SDKImpl) SetErrorInvoke(function types.InvokeFunction) {
	errorMessageInvoke = function
}
