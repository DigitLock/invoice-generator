import { apiGet, apiGetBlob } from './api'
import type { InvoiceListResponse, InvoiceResponse } from '@/types/api'

export function fetchInvoices(page = 1, pageSize = 20): Promise<InvoiceListResponse> {
  return apiGet<InvoiceListResponse>(`/api/v1/invoices?page=${page}&page_size=${pageSize}`)
}

export function fetchInvoice(id: number): Promise<InvoiceResponse> {
  return apiGet<InvoiceResponse>(`/api/v1/invoices/${id}`)
}

export async function downloadInvoicePdf(id: number, invoiceNumber: string): Promise<void> {
  const blob = await apiGetBlob(`/api/v1/invoices/${id}/pdf`)
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `${invoiceNumber}.pdf`
  a.click()
  URL.revokeObjectURL(url)
}
