import type { LoginRequest, LoginResponse } from '@/types/api'

const AUTH_BASE = import.meta.env.VITE_AUTH_URL || 'http://localhost:8080'

export async function login(credentials: LoginRequest): Promise<LoginResponse> {
  const response = await fetch(`${AUTH_BASE}/api/v1/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(credentials),
  })

  const body = await response.json()

  if (!response.ok) {
    throw new Error(body.error || 'Login failed')
  }

  return body.data as LoginResponse
}
