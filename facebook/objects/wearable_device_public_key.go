package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type WearableDevicePublicKey struct {
	Base64EncodedPublicKey *string    `json:"base64_encoded_public_key,omitempty"`
	CreationTimeOnDevice   *time.Time `json:"creation_time_on_device,omitempty"`
	DeviceUUID             *string    `json:"device_uuid,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	KeyType                *string    `json:"key_type,omitempty"`
	OwnerID                *core.ID   `json:"owner_id,omitempty"`
	ProductUseCase         *string    `json:"product_use_case,omitempty"`
	Version                *string    `json:"version,omitempty"`
}
