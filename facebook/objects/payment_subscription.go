package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PaymentSubscription struct {
	Amount             *string               `json:"amount,omitempty"`
	AppParamData       *string               `json:"app_param_data,omitempty"`
	Application        *Application          `json:"application,omitempty"`
	BillingPeriod      *string               `json:"billing_period,omitempty"`
	CanceledReason     *string               `json:"canceled_reason,omitempty"`
	CreatedTime        *core.Time            `json:"created_time,omitempty"`
	Currency           *string               `json:"currency,omitempty"`
	ID                 *core.ID              `json:"id,omitempty"`
	LastPayment        *PaymentEnginePayment `json:"last_payment,omitempty"`
	NextBillTime       *core.Time            `json:"next_bill_time,omitempty"`
	NextPeriodAmount   *string               `json:"next_period_amount,omitempty"`
	NextPeriodCurrency *string               `json:"next_period_currency,omitempty"`
	NextPeriodProduct  *string               `json:"next_period_product,omitempty"`
	PaymentStatus      *string               `json:"payment_status,omitempty"`
	PendingCancel      *bool                 `json:"pending_cancel,omitempty"`
	PeriodStartTime    *core.Time            `json:"period_start_time,omitempty"`
	Product            *string               `json:"product,omitempty"`
	Status             *string               `json:"status,omitempty"`
	Test               *uint64               `json:"test,omitempty"`
	TrialAmount        *string               `json:"trial_amount,omitempty"`
	TrialCurrency      *string               `json:"trial_currency,omitempty"`
	TrialExpiryTime    *core.Time            `json:"trial_expiry_time,omitempty"`
	UpdatedTime        *core.Time            `json:"updated_time,omitempty"`
	User               *User                 `json:"user,omitempty"`
}

var PaymentSubscriptionFields = struct {
	Amount             string
	AppParamData       string
	Application        string
	BillingPeriod      string
	CanceledReason     string
	CreatedTime        string
	Currency           string
	ID                 string
	LastPayment        string
	NextBillTime       string
	NextPeriodAmount   string
	NextPeriodCurrency string
	NextPeriodProduct  string
	PaymentStatus      string
	PendingCancel      string
	PeriodStartTime    string
	Product            string
	Status             string
	Test               string
	TrialAmount        string
	TrialCurrency      string
	TrialExpiryTime    string
	UpdatedTime        string
	User               string
}{
	Amount:             "amount",
	AppParamData:       "app_param_data",
	Application:        "application",
	BillingPeriod:      "billing_period",
	CanceledReason:     "canceled_reason",
	CreatedTime:        "created_time",
	Currency:           "currency",
	ID:                 "id",
	LastPayment:        "last_payment",
	NextBillTime:       "next_bill_time",
	NextPeriodAmount:   "next_period_amount",
	NextPeriodCurrency: "next_period_currency",
	NextPeriodProduct:  "next_period_product",
	PaymentStatus:      "payment_status",
	PendingCancel:      "pending_cancel",
	PeriodStartTime:    "period_start_time",
	Product:            "product",
	Status:             "status",
	Test:               "test",
	TrialAmount:        "trial_amount",
	TrialCurrency:      "trial_currency",
	TrialExpiryTime:    "trial_expiry_time",
	UpdatedTime:        "updated_time",
	User:               "user",
}
