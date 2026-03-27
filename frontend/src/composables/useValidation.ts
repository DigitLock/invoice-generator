import { reactive } from 'vue'
import type { Invoice } from '@/types/invoice'

export interface ValidationErrors {
  sellerName: string
  sellerAddress: string
  buyerName: string
  buyerAddress: string
  buyerEmail: string
  bankName: string
  iban: string
  swift: string
  lineItems: string
}

function createEmptyErrors(): ValidationErrors {
  return {
    sellerName: '',
    sellerAddress: '',
    buyerName: '',
    buyerAddress: '',
    buyerEmail: '',
    bankName: '',
    iban: '',
    swift: '',
    lineItems: '',
  }
}

export function useValidation() {
  const errors = reactive<ValidationErrors>(createEmptyErrors())

  function validateEmail(email: string): boolean {
    if (!email) return true
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
  }

  function validateIban(iban: string): boolean {
    if (!iban) return false
    const cleaned = iban.replace(/\s/g, '')
    return /^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$/.test(cleaned.toUpperCase())
  }

  function validateSwift(swift: string): boolean {
    if (!swift) return false
    const cleaned = swift.replace(/\s/g, '')
    return /^[A-Z0-9]{8}$|^[A-Z0-9]{11}$/.test(cleaned.toUpperCase())
  }

  function validate(invoice: Invoice): boolean {
    Object.assign(errors, createEmptyErrors())

    if (!invoice.seller.name.trim()) {
      errors.sellerName = 'Company name is required'
    }
    if (!invoice.seller.address.trim()) {
      errors.sellerAddress = 'Address is required'
    }

    if (!invoice.buyer.name.trim()) {
      errors.buyerName = 'Company name is required'
    }
    if (!invoice.buyer.address.trim()) {
      errors.buyerAddress = 'Address is required'
    }
    if (invoice.buyer.email && !validateEmail(invoice.buyer.email)) {
      errors.buyerEmail = 'Invalid email format'
    }

    if (!invoice.bankAccount.bankName.trim()) {
      errors.bankName = 'Bank name is required'
    }
    if (!validateIban(invoice.bankAccount.iban)) {
      errors.iban = invoice.bankAccount.iban.trim()
        ? 'Invalid IBAN format (expected: 2 letters + 2 digits + 11-30 alphanumeric)'
        : 'IBAN is required'
    }
    if (!validateSwift(invoice.bankAccount.swift)) {
      errors.swift = invoice.bankAccount.swift.trim()
        ? 'Invalid SWIFT format (expected: 8 or 11 alphanumeric characters)'
        : 'SWIFT is required'
    }

    const hasValidItem = invoice.items.some(
      (item) => item.description.trim() !== '',
    )
    if (!hasValidItem) {
      errors.lineItems = 'At least one line item with a description is required'
    }

    return Object.values(errors).every((e) => e === '')
  }

  function clearErrors() {
    Object.assign(errors, createEmptyErrors())
  }

  return { errors, validate, clearErrors }
}
