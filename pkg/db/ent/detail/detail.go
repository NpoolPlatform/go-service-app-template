// Code generated by ent, DO NOT EDIT.

package detail

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the detail type in the database.
	Label = "detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldIoType holds the string denoting the io_type field in the database.
	FieldIoType = "io_type"
	// FieldIoSubType holds the string denoting the io_sub_type field in the database.
	FieldIoSubType = "io_sub_type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldFromCoinTypeID holds the string denoting the from_coin_type_id field in the database.
	FieldFromCoinTypeID = "from_coin_type_id"
	// FieldCoinUsdCurrency holds the string denoting the coin_usd_currency field in the database.
	FieldCoinUsdCurrency = "coin_usd_currency"
	// FieldIoExtra holds the string denoting the io_extra field in the database.
	FieldIoExtra = "io_extra"
	// FieldFromOldID holds the string denoting the from_old_id field in the database.
	FieldFromOldID = "from_old_id"
	// Table holds the table name of the detail in the database.
	Table = "details"
)

// Columns holds all SQL columns for detail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldCoinTypeID,
	FieldIoType,
	FieldIoSubType,
	FieldAmount,
	FieldFromCoinTypeID,
	FieldCoinUsdCurrency,
	FieldIoExtra,
	FieldFromOldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/service-template/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultIoType holds the default value on creation for the "io_type" field.
	DefaultIoType string
	// DefaultIoSubType holds the default value on creation for the "io_sub_type" field.
	DefaultIoSubType string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount decimal.Decimal
	// DefaultFromCoinTypeID holds the default value on creation for the "from_coin_type_id" field.
	DefaultFromCoinTypeID func() uuid.UUID
	// DefaultCoinUsdCurrency holds the default value on creation for the "coin_usd_currency" field.
	DefaultCoinUsdCurrency decimal.Decimal
	// DefaultIoExtra holds the default value on creation for the "io_extra" field.
	DefaultIoExtra string
	// DefaultFromOldID holds the default value on creation for the "from_old_id" field.
	DefaultFromOldID func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
