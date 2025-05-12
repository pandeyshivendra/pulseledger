package enums

type AuditAction string

const (
	ActionCreate AuditAction = "CREATE"
	ActionUpdate AuditAction = "UPDATE"
	ActionDelete AuditAction = "DELETE"
)
