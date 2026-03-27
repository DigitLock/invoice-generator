import type { Invoice } from '@/types/invoice'

export async function generatePdf(_invoice: Invoice): Promise<Blob> {
  // TODO: implement client-side PDF generation with jsPDF or pdf-lib
  throw new Error('PDF generation not yet implemented')
}
