import { apiGet, apiPost, apiPut, apiDelete } from './api'
import type { CompanyResponse } from '@/types/api'

export interface CompanyInput {
  name: string
  contact_person: string
  address: string
  phone?: string | null
  vat_number?: string | null
  reg_number?: string | null
}

export function fetchCompanies(): Promise<CompanyResponse[]> {
  return apiGet<CompanyResponse[]>('/api/v1/companies')
}

export function fetchCompany(id: number): Promise<CompanyResponse> {
  return apiGet<CompanyResponse>(`/api/v1/companies/${id}`)
}

export function createCompany(data: CompanyInput): Promise<CompanyResponse> {
  return apiPost<CompanyResponse>('/api/v1/companies', data)
}

export function updateCompany(id: number, data: CompanyInput): Promise<CompanyResponse> {
  return apiPut<CompanyResponse>(`/api/v1/companies/${id}`, data)
}

export function deleteCompany(id: number): Promise<void> {
  return apiDelete(`/api/v1/companies/${id}`)
}
