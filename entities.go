package mundipagg

import "time"

type Boleto struct {
	Bank           Bank              `json:"bank,omitempty"`
	Instructions   string            `json:"instructions,omitempty"` // max: 256 characters
	DueAt          *time.Time        `json:"due_at,omitempty"`
	NossoNumero    string            `json:"nosso_numero,omitempty"`
	Type           BoletoType        `json:"type,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	DocumentNumber string            `json:"document_number,omitempty"`
}

type Bank string

const (
	BancoDoBrasil Bank = "001"
	Santander     Bank = "033"
	Caixa         Bank = "104"
	Bradesco      Bank = "237"
	Itau          Bank = "341"
	Citibank      Bank = "745"
)

type BoletoType string

const (
	Duplicata BoletoType = "DM"
	Proposta  BoletoType = "BDP"
)

type CreditCard struct {
	Installments         int           `json:"installments,omitempty"`
	StatementDescriptor  string        `json:"statement_descriptor,omitempty"` // ax 22 characters
	OperationType        OperationType `json:"Operation_type,omitempty"`
	CardCredit           *CreditCard   `json:"credit_card,omitempty"`
	CardID               string        `json:"card_id,omitempty"`
	CardToken            string        `json:"card_token,omitempty"`
	Recurrence           bool          `json:"recurrence,omitempty"`
	Metadata             map[string]string
	ExtendedLimitEnabled bool         `json:"extended_limit_enabled,omitempty"`
	ExtendedLimitCode    string       `json:"extended_limit_code,omitempty"`
	MerchantCategoryCode int          `json:"merchant_id,omitempty"`
	Authentication       *interface{} // TODO ----------------------
	AutoRecovery         bool         `json:"auto_recovery,omitempty"`
	Payload              *interface{} `json:"payload,omitempty"` // TODO ---------------
}

type OperationType string

const (
	AuthAndCapture OperationType = "auth_and_capture"
	Auth           OperationType = "auth_only"
	PreAuth        OperationType = "pre_auth"
)

type CreditCardToken struct {
	Token   string             `json:"token,omitempty"`
	Options *CreditCardOptions `json:"options,omitempty"`
}

type CreditCardOptions struct {
	VerifyCard bool `json:"verify_card,omitempty"`
}

type Customer struct {
	Name     string            `json:"name,omitempty"`
	Email    string            `json:"email,omitempty"`
	Code     string            `json:"code,omitempty"`
	Document string            `json:"document,omitempty"`
	Type     DocumentType      `json:"type,omitempty"`
	Gender   string            `json:"gender,omitempty"`
	Address  *Address          `json:"address,omitempty"`
	Phones   *Phones           `json:"phones,omitempty"`
	Birthday *time.Time        `json:"birthday,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type DocumentType string

const (
	CPF  DocumentType = "individual"
	CNPJ DocumentType = "company"
)

type Address struct {
	ID        string     `json:"id,omitempty"`
	Line1     string     `json:"line_1,omitempty"` // número, Rua, Bairro - nesta ordem e separados por vírgula
	Line2     string     `json:"line_2,omitempty"` // informações complementares do endereço, tais como andar, apto, sala
	Zipcode   string     `json:"zip_code,omitempty"`
	City      string     `json:"city,omitempty"`
	State     string     `json:"state,omitempty"`
	Country   string     `json:"country,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Phones struct {
	HomePhone   *Phone `json:"home_phone,omitempty"`
	MobilePhone *Phone `json:"mobile_phone,omitempty"`
}

type Phone struct {
	CountryCode string `json:"country_code,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
	Number      string `json:"number,omitempty"`
}

type PriceSchema struct {
	Price         int             `json:"price,omitempty"`
	MinimumPrice  int             `json:"minimum_price,omitempty"`
	SchemaType    SchemaType      `json:"schema_type,omitempty"`
	Quantity      int             `json:"quantity,omitempty"`
	PriceBrackets *[]PriceBracket `json:"price_brackets,omitempty"`
}

type SchemaType string

const (
	Unit    SchemaType = "unit"
	Package SchemaType = "package"
	Volume  SchemaType = "volume"
	Tier    SchemaType = "tier"
)

type PriceBracket struct {
	StartedQuantity int64 `json:"start_quantity,omitempty"`
	EndQuantity     int64 `json:"end_quantity,omitempty"`
	OveragePrice    int64 `json:"overage_price,omitempty"`
	Price           int64 `json:"price,omitempty"`
}

type Subscription struct {
	Code                 string             `json:"code,omitempty"`
	PaymentMethod        Method             `json:"payment_method,omitempty"`
	Currency             Currency           `json:"currency,omitempty"`
	StartAt              *time.Time         `json:"start_at,omitempty"`
	MinimumPrice         int                `json:"minimum_price,omitempty"`
	Interval             Interval           `json:"interval,omitempty"`
	IntervalCount        int                `json:"interval_count,omitempty"`
	BillingType          BillingType        `json:"billing_type,omitempty"`
	BillingDay           int                `json:"billing_day,omitempty"`
	Description          string             `json:"description,omitempty"`
	Installments         int                `json:"installments,omitempty"`
	StatementDescriptor  string             `json:"statement_descriptor,omitempty"`
	CustomerID           string             `json:"customer_id,omitempty"`
	Customer             *Customer          `json:"customer,omitempty"`
	Discounts            *[]BillExtras      `json:"discount,omitempty"`
	Increments           *[]BillExtras      `json:"increments,omitempty"`
	Items                *[]Item            `json:"items,omitempty"`
	Setup                *Setup             `json:"setup,omitempty"`
	GatewayAffiliationID string             `json:"gateway_affiliation_id,omitempty"`
	BoletoDueDays        int                `json:"boleto_due_days,omitempty"`
	Card                 *SubscriptionCards `json:"card,omitempty"`
	Boleto               *Boleto            `json:"boleto,omitempty"`
	Metadata             map[string]string  `json:"metadata,omitempty"`
}

type Method string

const (
	MethodCreditCard Method = "credit_card"
	MethodDebit      Method = "debit_card"
	MethodBoleto     Method = "boleto"
)

type SubscriptionCards struct {
	CardID string `json:"card_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type Currency string

const (
	BRL Currency = "BRL"
	ARS Currency = "ARS"
	BOB Currency = "BOB"
	CLP Currency = "CLP"
	COP Currency = "COP"
	MXN Currency = "MXN"
	PYG Currency = "PYG"
	USD Currency = "USD"
	UYU Currency = "UYU"
	EUR Currency = "EUR"
)

type Interval string

const (
	Day   Interval = "day"
	Week  Interval = "week"
	Month Interval = "month"
	Year  Interval = "year"
)

type BillingType string

const (
	Prepaid  BillingType = "prepaid"
	PostPaid BillingType = "postpaid"
	ExactDay BillingType = "exact_day"
)

type BillExtras struct {
	ID           string       `json:"id,omitempty"`
	Cycles       int          `json:"cycles,omitempty"`
	Value        int          `json:"value,omitempty"`
	DiscountType DiscountType `json:"discount_type,omitempty"`
	ItemID       string       `json:"item_id,omitempty"`
	Status       StatusType   `json:"status,omitempty"`
	CreatedAt    *time.Time   `json:"created_at,omitempty"`
}

type DiscountType string

const (
	Flat       DiscountType = "flat"
	Percentage DiscountType = "percentage"
)

type StatusType string

const (
	Active   StatusType = "active"
	Deleted  StatusType = "deleted"
	Inactive StatusType = "inactive"
)

type Item struct {
	ID            string        `json:"id,omitempty"`
	Description   string        `json:"description,omitempty"`
	Cycles        int           `json:"cycles,omitempty"`
	Quantity      int           `json:"quantity,omitempty"`
	Status        StatusType    `json:"status,omitempty"`
	PricingSchema *PriceSchema  `json:"pricing_scheme,omitempty"`
	CreatedAt     *time.Time    `json:"created_at,omitempty"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
	DeletedAt     *time.Time    `json:"deleted_at,omitempty"`
	Discounts     *[]BillExtras `json:"discounts,omitempty"`
	Increments    *[]BillExtras `json:"increments,omitempty"`
	Name          string        `json:"name,omitempty"`
}

type Setup struct {
	Amount      int      `json:"amount"`
	Description string   `json:"description"`
	Payment     *Payment `json:"payment"`
}

type Payment struct {
	PaymentMethod        Method            `json:"payment_method,omitempty"`
	CreditCard           *CreditCard       `json:"credit_card,omitempty"`
	Voucher              interface{}       `json:"voucher,omitempty"`
	Boleto               *Boleto           `json:"boleto,omitempty"`
	BankTransfer         interface{}       `json:"bank_transfer,omitempty"`
	Checkout             interface{}       `json:"checkout,omitempty"`
	Cash                 interface{}       `json:"cash,omitempty"`
	Amount               int64             `json:"amount,omitempty"`
	Metadata             map[string]string `json:"metadata,omitempty"`
	GatewayAffiliationID string            `json:"gateway_affiliation_id,omitempty"`
}
