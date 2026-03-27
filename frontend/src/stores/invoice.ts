import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Invoice, LineItem, Seller, Buyer, BankAccount } from '@/types/invoice'

function createEmptySeller(): Seller {
  return {
    name: '',
    contactPerson: '',
    address: '',
    phone: '',
    vatNumber: '',
    regNumber: '',
  }
}

function createEmptyBuyer(): Buyer {
  return {
    name: '',
    contactPerson: '',
    email: '',
    address: '',
    vatNumber: '',
    regNumber: '',
  }
}

function createEmptyBankAccount(): BankAccount {
  return {
    bankName: '',
    bankAddress: '',
    accountHolder: '',
    iban: '',
    swift: '',
    currency: 'EUR',
  }
}

export function createEmptyLineItem(): LineItem {
  return {
    description: '',
    quantity: '',
    unitPrice: '',
    total: '0.00',
  }
}

function createEmptyInvoice(): Invoice {
  const today = new Date().toISOString().slice(0, 10)
  return {
    invoiceNumber: '',
    issueDate: today,
    dueDate: '',
    currency: 'EUR',
    vatRate: '0.00',
    contractReference: '',
    externalReference: '',
    notes: '',
    seller: createEmptySeller(),
    buyer: createEmptyBuyer(),
    bankAccount: createEmptyBankAccount(),
    items: [createEmptyLineItem()],
    subtotal: '0.00',
    vatAmount: '0.00',
    total: '0.00',
  }
}

export const useInvoiceStore = defineStore('invoice', () => {
  const invoice = ref<Invoice>(createEmptyInvoice())

  function resetInvoice() {
    invoice.value = createEmptyInvoice()
  }

  return {
    invoice,
    resetInvoice,
  }
})
