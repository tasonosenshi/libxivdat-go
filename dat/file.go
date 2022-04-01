package dat

import "io/fs"

const (
	/// Header size in bytes.
	HEADER_SIZE uint32 = 0x11

	/// Offset of the max_size header value from the actual file size on disk.
	/// This value should be added to the `max_size` in the header to produce size of the file on disk.
	/// Files have a null-padded "footer" of 15 bytes that cannot be omitted, as well as the 17 byte header.
	MAX_SIZE_OFFSET uint32 = 32

	/// Index of the `file_type` header record.
	INDEX_FILE_TYPE uintptr = 0x00

	/// Index of the `max_size` header record.
	INDEX_MAX_SIZE uintptr = 0x04

	/// Index of the `content_size` header record.
	INDEX_CONTENT_SIZE uintptr = 0x08
)

/// A reference to an open DAT file on the system. This emulates the standard lib
/// [`std::fs::File`] but provides additional DAT-specific functionality.
///
/// Reads and writes to DAT files are performed only on the data contents of the file.
/// XOR masks are automatically applied as necessary.
///
/// # Examples
/// ```rust
/// use libxivdat::dat_file::File;
/// use libxivdat::dat_type::Type;
/// use std::io::Read;
///
/// let mut dat_file = match File::open("./resources/TEST_XOR.DAT") {
///     Ok(dat_file) => dat_file,
///     Err(_) => panic!("Something broke!")
/// };
///
/// match dat_file.file_type() {
///     Type::Macro => {
///         let mut macro_bytes = vec![0u8; dat_file.content_size() as usize - 1];
///         match dat_file.read(&mut macro_bytes) {
///             Ok(count) => println!("Read {} bytes.", count),
///             Err(_) => panic!("Reading broke!")
///         }
///     },
///     _ => panic!("Not a macro file!")
/// };
/// ```
type File struct {
	/// Size in bytes of the readable content of the DAT file. This size includes a trailing null byte.
	/// The size of readable content is 1 less than this value.
	content_size uint32

	/// Type of the file. This will be inferred from the header when converting directly from a `File`.
	file_type Type

	/// A single byte that marks the end of the header. This is `0xFF` for most DAT files, but occasionally varies.
	/// The purpose of this byte is unknown.
	header_end_byte uint8

	/// Maximum allowed size of the content in bytes. The writeable size is 1 byte less than this value.
	/// Excess available space not used by content is null padded.
	///
	/// Altering this value from the defaults provided for each file type may
	/// produce undefined behavior in the game client.
	max_size uint32

	/// The underlying File handle.
	raw_file fs.File
}
