import { useMutation } from '@tanstack/vue-query'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/user'
import { api } from '@/lib/axios'

export interface LoginPayload {
  email: string
  password: string
}

export interface RegisterPayload {
  name: string
  email: string
  password: string
  org_name: string
}

export interface Organization {
  id: string
  name: string
}

export interface AuthResponse {
  message: string
  name: string
  email: string
  orgs: Organization[]
}

export function useLogin() {
  const router = useRouter()
  const { setUser } = useAuthStore()

  return useMutation({
    mutationFn: (payload: LoginPayload) =>
      api.post<AuthResponse>('/auth/login', payload).then((res) => res.data),
    onSuccess(data) {
      setUser({ name: data.name, email: data.email, orgs: data.orgs })
      router.push('/')
    },
  })
}

export function useRegister() {
  const router = useRouter()
  const { setUser } = useAuthStore()

  return useMutation({
    mutationFn: (payload: RegisterPayload) =>
      api.post<AuthResponse>('/auth/register', payload).then((res) => res.data),
    onSuccess(data) {
      setUser({ name: data.name, email: data.email, orgs: data.orgs })
      router.push('/')
    },
  })
}
