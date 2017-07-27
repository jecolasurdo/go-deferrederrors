package injectionbehavior

// InjectionBehavior instructs the system on how to thread values through the chain of actions.
type InjectionBehavior int

const (
	// NotSpecified means that no behavior has been declared. This value is assumed to mean UsePrevious
	NotSpecified InjectionBehavior = iota

	// UsePrevious means that your intention is to inject the value supplied by the previous action in the chain.
	// If no previous action was present in the chain, nil is injected.
	// If a value is supplied in the Argument to the chain function, it is ignored.
	UsePrevious

	// OverridePrevious means that your intention is to inject the value supplied in the Argument.
	// If a previous value is presnet in the chain, it is ignored.
	OverridePrevious
)
