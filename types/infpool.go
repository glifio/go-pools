package types

type RejectionReason string

const (
	RejectionReasonNone           RejectionReason = "none"
	RejectionReasonPayUP          RejectionReason = "Agent is behind on payments"
	RejectionReasonCap            RejectionReason = "Agent reached its cap"
	RejectionReasonGCRED          RejectionReason = "Agent has a GCRED score below the threshold"
	RejectionReasonZeroCollateral RejectionReason = "Agent has zero collateral"
	RejectionReasonLTV            RejectionReason = "Agent fails the LTV check"
	RejectionReasonDTE            RejectionReason = "Agent fails the DTE check"
	RejectionReasonDTI            RejectionReason = "Agent fails the DTI check"
)
