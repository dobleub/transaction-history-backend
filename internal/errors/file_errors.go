package errors

const (
	// ErrFileNotFound is returned when a file is not found
	ErrFileNotFound = "file not found"
	// ErrFileNotCSV is returned when a file is not a CSV file
	ErrFileNotCSV = "file is not a CSV file"
	// ErrFileNotJSON is returned when a file is not a JSON file
	ErrFileNotJSON = "file is not a JSON file"
	// ErrFileNotYAML is returned when a file is not a YAML file
	ErrFileNotYAML = "file is not a YAML file"
	// Transactions not found
	ErrTransactionsNotFound = "no transactions found"
)
