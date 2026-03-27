export interface Seller {
  name: string
  contactPerson: string
  address: string
  phone: string
  vatNumber: string
  regNumber: string
}

export interface Buyer {
  name: string
  contactPerson: string
  email: string
  address: string
  vatNumber: string
  regNumber: string
}

export interface BankAccount {
  bankName: string
  bankAddress: string
  accountHolder: string
  iban: string
  swift: string
  currency: string
}

export interface LineItem {
  description: string
  quantity: string
  unitPrice: string
  total: string
}

export type Currency = 'EUR' | 'RSD' | string

export interface Invoice {
  invoiceNumber: string
  issueDate: string
  dueDate: string
  currency: Currency
  vatRate: string
  contractReference: string
  externalReference: string
  notes: string
  seller: Seller
  buyer: Buyer
  bankAccount: BankAccount
  items: LineItem[]
  subtotal: string
  vatAmount: string
  total: string
}
