package core

type Role uint32
type State uint32

// AT-SPI2 roles identifiers, taken from https://gitlab.gnome.org/GNOME/at-spi2-core/-/blob/main/atspi/atspi-constants.h
const (
	ATSPI_ROLE_INVALID Role = iota
	ATSPI_ROLE_ACCELERATOR_LABEL
	ATSPI_ROLE_ALERT
	ATSPI_ROLE_ANIMATION
	ATSPI_ROLE_ARROW
	ATSPI_ROLE_CALENDAR
	ATSPI_ROLE_CANVAS
	ATSPI_ROLE_CHECK_MENU_ITEM
	ATSPI_ROLE_CHECK_BOX
	ATSPI_ROLE_COLOR_CHOOSER
	ATSPI_ROLE_COLUMN_HEADER
	ATSPI_ROLE_COMBO_BOX
	ATSPI_ROLE_DATE_EDITOR
	ATSPI_ROLE_DESKTOP_ICON
	ATSPI_ROLE_DESKTOP_FRAME
	ATSPI_ROLE_DIAL
	ATSPI_ROLE_DIALOG
	ATSPI_ROLE_DIRECTORY_PANE
	ATSPI_ROLE_DRAWING_AREA
	ATSPI_ROLE_FILE_CHOOSER
	ATSPI_ROLE_FILLER
	ATSPI_ROLE_FOCUS_TRAVERSABLE
	ATSPI_ROLE_FONT_CHOOSER
	ATSPI_ROLE_FRAME
	ATSPI_ROLE_GLASS_PANE
	ATSPI_ROLE_HTML_CONTAINER
	ATSPI_ROLE_ICON
	ATSPI_ROLE_IMAGE
	ATSPI_ROLE_INTERNAL_FRAME
	ATSPI_ROLE_LABEL
	ATSPI_ROLE_LAYERED_PANE
	ATSPI_ROLE_LIST
	ATSPI_ROLE_LIST_ITEM
	ATSPI_ROLE_MENU
	ATSPI_ROLE_MENU_BAR
	ATSPI_ROLE_MENU_ITEM
	ATSPI_ROLE_OPTION_PANE
	ATSPI_ROLE_PAGE_TAB
	ATSPI_ROLE_PAGE_TAB_LIST
	ATSPI_ROLE_PANEL
	ATSPI_ROLE_PASSWORD_TEXT
	ATSPI_ROLE_POPUP_MENU
	ATSPI_ROLE_PROGRESS_BAR
	ATSPI_ROLE_BUTTON
	ATSPI_ROLE_RADIO_BUTTON
	ATSPI_ROLE_RADIO_MENU_ITEM
	ATSPI_ROLE_ROOT_PANE
	ATSPI_ROLE_ROW_HEADER
	ATSPI_ROLE_SCROLL_BAR
	ATSPI_ROLE_SCROLL_PANE
	ATSPI_ROLE_SEPARATOR
	ATSPI_ROLE_SLIDER
	ATSPI_ROLE_SPIN_BUTTON
	ATSPI_ROLE_SPLIT_PANE
	ATSPI_ROLE_STATUS_BAR
	ATSPI_ROLE_TABLE
	ATSPI_ROLE_TABLE_CELL
	ATSPI_ROLE_TABLE_COLUMN_HEADER
	ATSPI_ROLE_TABLE_ROW_HEADER
	ATSPI_ROLE_TEAROFF_MENU_ITEM
	ATSPI_ROLE_TERMINAL
	ATSPI_ROLE_TEXT
	ATSPI_ROLE_TOGGLE_BUTTON
	ATSPI_ROLE_TOOL_BAR
	ATSPI_ROLE_TOOL_TIP
	ATSPI_ROLE_TREE
	ATSPI_ROLE_TREE_TABLE
	ATSPI_ROLE_UNKNOWN
	ATSPI_ROLE_VIEWPORT
	ATSPI_ROLE_WINDOW
	ATSPI_ROLE_EXTENDED
	ATSPI_ROLE_HEADER
	ATSPI_ROLE_FOOTER
	ATSPI_ROLE_PARAGRAPH

	ATSPI_ROLE_RULER
	ATSPI_ROLE_APPLICATION
	ATSPI_ROLE_AUTOCOMPLETE
	ATSPI_ROLE_EDITBAR
	ATSPI_ROLE_EMBEDDED
	ATSPI_ROLE_ENTRY
	ATSPI_ROLE_CHART
	ATSPI_ROLE_CAPTION
	ATSPI_ROLE_DOCUMENT_FRAME
	ATSPI_ROLE_HEADING
	ATSPI_ROLE_PAGE
	ATSPI_ROLE_SECTION
	ATSPI_ROLE_REDUNDANT_OBJECT
	ATSPI_ROLE_FORM
	ATSPI_ROLE_LINK
	ATSPI_ROLE_INPUT_METHOD_WINDOW
	ATSPI_ROLE_TABLE_ROW
	ATSPI_ROLE_TREE_ITEM
	ATSPI_ROLE_DOCUMENT_SPREADSHEET
	ATSPI_ROLE_DOCUMENT_PRESENTATION
	ATSPI_ROLE_DOCUMENT_TEXT
	ATSPI_ROLE_DOCUMENT_WEB
	ATSPI_ROLE_DOCUMENT_EMAIL
	ATSPI_ROLE_COMMENT
	ATSPI_ROLE_LIST_BOX
	ATSPI_ROLE_GROUPING
	ATSPI_ROLE_IMAGE_MAP
	ATSPI_ROLE_NOTIFICATION
	ATSPI_ROLE_INFO_BAR
)

