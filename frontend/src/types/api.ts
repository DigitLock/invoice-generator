export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  token: string
  user: {
    id: number
    family_id: number
    email: string
    name: string
  }
  expires_at: string
}

export interface CompanyResponse {
  id: number
  name: string
  contact_person: string
  address: string
  phone: string | null
  vat_number: string | null
  reg_number: string | null
  created_at: string
  updated_at: string
}

export interface ClientResponse {
  id: number
  name: string
  contact_person: string | null
  email: string | null
  address: string
  vat_number: string | null
  reg_number: string | null
  contract_reference: string | null
  contract_notes: string | null
  status: 'active' | 'inactive'
  created_at: string
  updated_at: string
}

export interface BankAccountResponse {
  id: number
  company_id: number
  bank_name: string
  bank_address: string
  account_holder: string
  iban: string
  swift: string
  currency: string
  is_default: boolean
  created_at: string
  updated_at: string
}

export interface InvoiceItemResponse {
  id: number
  invoice_id: number
  description: string
  quantity: string
  unit_price: string
  total: string
  created_at: string
  updated_at: string
}

export interface InvoiceResponse {
  id: number
  invoice_number: string
  user_id: number
  family_id: number
  company_id: number
  client_id: number
  bank_account_id: number
  issue_date: string
  due_date: string
  currency: string
  status: string
  is_overdue: boolean
  vat_rate: string
  subtotal: string
  vat_amount: string
  total: string
  contract_reference: string | null
  external_reference: string | null
  notes: string | null
  company?: CompanyResponse
  client?: ClientResponse
  bank_account?: BankAccountResponse
  items?: InvoiceItemResponse[]
  created_at: string
  updated_at: string
}

export interface InvoiceListItem {
  id: number
  invoice_number: string
  issue_date: string
  due_date: string
  status: string
  is_overdue: boolean
  currency: string
  company_name: string
  client_name: string
  subtotal: string
  vat_amount: string
  total: string
  items_count: number
  created_at: string
  updated_at: string
}

export interface PaginationMeta {
  page: number
  page_size: number
  total_items: number
  total_pages: number
  has_next: boolean
  has_previous: boolean
}

export interface InvoiceListResponse {
  invoices: InvoiceListItem[]
  pagination: PaginationMeta
}

export interface ApiError {
  error: string
  details?: { field: string; message: string }[]
}
