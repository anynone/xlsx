package xlsx

// XLSXWorksheet directly maps the worksheet element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXWorksheet struct {
	Dimension     XLSXDimension
	SheetViews    XLSXSheetViews
	SheetFormatPr XLSXSheetFormatPr
	SheetData     XLSXSheetData
}

// XLSXDimension directly maps the dimension element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXDimension struct {
	Ref string "attr"
}

// XLSXSheetViews directly maps the sheetViews element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXSheetViews struct {
	SheetView []XLSXSheetView
}

// XLSXSheetView directly maps the sheetView element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXSheetView struct {
	TabSelected    string "attr"
	WorkbookViewID string "attr"
	Selection      XLSXSelection
}


// XLSXSelection directly maps the selection element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXSelection struct {
	ActiveCell string "attr"
	SQRef      string "attr"
}

// XLSXSheetFormatPr directly maps the sheetFormatPr element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXSheetFormatPr struct {
	BaseColWidth     string "attr"
	DefaultRowHeight string "attr"
}

// XLSXSheetData directly maps the sheetData element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXSheetData struct {
	Row []XLSXRow
}

// XLSXRow directly maps the row element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXRow struct {
	R     string "attr"
	Spans string "attr"
	C     []XLSXC
}

// XLSXC directly maps the c element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXC struct {
	R string "attr"
	T string "attr"
	V XLSXV
}


// XLSXV directly maps the v element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much
// as I need.
type XLSXV struct {
	Data string "chardata"
}

