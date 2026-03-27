import { readFile, writeFile } from 'fs/promises'
import { PDFDocument, StandardFonts, rgb, degrees } from 'pdf-lib'

const PDF_PATH = new URL('../frontend/public/invoice-preview.pdf', import.meta.url)

const pdfBytes = await readFile(PDF_PATH)
const doc = await PDFDocument.load(pdfBytes)
const font = await doc.embedFont(StandardFonts.HelveticaBold)

for (const page of doc.getPages()) {
  const { width, height } = page.getSize()
  const text = 'EXAMPLE'
  const fontSize = 80
  const textWidth = font.widthOfTextAtSize(text, fontSize)

  page.drawText(text, {
    x: (width - textWidth * Math.cos(Math.PI / 4)) / 2,
    y: (height - textWidth * Math.sin(Math.PI / 4)) / 2,
    size: fontSize,
    font,
    color: rgb(200 / 255, 200 / 255, 200 / 255),
    opacity: 0.3,
    rotate: degrees(45),
  })
}

const output = await doc.save()
await writeFile(PDF_PATH, output)
console.log('Watermark added to invoice-preview.pdf')
