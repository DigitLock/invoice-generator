export function useValidation() {
  function validateEmail(_email: string): boolean {
    // TODO: implement email format validation
    return true
  }

  function validateIban(_iban: string): boolean {
    // TODO: implement IBAN format validation
    return true
  }

  function validateSwift(_swift: string): boolean {
    // TODO: implement SWIFT/BIC format validation
    return true
  }

  return { validateEmail, validateIban, validateSwift }
}
