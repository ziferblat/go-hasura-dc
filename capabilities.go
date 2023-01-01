package dc

import "context"

type CapabilitiesService interface {
	Get(ctx context.Context, conf any) (*CapabilitiesResponse, error)
}

type CapabilitiesResponse struct {
	Capabilities Capabilities `json:"capabilities"`

	ConfigSchemas ConfigSchemaResponse `json:"config_schemas"`

	DisplayName string `json:"display_name,omitempty"`

	ReleaseName string `json:"release_name,omitempty"`
}

type Capabilities struct{}

type ConfigSchemaResponse struct{}
