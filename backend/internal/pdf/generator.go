package pdf

import (
	"bytes"
	"embed"
	"fmt"
	"math"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/shopspring/decimal"
)

//go:embed fonts/Roboto-Regular.ttf fonts/Roboto-Bold.ttf
var fontsFS embed.FS

const (
	pageW      = 210.0
	marginL    = 20.0
	marginR    = 20.0
	contentW   = pageW - marginL - marginR
	fontName   = "Roboto"
	lineHeight = 4.5
	rowPad     = 4.0
)

type InvoiceData struct {
	InvoiceNumber     string
	IssueDate         string // DD.MM.YYYY
	DueDate           string
	ContractReference string
	ExternalReference string
	Currency          string
	VatRate           string

	Seller  PartyData
	Buyer   PartyData
	Bank    BankData
	Items   []ItemData
	Notes   string

	Subtotal  decimal.Decimal
	VatAmount decimal.Decimal
	Total     decimal.Decimal
}

type PartyData struct {
	Name          string
	ContactPerson string
	Email         string
	Address       string
	Phone         string
	VatNumber     string
	RegNumber     string
}

type BankData struct {
	AccountHolder string
	BankName      string
	BankAddress   string
	IBAN          string
	SWIFT         string
}

type ItemData struct {
	Description string
	Quantity    string
	UnitPrice   decimal.Decimal
	Total       decimal.Decimal
}

func formatAmount(d decimal.Decimal, currency string) string {
	f, _ := d.Float64()
	// Format with commas and 2 decimal places
	neg := f < 0
	if neg {
		f = -f
	}
	intPart := int64(f)
	fracPart := int64(math.Round((f - float64(intPart)) * 100))
	if fracPart >= 100 {
		intPart++
		fracPart -= 100
	}

	// Add comma separators
	s := fmt.Sprintf("%d", intPart)
	if len(s) > 3 {
		var parts []string
		for len(s) > 3 {
			parts = append([]string{s[len(s)-3:]}, parts...)
			s = s[:len(s)-3]
		}
		parts = append([]string{s}, parts...)
		s = strings.Join(parts, ",")
	}

	result := fmt.Sprintf("%s.%02d %s", s, fracPart, currency)
	if neg {
		result = "-" + result
	}
	return result
}

