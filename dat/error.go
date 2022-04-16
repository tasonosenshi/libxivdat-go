package dat

/// Wrapper error for any error related to processing a binary DAT file.
type Error interface {
	error
}

type ErrorType int

const (
	/// Attempted to read a byte stream as UTF-8 text when it didn't contain
	/// valid UTF-8.
	BadEncoding ErrorType = iota

	/// The header data is incorrect. The file is probably not a binary DAT file,
	/// but may be a plaintext DAT.
	BadHeader

	/// Data provided exceeds the maximum length specified in the header or the
	/// maximum possible length.
	Overflow

	/// Data provided is shorter than the content_size specified in the header or
	/// the minimum possible length.
	Underflow

	/// Unexpectedly hit the EOF when attempting to read a block of data.
	EndOfFile

	/// Wrapper for various `std::io::Error` errors. Represents an error reading or writing a
	/// file on disk.
	FileIO

	/// Attempted to use a type-specific function on the incorrect [`DATType`](crate::dat_type::DATType)
	IncorrectType

	/// Invalid input for a function
	InvalidInput
)
