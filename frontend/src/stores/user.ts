import type { Organization } from '@/composables/useAuth'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const name = ref<string | null>(null)
  const email = ref<string | null>(null)
  const orgs = ref<Organization[]>([])

  function setUser(userData: { name: string; email: string; orgs: Organization[] }) {
    name.value = userData.name
    email.value = userData.email
    orgs.value = userData.orgs
  }

  const isAuthenticated = () => !!email.value

  return { name, email, orgs, setUser, isAuthenticated }
})
