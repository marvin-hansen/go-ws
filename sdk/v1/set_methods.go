package v1

import "go-ws/sdk/types"

func (s SDKImpl) SetErrorInvoke(function types.InvokeFunction) {
	errorMessageInvoke = function
}

func (s SDKImpl) SetHeartBeatInvoke(function types.InvokeFunction) {
	heartBeatInvoke = function
}

func (s SDKImpl) SetReconnectInvoke(function types.InvokeFunction) {
}
