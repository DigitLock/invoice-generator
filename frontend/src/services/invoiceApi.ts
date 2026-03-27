import { apiGet, apiPost, apiPut, apiPatch, apiDelete, apiGetBlob } from './api'
import type { InvoiceListResponse, InvoiceResponse } from '@/types/api'

export interface InvoiceItemInput {
  description: string
  quantity: string
  unit_price: string
}

export interface CreateInvoiceInput {
  company_id: number
  client_id: number
  bank_account_id: number
  issue_date: string
  due_date: string
  currency: string
  vat_rate: string
  contract_reference?: string | null
  external_reference?: string | null
  notes?: string | null
  items: InvoiceItemInput[]
}

export interface UpdateInvoiceInput extends CreateInvoiceInput {
  invoice_number: string
}

export function fetchInvoices(params: {
  page?: number
  page_size?: number
  status?: string
  search?: string
  date_from?: string
  date_to?: string
} = {}): Promise<InvoiceListResponse> {
  const query = new URLSearchParams()
  query.set('page', String(params.page ?? 1))
  query.set('page_size', String(params.page_size ?? 20))
  if (params.status) query.set('status', params.status)
  if (params.search) query.set('search', params.search)
  if (params.date_from) query.set('date_from', params.date_from)
  if (params.date_to) query.set('date_to', params.date_to)
  return apiGet<InvoiceListResponse>(`/api/v1/invoices?${query}`)
}

export function fetchInvoice(id: number): Promise<InvoiceResponse> {
  return apiGet<InvoiceResponse>(`/api/v1/invoices/${id}`)
}

export function createInvoice(data: CreateInvoiceInput): Promise<InvoiceResponse> {
  return apiPost<InvoiceResponse>('/api/v1/invoices', data)
}

export function updateInvoice(id: number, data: UpdateInvoiceInput): Promise<InvoiceResponse> {
  return apiPut<InvoiceResponse>(`/api/v1/invoices/${id}`, data)
}

export function deleteInvoice(id: number): Promise<void> {
  return apiDelete(`/api/v1/invoices/${id}`)
}

export function updateInvoiceStatus(id: number, status: string): Promise<InvoiceResponse> {
  return apiPatch<InvoiceResponse>(`/api/v1/invoices/${id}/status`, { status })
}

export function updateInvoiceOverdue(id: number, isOverdue: boolean): Promise<InvoiceResponse> {
  return apiPatch<InvoiceResponse>(`/api/v1/invoices/${id}/overdue`, { is_overdue: isOverdue })
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
