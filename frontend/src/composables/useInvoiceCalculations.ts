import { watch } from 'vue'
import type { Invoice } from '@/types/invoice'
import type { Ref } from 'vue'

function toNum(val: string): number {
  const n = parseFloat(val)
  return isNaN(n) ? 0 : n
}

function toFixed2(n: number): string {
  return n.toFixed(2)
}

function calcItemTotal(qty: string, price: string): string {
  return toFixed2(Math.round(toNum(qty) * toNum(price) * 100) / 100)
}

export function useInvoiceCalculations(invoice: Ref<Invoice>) {
  watch(
    () => invoice.value.items.map((i) => [i.quantity, i.unitPrice]),
    () => {
      for (const item of invoice.value.items) {
        item.total = calcItemTotal(item.quantity, item.unitPrice)
      }
    },
    { deep: true, immediate: true },
  )

  watch(
    () => [
      invoice.value.items.map((i) => i.total),
      invoice.value.vatRate,
    ],
    () => {
      const subtotal = invoice.value.items.reduce(
        (sum, item) => sum + toNum(item.total),
        0,
      )
      const vatAmount =
        Math.round(subtotal * (toNum(invoice.value.vatRate) / 100) * 100) / 100
      const total = Math.round((subtotal + vatAmount) * 100) / 100

      invoice.value.subtotal = toFixed2(subtotal)
      invoice.value.vatAmount = toFixed2(vatAmount)
      invoice.value.total = toFixed2(total)
    },
    { deep: true, immediate: true },
  )
}
