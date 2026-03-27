import { apiGet, apiPost, apiPut, apiDelete } from './api'
import type { BankAccountResponse } from '@/types/api'

export interface BankAccountInput {
  bank_name: string
  bank_address: string
  account_holder: string
  iban: string
  swift: string
  currency: string
  is_default: boolean
}

export function fetchBankAccounts(companyId: number): Promise<BankAccountResponse[]> {
  return apiGet<BankAccountResponse[]>(`/api/v1/companies/${companyId}/bank-accounts`)
}

export function createBankAccount(companyId: number, data: BankAccountInput): Promise<BankAccountResponse> {
  return apiPost<BankAccountResponse>(`/api/v1/companies/${companyId}/bank-accounts`, data)
}

export function updateBankAccount(id: number, data: BankAccountInput): Promise<BankAccountResponse> {
  return apiPut<BankAccountResponse>(`/api/v1/bank-accounts/${id}`, data)
}

export function deleteBankAccount(id: number): Promise<void> {
  return apiDelete(`/api/v1/bank-accounts/${id}`)
}
