import { computed } from 'vue'
import type { Invoice } from '@/types/invoice'
import type { Ref } from 'vue'

export function useInvoiceCalculations(invoice: Ref<Invoice>) {
  const subtotal = computed(() => {
    // TODO: sum of item totals (quantity * unitPrice)
    return '0.00'
  })

  const vatAmount = computed(() => {
    // TODO: subtotal * (vatRate / 100)
    return '0.00'
  })

  const total = computed(() => {
    // TODO: subtotal + vatAmount
    return '0.00'
  })

  return { subtotal, vatAmount, total }
}
