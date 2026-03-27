import { jsPDF } from 'jspdf'
import type { Invoice } from '@/types/invoice'
import { robotoRegularBase64 } from '@/assets/fonts/roboto-regular-base64'
import { robotoBoldBase64 } from '@/assets/fonts/roboto-bold-base64'

const PAGE_W = 210
const MARGIN_L = 20
const MARGIN_R = 20
const CONTENT_W = PAGE_W - MARGIN_L - MARGIN_R
const FONT_NAME = 'Roboto'
const LINE_HEIGHT = 4.5

function formatDate(iso: string): string {
  if (!iso) return ''
  const [y, m, d] = iso.split('-')
  return `${d}.${m}.${y}`
}

function formatAmount(amount: string, currency: string): string {
  const num = parseFloat(amount)
  if (isNaN(num)) return `0.00 ${currency}`
  const formatted = num.toLocaleString('en-US', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  })
  return `${formatted} ${currency}`
}

function registerFonts(doc: jsPDF) {
  doc.addFileToVFS('Roboto-Regular.ttf', robotoRegularBase64)
  doc.addFont('Roboto-Regular.ttf', FONT_NAME, 'normal')
  doc.addFileToVFS('Roboto-Bold.ttf', robotoBoldBase64)
  doc.addFont('Roboto-Bold.ttf', FONT_NAME, 'bold')
  doc.setFont(FONT_NAME, 'normal')
}

