package libxivdatgo

/// Enumeration of known FFXIV DAT file types.
/// The value of each element represents the first 4 header bytes as a little-endian i32.
/// These bytes are known static values that differentiate file types.
///
/// File types may be referenced using a human readable descriptor -- `DATType.GoldSaucer` --
/// or the filename used by FFXIV -- `DATType.GS`. These methods are interchangable and considered
/// equivalent. `DATType.GoldSaucer == DATType.GS`.
type DATType uint32

const (
	/// ACQ.DAT (Acquaintances?)
	RecentTells DATType = 0x00640006
	ACQ

	/// GEARSET.DAT
	Gearset DATType = 0x006b0005
	GEARSET

	/// GS.DAT
	GoldSaucer DATType = 0x0067000A
	GS

	/// HOTBAR.DAT
	Hotbar DATType = 0x00040002
	HOTBAR

	/// ITEMFDR.DAT
	ItemFinder DATType = 0x00CA0008
	ITEMFDR

	/// ITEMODR.DAT
	ItemOrder DATType = 0x00670007
	ITEMODR

	/// KEYBIND.DAT
	Keybind DATType = 0x00650003
	KEYBIND

	/// LOGFLTR.DAT
	LogFilter DATType = 0x00030004
	LOGFLTR

	/// MACRO.DAT (Character) & MACROSYS.DAT (Global)
	Macro DATType = 0x00020001
	MACRO
	MACROSYS

	/// UISAVE.DAT
	UISave DATType = 0x00010009
	UISAVE

	Unknown DATType = 0
)

var datTypes []DATType = []DATType{
	RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, LogFilter, Macro, UISave,
}

func (dtype DATType) From(header uint32) DATType {
	for _, datType := range datTypes {
		if dtype == datType {
			return datType
		}
	}

	return Unknown
}

func getMaskForType(fileType DATType) uint8 {
	switch fileType {
	case RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, Macro:
		return 0x73
	case Hotbar, UISave:
		return 0x31
	case LogFilter:
		return 0x00
	default:
		return 0
	}
}

func getDefaultEndByteForType(fileType DATType) uint8 {
	switch fileType {
	case RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, Macro:
		return 0xFF
	case Hotbar:
		return 0x31
	case UISave:
		return 0x00
	case LogFilter:
		return 0x21
	default:
		return 0
	}
}

func getDefaultMaxSizeForType(fileType DATType) uint32 {
	switch fileType {
	case RecentTells:
		return 2048
	case Gearset:
		return 44849
	case GoldSaucer:
		return 649
	case Hotbar:
		return 204800
	case ItemFinder:
		return 14030
	case ItemOrder:
		return 15193
	case Keybind:
		return 20480
	case LogFilter:
		return 2048
	case Macro:
		return 286720
	case UISave:
		return 64512
	default:
		return 0
	}
}
