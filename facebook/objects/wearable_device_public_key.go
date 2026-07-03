package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WearableDevicePublicKey struct {
	Base64EncodedPublicKey *string    `json:"base64_encoded_public_key,omitempty"`
	CreationTimeOnDevice   *core.Time `json:"creation_time_on_device,omitempty"`
	DeviceUUID             *string    `json:"device_uuid,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	KeyType                *string    `json:"key_type,omitempty"`
	OwnerID                *core.ID   `json:"owner_id,omitempty"`
	ProductUseCase         *string    `json:"product_use_case,omitempty"`
	Version                *string    `json:"version,omitempty"`
}

var WearableDevicePublicKeyFields = struct {
	Base64EncodedPublicKey string
	CreationTimeOnDevice   string
	DeviceUUID             string
	ID                     string
	KeyType                string
	OwnerID                string
	ProductUseCase         string
	Version                string
}{
	Base64EncodedPublicKey: "base64_encoded_public_key",
	CreationTimeOnDevice:   "creation_time_on_device",
	DeviceUUID:             "device_uuid",
	ID:                     "id",
	KeyType:                "key_type",
	OwnerID:                "owner_id",
	ProductUseCase:         "product_use_case",
	Version:                "version",
}