export async function generatePdf(invoice: Invoice): Promise<Blob> {
  const doc = new jsPDF({ unit: 'mm', format: 'a4' })
  registerFonts(doc)
  let y = 20

  function addLine(yPos: number) {
    doc.setDrawColor(200)
    doc.setLineWidth(0.3)
    doc.line(MARGIN_L, yPos, PAGE_W - MARGIN_R, yPos)
  }

  function checkPageBreak(needed: number) {
    if (y + needed > 277) {
      doc.addPage()
      y = 20
    }
  }

  // --- 1. HEADER (compact) ---
  const invoiceNum = invoice.invoiceNumber || 'N/A'

  doc.setFont(FONT_NAME, 'bold')
  doc.setFontSize(16)
  doc.text(`Invoice #${invoiceNum}`, MARGIN_L, y)
  y += 7

  doc.setFont(FONT_NAME, 'normal')
  doc.setFontSize(9)
  doc.setTextColor(100)
  const dateLine = `Issue Date: ${formatDate(invoice.issueDate)}  |  Due Date: ${formatDate(invoice.dueDate)}`
  doc.text(dateLine, MARGIN_L, y)
  y += 5

  const refParts: string[] = []
  if (invoice.contractReference) refParts.push(`Contract Ref: ${invoice.contractReference}`)
  if (invoice.externalReference) refParts.push(`External Ref: ${invoice.externalReference}`)
  if (refParts.length > 0) {
    doc.text(refParts.join('  |  '), MARGIN_L, y)
    y += 5
  }

  doc.setTextColor(0)
  y += 2
  addLine(y)
  y += 8

  // --- 2. SELLER & 3. BUYER (side by side) ---
  const colW = CONTENT_W / 2
  const sellerX = MARGIN_L
  const buyerX = MARGIN_L + colW + 5
  const colMaxW = colW - 5

  doc.setFontSize(10)
  doc.setFont(FONT_NAME, 'bold')
  doc.text('From:', sellerX, y)
  doc.text('Bill To:', buyerX, y)
  y += 6

  doc.setFontSize(9)
  doc.setFont(FONT_NAME, 'normal')

  const sellerEntries: string[] = []
  if (invoice.seller.name) sellerEntries.push(invoice.seller.name)
  if (invoice.seller.contactPerson) sellerEntries.push(invoice.seller.contactPerson)
  if (invoice.seller.address) sellerEntries.push(invoice.seller.address)
  if (invoice.seller.phone) sellerEntries.push(`Phone: ${invoice.seller.phone}`)
  if (invoice.seller.vatNumber) sellerEntries.push(`VAT: ${invoice.seller.vatNumber}`)
  if (invoice.seller.regNumber) sellerEntries.push(`Reg No: ${invoice.seller.regNumber}`)

  const buyerEntries: string[] = []
  if (invoice.buyer.name) buyerEntries.push(invoice.buyer.name)
  if (invoice.buyer.contactPerson) buyerEntries.push(invoice.buyer.contactPerson)
  if (invoice.buyer.email) buyerEntries.push(invoice.buyer.email)
  if (invoice.buyer.address) buyerEntries.push(invoice.buyer.address)
  if (invoice.buyer.vatNumber) buyerEntries.push(`VAT: ${invoice.buyer.vatNumber}`)
  if (invoice.buyer.regNumber) buyerEntries.push(`Reg No: ${invoice.buyer.regNumber}`)

  const maxEntries = Math.max(sellerEntries.length, buyerEntries.length)
  for (let i = 0; i < maxEntries; i++) {
    const sellerWrapped: string[] = sellerEntries[i]
      ? doc.splitTextToSize(sellerEntries[i], colMaxW)
      : []
    const buyerWrapped: string[] = buyerEntries[i]
      ? doc.splitTextToSize(buyerEntries[i], colMaxW)
      : []

    if (sellerWrapped.length > 0) doc.text(sellerWrapped, sellerX, y)
    if (buyerWrapped.length > 0) doc.text(buyerWrapped, buyerX, y)

    const sellerH = sellerWrapped.length * LINE_HEIGHT
    const buyerH = buyerWrapped.length * LINE_HEIGHT
    y += Math.max(sellerH, buyerH, LINE_HEIGHT) + 1
  }
  y += 6

  // --- 4. LINE ITEMS TABLE ---
  const colDescX = MARGIN_L
  const descEndX = MARGIN_L + CONTENT_W * 0.48
  const qtyStartX = descEndX
  const qtyEndX = MARGIN_L + CONTENT_W * 0.60
  const priceStartX = qtyEndX
  const priceEndX = MARGIN_L + CONTENT_W * 0.82
  const totalStartX = priceEndX
  const totalEndX = PAGE_W - MARGIN_R
  const colQtyCenterX = (qtyStartX + qtyEndX) / 2
  const colPriceCenterX = (priceStartX + priceEndX) / 2
  const tableHeaderH = 9
  const ROW_PAD = 3

  // Table header background
  doc.setFillColor(245, 245, 245)
  doc.rect(MARGIN_L, y - 1, CONTENT_W, tableHeaderH, 'F')

  addLine(y - 1)

  doc.setFont(FONT_NAME, 'bold')
  doc.setFontSize(9)
  y += (tableHeaderH / 2) + 1
  doc.text('Description', colDescX + 1, y)
  doc.text('Qty', colQtyCenterX, y, { align: 'center' })
  doc.text('Unit Price', colPriceCenterX, y, { align: 'center' })
  doc.text('Total', totalEndX, y, { align: 'right' })
  y = y + (tableHeaderH / 2) - 1
  addLine(y)
  y += ROW_PAD

  // Table rows
  doc.setFont(FONT_NAME, 'normal')
  const descMaxW = descEndX - colDescX - 2
  for (let idx = 0; idx < invoice.items.length; idx++) {
    const item = invoice.items[idx]
    if (!item.description && !item.quantity && !item.unitPrice) continue
    checkPageBreak(10)

    const descLines = doc.splitTextToSize(item.description || '', descMaxW)
    doc.text(descLines, colDescX + 1, y)
    doc.text(item.quantity || '0', colQtyCenterX, y, { align: 'center' })
    doc.text(
      formatAmount(item.unitPrice, invoice.currency),
      colPriceCenterX,
      y,
      { align: 'center' },
    )
    doc.text(
      formatAmount(item.total, invoice.currency),
      totalEndX,
      y,
      { align: 'right' },
    )

    const rowH = Math.max(descLines.length * LINE_HEIGHT, 6)
    y += rowH + ROW_PAD

    // Light row separator with padding (except after last item)
    if (idx < invoice.items.length - 1) {
      doc.setDrawColor(230, 230, 230)
      doc.setLineWidth(0.1)
      doc.line(MARGIN_L, y, PAGE_W - MARGIN_R, y)
      y += ROW_PAD
    }
  }

  y += 2
  addLine(y)
  y += 7

  // --- 5. TOTALS ---
  checkPageBreak(30)
  const totalsLabelX = MARGIN_L + CONTENT_W * 0.6
  const totalsValueX = PAGE_W - MARGIN_R

  doc.setFontSize(9)
  doc.setFont(FONT_NAME, 'normal')
  doc.text('Subtotal:', totalsLabelX, y)
  doc.text(formatAmount(invoice.subtotal, invoice.currency), totalsValueX, y, {
    align: 'right',
  })
  y += 6

  doc.text(`VAT (${invoice.vatRate}%):`, totalsLabelX, y)
  doc.text(formatAmount(invoice.vatAmount, invoice.currency), totalsValueX, y, {
    align: 'right',
  })
  y += 5

  // Total row with background highlight
  const totalRowH = 8
  doc.setFillColor(245, 245, 245)
  doc.rect(totalsLabelX - 2, y - 1, totalsValueX - totalsLabelX + 4, totalRowH, 'F')

  doc.setFont(FONT_NAME, 'bold')
  doc.setFontSize(11)
  y += 4
  doc.text('Total:', totalsLabelX, y)
  doc.text(formatAmount(invoice.total, invoice.currency), totalsValueX, y, {
    align: 'right',
  })
  y += 12

  // --- 6. PAYMENT DETAILS ---
  checkPageBreak(40)

  doc.setFont(FONT_NAME, 'bold')
  doc.setFontSize(10)
  doc.text('Payment Details', MARGIN_L, y)
  y += 7

  doc.setFontSize(9)
  doc.setFont(FONT_NAME, 'normal')

  const labelCol = MARGIN_L
  const valueCol = MARGIN_L + 30
  const valueMaxW = CONTENT_W - 30

  const paymentPairs: [string, string][] = []
  if (invoice.bankAccount.accountHolder) {
    paymentPairs.push(['Account Holder:', invoice.bankAccount.accountHolder])
  }
  if (invoice.bankAccount.bankName) {
    paymentPairs.push(['Bank:', invoice.bankAccount.bankName])
  }
  if (invoice.bankAccount.bankAddress) {
    paymentPairs.push(['Bank Address:', invoice.bankAccount.bankAddress])
  }
  if (invoice.bankAccount.iban) {
    paymentPairs.push(['IBAN:', invoice.bankAccount.iban])
  }
  if (invoice.bankAccount.swift) {
    paymentPairs.push(['SWIFT:', invoice.bankAccount.swift])
  }

  for (const [label, value] of paymentPairs) {
    const wrapped: string[] = doc.splitTextToSize(value, valueMaxW)
    checkPageBreak(wrapped.length * LINE_HEIGHT + 2)
    doc.setFont(FONT_NAME, 'bold')
    doc.text(label, labelCol, y)
    doc.setFont(FONT_NAME, 'normal')
    doc.text(wrapped, valueCol, y)
    y += wrapped.length * LINE_HEIGHT + 1
  }

  // --- 7. NOTES ---
  if (invoice.notes) {
    y += 7
    checkPageBreak(20)
    doc.setFont(FONT_NAME, 'bold')
    doc.setFontSize(10)
    doc.text('Notes', MARGIN_L, y)
    y += 6
    doc.setFont(FONT_NAME, 'normal')
    doc.setFontSize(9)
    const noteLines: string[] = doc.splitTextToSize(invoice.notes, CONTENT_W)
    doc.text(noteLines, MARGIN_L, y)
  }

  const filename = invoice.invoiceNumber
    ? `${invoice.invoiceNumber}.pdf`
    : 'invoice.pdf'

  doc.save(filename)

  return doc.output('blob')
}
