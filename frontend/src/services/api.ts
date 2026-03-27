import type { ApiError } from '@/types/api'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8081'

export class ApiRequestError extends Error {
  status: number
  details?: { field: string; message: string }[]

  constructor(status: number, body: ApiError) {
    super(body.error)
    this.status = status
    this.details = body.details
  }
}

function getToken(): string | null {
  return localStorage.getItem('jwt_token')
}

export async function apiFetch<T>(
  path: string,
  options: RequestInit = {},
): Promise<T> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...((options.headers as Record<string, string>) || {}),
  }

  const token = getToken()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const response = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers,
  })

  if (response.status === 401) {
    localStorage.removeItem('jwt_token')
    localStorage.removeItem('user')
    window.location.href = '/login'
    throw new ApiRequestError(401, { error: 'Session expired' })
  }

  if (response.status === 204) {
    return undefined as T
  }

  const body = await response.json()

  if (!response.ok) {
    throw new ApiRequestError(response.status, body)
  }

  return body as T
}

export function apiGet<T>(path: string): Promise<T> {
  return apiFetch<T>(path)
}

export function apiPost<T>(path: string, data: unknown): Promise<T> {
  return apiFetch<T>(path, { method: 'POST', body: JSON.stringify(data) })
}

export function apiPut<T>(path: string, data: unknown): Promise<T> {
  return apiFetch<T>(path, { method: 'PUT', body: JSON.stringify(data) })
}

export function apiPatch<T>(path: string, data: unknown): Promise<T> {
  return apiFetch<T>(path, { method: 'PATCH', body: JSON.stringify(data) })
}

export function apiDelete(path: string): Promise<void> {
  return apiFetch<void>(path, { method: 'DELETE' })
}

export async function apiGetBlob(path: string): Promise<Blob> {
  const headers: Record<string, string> = {}
  const token = getToken()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const response = await fetch(`${API_BASE}${path}`, { headers })

  if (!response.ok) {
    const body = await response.json()
    throw new ApiRequestError(response.status, body)
  }

  return response.blob()
}
