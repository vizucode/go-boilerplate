package models

import (
	"time"

	"github.com/google/uuid"
)

type TbTransaction struct {
	Id                    uuid.UUID `gorm:"column:id"`
	TenantId              uuid.UUID `gorm:"column:tenant_id"`
	UserId                uuid.UUID `gorm:"column:user_id"`
	TrxCode               string    `gorm:"column:trx_code"`
	BookingType           string    `gorm:"column:booking_type"`
	FirstName             string    `gorm:"column:first_name"`
	LastName              string    `gorm:"column:last_name"`
	PhoneNumber           string    `gorm:"column:phone_number"`
	UrgentNumber          string    `gorm:"column:phone_number"`
	Email                 string    `gorm:"column:email"`
	ProviderName          string    `gorm:"column:provider_name"`
	PaymentCode           string    `gorm:"column:payment_code"`
	PaymentChannel        string    `gorm:"column:payment_channel"`
	PaymentType           string    `gorm:"column:payment_type"`
	Status                string    `gorm:"column:status"`
	AdminFee              float64   `gorm:"column:admin_fee"`
	PaymentFee            float64   `gorm:"column:payment_fee"`
	TotalOrderPrice       float64   `gorm:"column:total_order_price"`
	PriceTax              float64   `gorm:"column:price_tax"`
	PromoCode             string    `gorm:"column:promo_code"`
	DiscountTotalPrice    float64   `gorm:"column:discount_total_price"`
	GrandTotalPrice       float64   `gorm:"column:grand_total_price"`
	QrUrl                 string    `gorm:"column:qr_url"`
	QrString              string    `gorm:"column:qr_string"`
	PaymentWebRedirect    string    `gorm:"column:payment_web_redirect"`
	PaymentMobileRedirect string    `gorm:"column:payment_mobile_redirect"`
	PaymentProviderId     string    `gorm:"column:payment_provider_id"`
	PaymentProviderName   string    `gorm:"column:payment_provider_name"`
	ExpiredAt             time.Time `gorm:"column:expired_at"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at"`
	BookingDate           time.Time `gorm:"column:booking_date"`
	Reason                string    `gorm:"column:reason"`

	Ntsa       float64 `gorm:"column:ntsa"`
	Commission float64 `gorm:"column:commision"`
	TotalFare  float64 `gorm:"column:totalfare"`

	FlightTransaction  []TbFlightTransaction `gorm:"foreignKey:TransactionId;reference:Id;constraint:OnUpdaye:CASCADE,OnDelete:CASCADE"`
	TbHotelTransaction []TbHotelTransaction  `gorm:"foreignKey:TransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Passenger          []TbFlightPassengers  `gorm:"foreignKey:TransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Insurance          []TbFlightInsurance   `gorm:"foreignKey:TransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TbHotelTransaction struct {
	Id                 uuid.UUID `gorm:"column:id;primaryKey"`
	TransactionId      uuid.UUID `gorm:"column:transaction_id"`
	SearchId           string    `gorm:"column:search_id"`
	BreakfastStatus    bool      `gorm:"column:breakfast_status"`
	BookingCode        string    `gorm:"column:booking_code"`
	VoucherCode        string    `gorm:"column:voucher_code"`
	HotelId            string    `gorm:"column:hotel_id"`
	CheckInDate        time.Time `gorm:"column:check_in_date"`
	CheckOutDate       time.Time `gorm:"column:check_out_date"`
	CountryName        string    `gorm:"column:country_name"`
	CityName           string    `gorm:"column:city_name"`
	HotelName          string    `gorm:"column:hotel_name"`
	HotelAddress       string    `gorm:"column:hotel_address"`
	TotalRoom          int       `gorm:"column:total_room"`
	AdultCount         int       `gorm:"column:adult_count"`
	ChildCount         int       `gorm:"column:child_count"`
	TotalNight         int       `gorm:"column:total_night"`
	Status             string    `gorm:"column:status"`
	Reason             string    `gorm:"column:reason"`
	ThumbnailImage     string    `gorm:"column:thumbnail_image"`
	HotelNetAgentPrice float64   `gorm:"column:hotel_net_agent_price"`
	HotelNightPrice    float64   `gorm:"column:hotel_night_price"`
	RoomName           string    `gorm:"column:room_name"`
	ExtraRequest       string    `gorm:"column:extra_request"`
	NearestDestination string    `gorm:"column:nearest_destination"`
	PaxContactName     string    `gorm:"column:pax_contact_name"`
	PaxContactEmail    string    `gorm:"column:pax_contact_email"`
	PaxContactPhone    string    `gorm:"pax_contact_phone"`
	UpdatedAt          time.Time `gorm:"updated_at"`
	CreatedAt          time.Time `gorm:"created_at"`
}

type TbFlightTransaction struct {
	Id                      uuid.UUID `gorm:"column:id"`
	TransactionId           uuid.UUID `gorm:"column:transaction_id"`
	TotalFlightHour         int       `gorm:"column:total_flight_hour"`
	TotalFlightMinute       int       `gorm:"column:total_flight_minute"`
	BookingDate             time.Time `gorm:"column:booking_date"`
	PnrCode                 string    `gorm:"column:pnr_code"`
	BookingCode             string    `gorm:"column:booking_code"`
	InsuranceNumber         string    `gorm:"column:insurance_number"`
	Status                  string    `gorm:"column:status"`
	ConnectingTroughfare    bool      `gorm:"column:connecting_troughfare"`
	FlightNetAgentPrice     float64   `gorm:"column:flight_net_agent_price"`
	FlightPublishAgentPrice float64   `gorm:"column:flight_publish_agent_price"`
	AirlineCode             string    `gorm:"column:airline_code"`
	Adult                   int       `gorm:"column:adult"`
	Child                   int       `gorm:"column:child"`
	Infant                  int       `gorm:"column:infant"`
	DepartAirport           string    `gorm:"column:depart_airport"`
	ArrivalAirport          string    `gorm:"column:arrival_airport"`
	IsInternational         bool      `gorm:"column:is_international"`
	DepartDateTime          time.Time `gorm:"column:depart_datetime"`
	ArrivalDateTime         time.Time `gorm:"column:arrival_datetime"`
	FlightType              string    `gorm:"column:flight_type"`
	CreatedAt               time.Time `gorm:"column:created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at"`

	FlightJourney []TbFlightJourneys `gorm:"foreignKey:FlightTransactionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TbFlightAddons struct {
	Id                 uuid.UUID `gorm:"column:id"`
	FlightPassengersId uuid.UUID `gorm:"column:flight_passengers_id"`
	Type               string    `gorm:"column:type"`
	Value              string    `gorm:"column:value"`
	NetAgentPrice      float64   `gorm:"column:net_agent_price"`
	PublishPrice       float64   `gorm:"column:publish_price"`
	TotalPrice         float64   `gorm:"column:total_price"`
}

type TbFlightInsurance struct {
	Id            uuid.UUID `gorm:"column:id"`
	TransactionId uuid.UUID `gorm:"column:transaction_id"`
	InsuranceId   uuid.UUID `gorm:"column:insurance_id"`
	Title         string    `gorm:"column:title"`
	Currency      string    `gorm:"column:currency"`
	NetAgentPrice float64   `gorm:"column:net_agent_price"`
	PublishPrice  float64   `gorm:"column:publish_price"`
	TotalPrice    float64   `gorm:"column:total_price"`
}

type TbFlightPassengers struct {
	Id                     uuid.UUID `gorm:"column:id"`
	TransactionId          string    `gorm:"column:transaction_id"`
	PassengerNo            string    `gorm:"column:passenger_no"`
	TicketNumberDeparture  string    `gorm:"column:ticket_number_departure"`
	TicketNumberReturn     string    `gorm:"column:ticket_number_return"`
	Title                  string    `gorm:"column:title"`
	FirstName              string    `gorm:"column:first_name"`
	LastName               string    `gorm:"column:last_name"`
	AdultAssoc             int       `gorm:"column:adult_assoc"`
	NationalIdNumber       string    `gorm:"column:national_id_number"`
	PassportNumber         string    `gorm:"column:passport_number"`
	PassportIssuedDate     time.Time `gorm:"column:passport_issued_date"`
	PassportIssuedCountry  string    `gorm:"column:passport_issued_country"`
	PassportExpirationDate time.Time `gorm:"passport_expiration_date"`
	PassportNationality    string    `gorm:"column:passport_nationality"`
	DateOfBirth            time.Time `gorm:"column:date_of_birth"`
	IsAdultAssoc           int       `gorm:"column:is_adult_assoc"`
	PhoneNumber            string    `gorm:"column:phone_number"`
	Email                  string    `gorm:"column:email"`
	Type                   string    `gorm:"column:type"`

	AddOns []TbFlightAddons `gorm:"foreignKey:FlightPassengersId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TbFlightJourneys struct {
	Id                  uuid.UUID `gorm:"column:id"`
	FlightNo            string    `gorm:"column:flight_no"`
	Meal                bool      `gorm:"column:meal"`
	DurationHour        int       `gorm:"column:duration_hour"`
	DurationMinute      int       `gorm:"column:duration_minute"`
	FlightTransactionId string    `gorm:"column:flight_transaction_id"`
	Departure           string    `gorm:"column:departure"`
	Destination         string    `gorm:"column:destination"`
	StopOver            int       `gorm:"column:stopover"`
	DepartureDateTime   time.Time `gorm:"column:departure_datetime"`
	ArrivalDateTime     time.Time `gorm:"column:arrival_datetime"`
	JourneyType         string    `gorm:"column:journey_type"`
	ClassCode           string    `gorm:"column:class_code"`
	ClassName           string    `gorm:"column:class_name"`
	IataCode            string    `gorm:"column:iata_code"`
	ChangeDay           int       `gorm:"column:change_day"`
	FreeBaggageKilo     int       `gorm:"column:free_baggage_kilo"`
}

type TbLogTransaction struct {
	Id            uuid.UUID `gorm:"column:id"`
	TransactionId uuid.UUID `gorm:"column:transaction_id"`
	ProviderName  string    `gorm:"column:provider_name"`
	RequestBody   string    `gorm:"column:request_body"`
	Response      string    `gorm:"column:response"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
