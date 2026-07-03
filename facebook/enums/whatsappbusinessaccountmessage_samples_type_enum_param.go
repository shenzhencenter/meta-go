package enums

type WhatsappbusinessaccountmessageSamplesTypeEnumParam string

const (
	WhatsappbusinessaccountmessageSamplesTypeEnumParamInteractive WhatsappbusinessaccountmessageSamplesTypeEnumParam = "INTERACTIVE"
	WhatsappbusinessaccountmessageSamplesTypeEnumParamText        WhatsappbusinessaccountmessageSamplesTypeEnumParam = "TEXT"
)

func (value WhatsappbusinessaccountmessageSamplesTypeEnumParam) String() string {
	return string(value)
}
