package api

import (
	t "go-ws/sdk/types"
	v1 "go-ws/sdk/v1"
)

const version = t.V1

func NewSDK(url string) (sdk t.SDK) {
	sdk = getSDK(version, url)
	return sdk
}

func getSDK(version t.Version, url string) (sdk t.SDK) {
	switch version {
	case t.V1:
		// Bind interface to implementation matching the selected API version.
		sdk = v1.NewOemlSdkV1(url)
	default:
		sdk = v1.NewOemlSdkV1(url)
	}
	return sdk
}
