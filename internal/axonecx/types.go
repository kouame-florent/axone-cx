package axonecx

type TicketStatus string

const (
	TICKET_STATUS_NEW     TicketStatus = "NEW"     //after creation
	TICKET_STATUS_OPEN    TicketStatus = "OPEN"    //has been evaluated and is assigned
	TICKET_STATUS_PENDING TicketStatus = "PENDING" //need more informations from end user
	TICKET_STATUS_SOLVED  TicketStatus = "SOLVED"  //issue no longer exists, or the work has been completed
	TICKET_STATUS_CLOSED  TicketStatus = "CLOSED"  //resolved and a sufficient amount of time has passed (1 week)
)

type TicketType string

const (
	TICKET_TYPE_QUESTION TicketType = "QUESTION"
	TICKET_TYPE_PROBLEM  TicketType = "PROBLEM"
	TICKET_TYPE_TASK     TicketType = "TASK"
)
