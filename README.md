# a1notation

support for Microsoft Excel and Google Spreadsheet.

## Usage

rowNum and colNum start index 0.

```go
import "github.com/koluku/a1notation"

// CW101
index := a1notation.FromIndex(100, 100)
// ALL
colName := a1notation.ColumnLetterFrom("999")
// 9, 26
rowNum, colNum := a1notation.ToIndex("AA10")
```
