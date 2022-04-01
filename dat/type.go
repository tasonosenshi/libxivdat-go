package dat

import "errors"

/// Enumeration of known FFXIV DAT file types.
/// The value of each element represents the first 4 header bytes as a little-endian i32.
/// These bytes are known static values that differentiate file types.
///
/// File types may be referenced using a human readable descriptor -- `Type.GoldSaucer` --
/// or the filename used by FFXIV -- `Type.GS`. These methods are interchangable and considered
/// equivalent. `Type.GoldSaucer == Type.GS`.
type Type uint32

const (
	/// ACQ.DAT (Acquaintances?)
	RecentTells Type = 0x00640006
	ACQ

	/// GEARSET.DAT
	Gearset Type = 0x006b0005
	GEARSET

	/// GS.DAT
	GoldSaucer Type = 0x0067000A
	GS

	/// HOTBAR.DAT
	Hotbar Type = 0x00040002
	HOTBAR

	/// ITEMFDR.DAT
	ItemFinder Type = 0x00CA0008
	ITEMFDR

	/// ITEMODR.DAT
	ItemOrder Type = 0x00670007
	ITEMODR

	/// KEYBIND.DAT
	Keybind Type = 0x00650003
	KEYBIND

	/// LOGFLTR.DAT
	LogFilter Type = 0x00030004
	LOGFLTR

	/// MACRO.DAT (Character) & MACROSYS.DAT (Global)
	Macro Type = 0x00020001
	MACRO
	MACROSYS

	/// UISAVE.DAT
	UISave Type = 0x00010009
	UISAVE

	Unknown Type = 0
)

var Types []Type = []Type{
	RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, LogFilter, Macro, UISave,
}

func (dtype Type) From(header uint32) Type {
	for _, t := range Types {
		if dtype == t {
			return t
		}
	}

	return Unknown
}

func GetMaskForType(fileType Type) (uint8, error) {
	switch fileType {
	case RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, Macro:
		return 0x73, nil
	case Hotbar, UISave:
		return 0x31, nil
	case LogFilter:
		return 0x00, nil
	default:
		return 0, errors.New("unknown file type")
	}
}

func GetDefaultEndByteForType(fileType Type) (uint8, error) {
	switch fileType {
	case RecentTells, Gearset, GoldSaucer, ItemFinder, ItemOrder, Keybind, Macro:
		return 0xFF, nil
	case Hotbar:
		return 0x31, nil
	case UISave:
		return 0x00, nil
	case LogFilter:
		return 0x21, nil
	default:
		return 0, errors.New("unknown file type")
	}
}

func GetDefaultMaxSizeForType(fileType Type) (uint32, error) {
	switch fileType {
	case RecentTells:
		return 2048, nil
	case Gearset:
		return 44849, nil
	case GoldSaucer:
		return 649, nil
	case Hotbar:
		return 204800, nil
	case ItemFinder:
		return 14030, nil
	case ItemOrder:
		return 15193, nil
	case Keybind:
		return 20480, nil
	case LogFilter:
		return 2048, nil
	case Macro:
		return 286720, nil
	case UISave:
		return 64512, nil
	default:
		return 0, errors.New("unknown file type")
	}
}
