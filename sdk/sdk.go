package api

import (
	t "go-ws/sdk/types"
	v1 "go-ws/sdk/v1"
)

const apiVersion = t.V1

func NewSDK() (sdk t.SDK) {
	sdk = getSDK(apiVersion)
	return sdk
}

func getSDK(version t.Version) (sdk t.SDK) {
	switch version {
	case t.V1:
		// Bind interface to implementation matching the selected API version.
		sdk = v1.NewOemlSdkV1()
	default:
		sdk = v1.NewOemlSdkV1()
	}
	return sdk
}
