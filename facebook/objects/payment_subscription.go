package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PaymentSubscription struct {
	Amount             *string               `json:"amount,omitempty"`
	AppParamData       *string               `json:"app_param_data,omitempty"`
	Application        *Application          `json:"application,omitempty"`
	BillingPeriod      *string               `json:"billing_period,omitempty"`
	CanceledReason     *string               `json:"canceled_reason,omitempty"`
	CreatedTime        *time.Time            `json:"created_time,omitempty"`
	Currency           *string               `json:"currency,omitempty"`
	ID                 *core.ID              `json:"id,omitempty"`
	LastPayment        *PaymentEnginePayment `json:"last_payment,omitempty"`
	NextBillTime       *time.Time            `json:"next_bill_time,omitempty"`
	NextPeriodAmount   *string               `json:"next_period_amount,omitempty"`
	NextPeriodCurrency *string               `json:"next_period_currency,omitempty"`
	NextPeriodProduct  *string               `json:"next_period_product,omitempty"`
	PaymentStatus      *string               `json:"payment_status,omitempty"`
	PendingCancel      *bool                 `json:"pending_cancel,omitempty"`
	PeriodStartTime    *time.Time            `json:"period_start_time,omitempty"`
	Product            *string               `json:"product,omitempty"`
	Status             *string               `json:"status,omitempty"`
	Test               *uint64               `json:"test,omitempty"`
	TrialAmount        *string               `json:"trial_amount,omitempty"`
	TrialCurrency      *string               `json:"trial_currency,omitempty"`
	TrialExpiryTime    *time.Time            `json:"trial_expiry_time,omitempty"`
	UpdatedTime        *time.Time            `json:"updated_time,omitempty"`
	User               *User                 `json:"user,omitempty"`
}