func Generate(inv InvoiceData) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetAutoPageBreak(false, 0)

	// Load fonts from embedded FS
	regularBytes, err := fontsFS.ReadFile("fonts/Roboto-Regular.ttf")
	if err != nil {
		return nil, fmt.Errorf("load regular font: %w", err)
	}
	boldBytes, err := fontsFS.ReadFile("fonts/Roboto-Bold.ttf")
	if err != nil {
		return nil, fmt.Errorf("load bold font: %w", err)
	}

	pdf.AddUTF8FontFromBytes(fontName, "", regularBytes)
	pdf.AddUTF8FontFromBytes(fontName, "B", boldBytes)

	pdf.AddPage()
	y := 20.0

	addLine := func(yPos float64) {
		pdf.SetDrawColor(200, 200, 200)
		pdf.SetLineWidth(0.3)
		pdf.Line(marginL, yPos, pageW-marginR, yPos)
	}

	checkPageBreak := func(needed float64) {
		if y+needed > 277 {
			pdf.AddPage()
			y = 20
		}
	}

	// Helper to wrap text and return lines
	wrapText := func(text string, maxW float64) []string {
		if text == "" {
			return nil
		}
		return pdf.SplitText(text, maxW)
	}

	// --- 1. HEADER ---
	invoiceNum := inv.InvoiceNumber
	if invoiceNum == "" {
		invoiceNum = "N/A"
	}

	pdf.SetFont(fontName, "B", 16)
	pdf.SetXY(marginL, y)
	pdf.Cell(contentW, 7, fmt.Sprintf("Invoice #%s", invoiceNum))
	y += 7

	pdf.SetFont(fontName, "", 9)
	pdf.SetTextColor(100, 100, 100)
	dateLine := fmt.Sprintf("Issue Date: %s", inv.IssueDate)
	if inv.DueDate != "" {
		dateLine += fmt.Sprintf("  |  Due Date: %s", inv.DueDate)
	}
	pdf.SetXY(marginL, y)
	pdf.Cell(contentW, 5, dateLine)
	y += 5

	var refParts []string
	if inv.ContractReference != "" {
		refParts = append(refParts, "Contract Ref: "+inv.ContractReference)
	}
	if inv.ExternalReference != "" {
		refParts = append(refParts, "External Ref: "+inv.ExternalReference)
	}
	if len(refParts) > 0 {
		pdf.SetXY(marginL, y)
		pdf.Cell(contentW, 5, strings.Join(refParts, "  |  "))
		y += 5
	}

	pdf.SetTextColor(0, 0, 0)
	y += 2
	addLine(y)
	y += 8

	// --- 2. SELLER & 3. BUYER (side by side) ---
	colW := contentW / 2
	sellerX := marginL
	buyerX := marginL + colW + 5
	colMaxW := colW - 5

	pdf.SetFont(fontName, "B", 10)
	pdf.SetXY(sellerX, y)
	pdf.Cell(colMaxW, 6, "From:")
	pdf.SetXY(buyerX, y)
	pdf.Cell(colMaxW, 6, "Bill To:")
	y += 6

	pdf.SetFont(fontName, "", 9)

	var sellerEntries []string
	if inv.Seller.Name != "" {
		sellerEntries = append(sellerEntries, inv.Seller.Name)
	}
	if inv.Seller.ContactPerson != "" {
		sellerEntries = append(sellerEntries, inv.Seller.ContactPerson)
	}
	if inv.Seller.Address != "" {
		sellerEntries = append(sellerEntries, inv.Seller.Address)
	}
	if inv.Seller.Phone != "" {
		sellerEntries = append(sellerEntries, "Phone: "+inv.Seller.Phone)
	}
	if inv.Seller.VatNumber != "" {
		sellerEntries = append(sellerEntries, "VAT: "+inv.Seller.VatNumber)
	}
	if inv.Seller.RegNumber != "" {
		sellerEntries = append(sellerEntries, "Reg No: "+inv.Seller.RegNumber)
	}

	var buyerEntries []string
	if inv.Buyer.Name != "" {
		buyerEntries = append(buyerEntries, inv.Buyer.Name)
	}
	if inv.Buyer.ContactPerson != "" {
		buyerEntries = append(buyerEntries, inv.Buyer.ContactPerson)
	}
	if inv.Buyer.Email != "" {
		buyerEntries = append(buyerEntries, inv.Buyer.Email)
	}
	if inv.Buyer.Address != "" {
		buyerEntries = append(buyerEntries, inv.Buyer.Address)
	}
	if inv.Buyer.VatNumber != "" {
		buyerEntries = append(buyerEntries, "VAT: "+inv.Buyer.VatNumber)
	}
	if inv.Buyer.RegNumber != "" {
		buyerEntries = append(buyerEntries, "Reg No: "+inv.Buyer.RegNumber)
	}

	maxEntries := len(sellerEntries)
	if len(buyerEntries) > maxEntries {
		maxEntries = len(buyerEntries)
	}

	for i := 0; i < maxEntries; i++ {
		var sellerWrapped, buyerWrapped []string
		if i < len(sellerEntries) {
			sellerWrapped = wrapText(sellerEntries[i], colMaxW)
		}
		if i < len(buyerEntries) {
			buyerWrapped = wrapText(buyerEntries[i], colMaxW)
		}

		for j, line := range sellerWrapped {
			pdf.SetXY(sellerX, y+float64(j)*lineHeight)
			pdf.Cell(colMaxW, lineHeight, line)
		}
		for j, line := range buyerWrapped {
			pdf.SetXY(buyerX, y+float64(j)*lineHeight)
			pdf.Cell(colMaxW, lineHeight, line)
		}

		sellerH := float64(len(sellerWrapped)) * lineHeight
		buyerH := float64(len(buyerWrapped)) * lineHeight
		h := math.Max(sellerH, buyerH)
		if h < lineHeight {
			h = lineHeight
		}
		y += h + 1
	}
	y += 6

	// --- 4. LINE ITEMS TABLE ---
	colDescX := marginL
	descEndX := marginL + contentW*0.48
	qtyStartX := descEndX
	qtyEndX := marginL + contentW*0.60
	priceStartX := qtyEndX
	priceEndX := marginL + contentW*0.82
	totalEndX := pageW - marginR
	colQtyCenterX := (qtyStartX + qtyEndX) / 2
	colPriceCenterX := (priceStartX + priceEndX) / 2
	tableHeaderH := 9.0

	// Table header background
	headerTop := y
	pdf.SetFillColor(245, 245, 245)
	pdf.Rect(marginL, headerTop, contentW, tableHeaderH, "F")
	addLine(headerTop)

	pdf.SetFont(fontName, "B", 9)
	headerBaseline := headerTop + tableHeaderH/2 + 1.5

	pdf.SetXY(colDescX+1, headerBaseline-3)
	pdf.Cell(descEndX-colDescX-2, lineHeight, "Description")

	// Centered headers
	qtyW := qtyEndX - qtyStartX
	pdf.SetXY(colQtyCenterX-qtyW/2, headerBaseline-3)
	pdf.CellFormat(qtyW, lineHeight, "Qty", "", 0, "C", false, 0, "")

	priceW := priceEndX - priceStartX
	pdf.SetXY(colPriceCenterX-priceW/2, headerBaseline-3)
	pdf.CellFormat(priceW, lineHeight, "Unit Price", "", 0, "C", false, 0, "")

	totalW := totalEndX - priceEndX
	pdf.SetXY(priceEndX, headerBaseline-3)
	pdf.CellFormat(totalW, lineHeight, "Total", "", 0, "R", false, 0, "")

	y = headerTop + tableHeaderH
	addLine(y)

	// Table rows
	pdf.SetFont(fontName, "", 9)
	descMaxW := descEndX - colDescX - 2

	for idx, item := range inv.Items {
		// 1. Draw light separator (skip first row)
		if idx > 0 {
			pdf.SetDrawColor(230, 230, 230)
			pdf.SetLineWidth(0.1)
			pdf.Line(marginL, y, pageW-marginR, y)
		}

		// 2. Space after separator, then advance to baseline
		y += rowPad + lineHeight

		checkPageBreak(10)

		// 3. Render first line of text
		descLines := wrapText(item.Description, descMaxW)
		if len(descLines) == 0 {
			descLines = []string{""}
		}

		pdf.SetXY(colDescX+1, y-3)
		pdf.Cell(descMaxW, lineHeight, descLines[0])

		pdf.SetXY(colQtyCenterX-qtyW/2, y-3)
		pdf.CellFormat(qtyW, lineHeight, item.Quantity, "", 0, "C", false, 0, "")

		pdf.SetXY(colPriceCenterX-priceW/2, y-3)
		pdf.CellFormat(priceW, lineHeight, formatAmount(item.UnitPrice, inv.Currency), "", 0, "C", false, 0, "")

		pdf.SetXY(priceEndX, y-3)
		pdf.CellFormat(totalW, lineHeight, formatAmount(item.Total, inv.Currency), "", 0, "R", false, 0, "")

		// 4. Additional wrapped description lines
		for line := 1; line < len(descLines); line++ {
			y += lineHeight
			pdf.SetXY(colDescX+1, y-3)
			pdf.Cell(descMaxW, lineHeight, descLines[line])
		}

		// 5. Space before next separator
		y += rowPad
	}

	y += 2
	addLine(y)
	y += 7

	// --- 5. TOTALS ---
	checkPageBreak(30)
	totalsLabelX := marginL + contentW*0.6
	totalsValueX := pageW - marginR
	totalsW := totalsValueX - totalsLabelX

	pdf.SetFont(fontName, "", 9)
	pdf.SetXY(totalsLabelX, y)
	pdf.Cell(30, 5, "Subtotal:")
	pdf.SetXY(totalsLabelX+30, y)
	pdf.CellFormat(totalsW-30, 5, formatAmount(inv.Subtotal, inv.Currency), "", 0, "R", false, 0, "")
	y += 6

	pdf.SetXY(totalsLabelX, y)
	pdf.Cell(30, 5, fmt.Sprintf("VAT (%s%%):", inv.VatRate))
	pdf.SetXY(totalsLabelX+30, y)
	pdf.CellFormat(totalsW-30, 5, formatAmount(inv.VatAmount, inv.Currency), "", 0, "R", false, 0, "")
	y += 5

	// Total row with background highlight
	totalRowH := 8.0
	pdf.SetFillColor(245, 245, 245)
	pdf.Rect(totalsLabelX-2, y-1, totalsW+4, totalRowH, "F")

	pdf.SetFont(fontName, "B", 11)
	y += 4
	pdf.SetXY(totalsLabelX, y-3)
	pdf.Cell(30, 5, "Total:")
	pdf.SetXY(totalsLabelX+30, y-3)
	pdf.CellFormat(totalsW-30, 5, formatAmount(inv.Total, inv.Currency), "", 0, "R", false, 0, "")
	y += 12

	// --- 6. PAYMENT DETAILS ---
	checkPageBreak(40)

	pdf.SetFont(fontName, "B", 10)
	pdf.SetXY(marginL, y)
	pdf.Cell(contentW, 6, "Payment Details")
	y += 7

	pdf.SetFont(fontName, "", 9)
	labelCol := marginL
	valueCol := marginL + 30.0
	valueMaxW := contentW - 30.0

	type pair struct{ label, value string }
	var paymentPairs []pair
	if inv.Bank.AccountHolder != "" {
		paymentPairs = append(paymentPairs, pair{"Account Holder:", inv.Bank.AccountHolder})
	}
	if inv.Bank.BankName != "" {
		paymentPairs = append(paymentPairs, pair{"Bank:", inv.Bank.BankName})
	}
	if inv.Bank.BankAddress != "" {
		paymentPairs = append(paymentPairs, pair{"Bank Address:", inv.Bank.BankAddress})
	}
	if inv.Bank.IBAN != "" {
		paymentPairs = append(paymentPairs, pair{"IBAN:", inv.Bank.IBAN})
	}
	if inv.Bank.SWIFT != "" {
		paymentPairs = append(paymentPairs, pair{"SWIFT:", inv.Bank.SWIFT})
	}

	for _, p := range paymentPairs {
		wrapped := wrapText(p.value, valueMaxW)
		if len(wrapped) == 0 {
			wrapped = []string{p.value}
		}
		checkPageBreak(float64(len(wrapped))*lineHeight + 2)

		pdf.SetFont(fontName, "B", 9)
		pdf.SetXY(labelCol, y)
		pdf.Cell(30, lineHeight, p.label)

		pdf.SetFont(fontName, "", 9)
		for j, line := range wrapped {
			pdf.SetXY(valueCol, y+float64(j)*lineHeight)
			pdf.Cell(valueMaxW, lineHeight, line)
		}
		y += float64(len(wrapped))*lineHeight + 1
	}

	// --- 7. NOTES ---
	if inv.Notes != "" {
		y += 7
		checkPageBreak(20)
		pdf.SetFont(fontName, "B", 10)
		pdf.SetXY(marginL, y)
		pdf.Cell(contentW, 6, "Notes")
		y += 6

		pdf.SetFont(fontName, "", 9)
		noteLines := wrapText(inv.Notes, contentW)
		for _, line := range noteLines {
			pdf.SetXY(marginL, y)
			pdf.Cell(contentW, lineHeight, line)
			y += lineHeight
		}
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("generate pdf: %w", err)
	}
	return buf.Bytes(), nil
}