var roleNames = map[Role]string{
	ATSPI_ROLE_INVALID:               "Invalid",
	ATSPI_ROLE_ACCELERATOR_LABEL:     "Accelerator Label",
	ATSPI_ROLE_ALERT:                 "Alert",
	ATSPI_ROLE_ANIMATION:             "Animation",
	ATSPI_ROLE_ARROW:                 "Arrow",
	ATSPI_ROLE_CALENDAR:              "Calendar",
	ATSPI_ROLE_CANVAS:                "Canvas",
	ATSPI_ROLE_CHECK_MENU_ITEM:       "Check Menu Item",
	ATSPI_ROLE_CHECK_BOX:             "Check Box",
	ATSPI_ROLE_COLOR_CHOOSER:         "Color Chooser",
	ATSPI_ROLE_COLUMN_HEADER:         "Column Header",
	ATSPI_ROLE_COMBO_BOX:             "Combo Box",
	ATSPI_ROLE_DATE_EDITOR:           "Date Editor",
	ATSPI_ROLE_DESKTOP_ICON:          "Desktop Icon",
	ATSPI_ROLE_DESKTOP_FRAME:         "Desktop Frame",
	ATSPI_ROLE_DIAL:                  "Dial",
	ATSPI_ROLE_DIALOG:                "Dialog",
	ATSPI_ROLE_DIRECTORY_PANE:        "Directory Pane",
	ATSPI_ROLE_DRAWING_AREA:          "Drawing Area",
	ATSPI_ROLE_FILE_CHOOSER:          "File Chooser",
	ATSPI_ROLE_FILLER:                "Filler",
	ATSPI_ROLE_FOCUS_TRAVERSABLE:     "Focus Traversable",
	ATSPI_ROLE_FONT_CHOOSER:          "Font Chooser",
	ATSPI_ROLE_FRAME:                 "Frame",
	ATSPI_ROLE_GLASS_PANE:            "Glass Pane",
	ATSPI_ROLE_HTML_CONTAINER:        "HTML Container",
	ATSPI_ROLE_ICON:                  "Icon",
	ATSPI_ROLE_IMAGE:                 "Image",
	ATSPI_ROLE_INTERNAL_FRAME:        "Internal Frame",
	ATSPI_ROLE_LABEL:                 "Label",
	ATSPI_ROLE_LAYERED_PANE:          "Layered Pane",
	ATSPI_ROLE_LIST:                  "List",
	ATSPI_ROLE_LIST_ITEM:             "List Item",
	ATSPI_ROLE_MENU:                  "Menu",
	ATSPI_ROLE_MENU_BAR:              "Menu Bar",
	ATSPI_ROLE_MENU_ITEM:             "Menu Item",
	ATSPI_ROLE_OPTION_PANE:           "Option Pane",
	ATSPI_ROLE_PAGE_TAB:              "Page Tab",
	ATSPI_ROLE_PAGE_TAB_LIST:         "Page Tab List",
	ATSPI_ROLE_PANEL:                 "Panel",
	ATSPI_ROLE_PASSWORD_TEXT:         "Password Text",
	ATSPI_ROLE_POPUP_MENU:            "Popup Menu",
	ATSPI_ROLE_PROGRESS_BAR:          "Progress Bar",
	ATSPI_ROLE_BUTTON:                "Button",
	ATSPI_ROLE_RADIO_BUTTON:          "Radio Button",
	ATSPI_ROLE_RADIO_MENU_ITEM:       "Radio Menu Item",
	ATSPI_ROLE_ROOT_PANE:             "Root Pane",
	ATSPI_ROLE_ROW_HEADER:            "Row Header",
	ATSPI_ROLE_SCROLL_BAR:            "Scroll Bar",
	ATSPI_ROLE_SCROLL_PANE:           "Scroll Pane",
	ATSPI_ROLE_SEPARATOR:             "Separator",
	ATSPI_ROLE_SLIDER:                "Slider",
	ATSPI_ROLE_SPIN_BUTTON:           "Spin Button",
	ATSPI_ROLE_SPLIT_PANE:            "Split Pane",
	ATSPI_ROLE_STATUS_BAR:            "Status Bar",
	ATSPI_ROLE_TABLE:                 "Table",
	ATSPI_ROLE_TABLE_CELL:            "Table Cell",
	ATSPI_ROLE_TABLE_COLUMN_HEADER:   "Table Column Header",
	ATSPI_ROLE_TABLE_ROW_HEADER:      "Table Row Header",
	ATSPI_ROLE_TEAROFF_MENU_ITEM:     "Tearoff Menu Item",
	ATSPI_ROLE_TERMINAL:              "Terminal",
	ATSPI_ROLE_TEXT:                  "Text",
	ATSPI_ROLE_TOGGLE_BUTTON:         "Toggle Button",
	ATSPI_ROLE_TOOL_BAR:              "Tool Bar",
	ATSPI_ROLE_TOOL_TIP:              "Tool Tip",
	ATSPI_ROLE_TREE:                  "Tree",
	ATSPI_ROLE_TREE_TABLE:            "Tree Table",
	ATSPI_ROLE_UNKNOWN:               "Unknown",
	ATSPI_ROLE_VIEWPORT:              "Viewport",
	ATSPI_ROLE_WINDOW:                "Window",
	ATSPI_ROLE_EXTENDED:              "Extended",
	ATSPI_ROLE_HEADER:                "Header",
	ATSPI_ROLE_FOOTER:                "Footer",
	ATSPI_ROLE_PARAGRAPH:             "Paragraph",
	ATSPI_ROLE_RULER:                 "Ruler",
	ATSPI_ROLE_APPLICATION:           "Application",
	ATSPI_ROLE_AUTOCOMPLETE:          "Autocomplete",
	ATSPI_ROLE_EDITBAR:               "Edit Bar",
	ATSPI_ROLE_EMBEDDED:              "Embedded",
	ATSPI_ROLE_ENTRY:                 "Entry",
	ATSPI_ROLE_CHART:                 "Chart",
	ATSPI_ROLE_CAPTION:               "Caption",
	ATSPI_ROLE_DOCUMENT_FRAME:        "Document Frame",
	ATSPI_ROLE_HEADING:               "Heading",
	ATSPI_ROLE_PAGE:                  "Page",
	ATSPI_ROLE_SECTION:               "Section",
	ATSPI_ROLE_REDUNDANT_OBJECT:      "Redundant Object",
	ATSPI_ROLE_FORM:                  "Form",
	ATSPI_ROLE_LINK:                  "Link",
	ATSPI_ROLE_INPUT_METHOD_WINDOW:   "Input Method Window",
	ATSPI_ROLE_TABLE_ROW:             "Table Row",
	ATSPI_ROLE_TREE_ITEM:             "Tree Item",
	ATSPI_ROLE_DOCUMENT_SPREADSHEET:  "Document Spreadsheet",
	ATSPI_ROLE_DOCUMENT_PRESENTATION: "Document Presentation",
	ATSPI_ROLE_DOCUMENT_TEXT:         "Document Text",
	ATSPI_ROLE_DOCUMENT_WEB:          "Document Web",
	ATSPI_ROLE_DOCUMENT_EMAIL:        "Document Email",
	ATSPI_ROLE_COMMENT:               "Comment",
	ATSPI_ROLE_LIST_BOX:              "List Box",
	ATSPI_ROLE_GROUPING:              "Grouping",
	ATSPI_ROLE_IMAGE_MAP:             "Image Map",
	ATSPI_ROLE_NOTIFICATION:          "Notification",
	ATSPI_ROLE_INFO_BAR:              "Info Bar",
}

func (r Role) String() string {
	if name, exists := roleNames[r]; exists {
		return name
	}
	return "Unknown Role"
}

const (
	AtspiLocation     = "/org/a11y/bus"
	AtspiBusName      = "org.a11y.Bus"
	AtspiInterface    = "org.a11y.atspi"
	AtspiRegistryPath = "/org/a11y/atspi/registry"
)
