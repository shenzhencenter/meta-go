package enums

type ApplicationcodelessEventMappingsPostMethodEnumParam string

const (
	ApplicationcodelessEventMappingsPostMethodEnumParamCodeless ApplicationcodelessEventMappingsPostMethodEnumParam = "CODELESS"
	ApplicationcodelessEventMappingsPostMethodEnumParamEymt     ApplicationcodelessEventMappingsPostMethodEnumParam = "EYMT"
)

func (value ApplicationcodelessEventMappingsPostMethodEnumParam) String() string {
	return string(value)
}
