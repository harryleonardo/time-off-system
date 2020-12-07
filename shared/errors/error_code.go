package errors

type ERRCODE string
type CHANNEL string

const (
	ERRCODEVPGEN001 = "VP-GEN-001" // - internal server error
	ERRCODEVPGEN002 = "VP-GEN-002" // - failed to binding request into struct
	ERRCODEVPGEN003 = "VP-GEN-003" // - validation error

	ERRCODEAUTH001 = "ERR-AUTH-001" // - failed to get credential

	ERRCODEEMP001 = "ERR-EMP-001" // - failed to get insert into profile table
	ERRCODEEMP002 = "ERR-EMP-002" // - failed to get insert into employee table
	ERRCODEEMP003 = "ERR-EMP-003" // - failed to get the detail based on employee_id
	ERRCODEEMP004 = "ERR-EMP-004" // - failed to get the list of employee based on company_id

	ERRCODEOFF001 = "ERR-OFF-001" // - quota less than request total
	ERRCODEOFF002 = "ERR-OFF-002" // - failed to insert into time-off table;
	ERRCODEOFF003 = "ERR-OFF-003" // - failed to get time-off detail;
	ERRCODEOFF004 = "ERR-OFF-004" // - failed to update employee info in approved action;
	ERRCODEOFF005 = "ERR-OFF-005" // - failed to update time off data;
	ERRCODEOFF006 = "ERR-OFF-006" // - failed to get list of time off data;
)
