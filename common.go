package gonanogui

type Cursor int

const (
	Arrow       Cursor = iota + 1 ///< The arrow cursor.
	IBeam                         ///< The I-beam cursor.
	Crosshair                     ///< The crosshair cursor.
	Hand                          ///< The hand cursor.
	HResize                       ///< The horizontal resize cursor.
	VResize                       ///< The vertical resize cursor.
	CursorCount                   ///< Not a cursor --- should always be last: enables a loop over the cursor types.
)
