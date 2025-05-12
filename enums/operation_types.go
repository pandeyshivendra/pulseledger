package enums

type OperationType uint8

func (o OperationType) Description() string {
	switch o {
	case NormalPurchase:
		return "Normal Purchase"
	case InstalmentPurchase:
		return "Purchase with Installments"
	case Withdrawal:
		return "Withdrawal"
	case CreditVoucher:
		return "Credit Voucher"
	default:
		return "Unknown Operation"
	}
}

const (
	NormalPurchase     OperationType = 1
	InstalmentPurchase OperationType = 2
	Withdrawal         OperationType = 3
	CreditVoucher      OperationType = 4
)

func AllOperationTypes() []OperationType {
	return []OperationType{
		NormalPurchase,
		InstalmentPurchase,
		Withdrawal,
		CreditVoucher,
	}
}
