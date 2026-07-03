package enums

type ApplicationcodelessEventMappingsMutationMethodEnumParam string

const (
	ApplicationcodelessEventMappingsMutationMethodEnumParamAdd     ApplicationcodelessEventMappingsMutationMethodEnumParam = "ADD"
	ApplicationcodelessEventMappingsMutationMethodEnumParamDelete  ApplicationcodelessEventMappingsMutationMethodEnumParam = "DELETE"
	ApplicationcodelessEventMappingsMutationMethodEnumParamReplace ApplicationcodelessEventMappingsMutationMethodEnumParam = "REPLACE"
)

func (value ApplicationcodelessEventMappingsMutationMethodEnumParam) String() string {
	return string(value)
}
