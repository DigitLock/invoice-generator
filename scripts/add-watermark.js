import { readFile, writeFile } from 'fs/promises'
import { PDFDocument, StandardFonts, rgb, degrees } from 'pdf-lib'

const PDF_PATH = new URL('../frontend/public/invoice-preview.pdf', import.meta.url)

const pdfBytes = await readFile(PDF_PATH)
const doc = await PDFDocument.load(pdfBytes)
const font = await doc.embedFont(StandardFonts.HelveticaBold)

const MARGIN = 50
const text = 'EXAMPLE'

for (const page of doc.getPages()) {
  const { width, height } = page.getSize()

  // Content area diagonal (corner to corner excluding margins)
  const contentW = width - 2 * MARGIN
  const contentH = height - 2 * MARGIN
  const diagonal = Math.sqrt(contentW * contentW + contentH * contentH)
  const angle = Math.atan2(contentH, contentW)

  // Size the font so the text spans the full diagonal
  const fontSize = diagonal / font.widthOfTextAtSize(text, 1)
  const textWidth = font.widthOfTextAtSize(text, fontSize)

  // Start at bottom-left content corner
  // Offset slightly to center the text height on the diagonal
  const textHeight = fontSize * 0.7 // approx cap height
  const offsetX = (textHeight / 2) * Math.sin(angle)
  const offsetY = (textHeight / 2) * Math.cos(angle)

  page.drawText(text, {
    x: MARGIN + offsetX,
    y: MARGIN - offsetY,
    size: fontSize,
    font,
    color: rgb(200 / 255, 200 / 255, 200 / 255),
    opacity: 0.3,
    rotate: degrees((angle * 180) / Math.PI),
  })
}

const output = await doc.save()
await writeFile(PDF_PATH, output)
console.log('Watermark added to invoice-preview.pdf')
