import { apiGet, apiPost, apiPut, apiDelete } from './api'
import type { ClientResponse } from '@/types/api'

export interface ClientInput {
  name: string
  contact_person?: string | null
  email?: string | null
  address: string
  vat_number?: string | null
  reg_number?: string | null
  contract_reference?: string | null
  contract_notes?: string | null
  status: string
}

export function fetchClients(status?: string): Promise<ClientResponse[]> {
  const query = status ? `?status=${status}` : ''
  return apiGet<ClientResponse[]>(`/api/v1/clients${query}`)
}

export function fetchClient(id: number): Promise<ClientResponse> {
  return apiGet<ClientResponse>(`/api/v1/clients/${id}`)
}

export function createClient(data: ClientInput): Promise<ClientResponse> {
  return apiPost<ClientResponse>('/api/v1/clients', data)
}

export function updateClient(id: number, data: ClientInput): Promise<ClientResponse> {
  return apiPut<ClientResponse>(`/api/v1/clients/${id}`, data)
}

export function deleteClient(id: number): Promise<void> {
  return apiDelete(`/api/v1/clients/${id}`)
}
