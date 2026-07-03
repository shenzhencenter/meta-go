package objects

type TargetingSubscriberUniverse struct {
	MessagingCustomerBaseForWhatsapp *IDName `json:"messaging_customer_base_for_whatsapp,omitempty"`
	MessengerSubscriberPool          *IDName `json:"messenger_subscriber_pool,omitempty"`
	MessengerSubscriberSource        *IDName `json:"messenger_subscriber_source,omitempty"`
	WhatsappSubscriberSource         *IDName `json:"whatsapp_subscriber_source,omitempty"`
}

var TargetingSubscriberUniverseFields = struct {
	MessagingCustomerBaseForWhatsapp string
	MessengerSubscriberPool          string
	MessengerSubscriberSource        string
	WhatsappSubscriberSource         string
}{
	MessagingCustomerBaseForWhatsapp: "messaging_customer_base_for_whatsapp",
	MessengerSubscriberPool:          "messenger_subscriber_pool",
	MessengerSubscriberSource:        "messenger_subscriber_source",
	WhatsappSubscriberSource:         "whatsapp_subscriber_source",
}
