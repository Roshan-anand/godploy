import { createRouter, createWebHashHistory } from 'vue-router'
import DashboardView from '@/views/DashboardView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import { useAuthStore } from '@/stores/user'
import { type AuthResponse } from '@/composables/useAuth'
import { api } from '@/lib/axios'
import { queryClient } from '@/main'
import axios from 'axios'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'dashboard', component: DashboardView },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/register', name: 'register', component: RegisterView },
  ],
})

const PUBLIC_ROUTES = ['login', 'register']

router.beforeEach(async (to) => {
  const { setUser } = useAuthStore()

  // skip guard for public routes
  if (PUBLIC_ROUTES.includes(to.name as string)) return true

  const { isAuthenticated } = useAuthStore()

  if (isAuthenticated()) return true

  // no user data, verify with server
  // const res = await api.get<AuthResponse>('/auth/user')

  try {
    const res = await queryClient.fetchQuery({
      queryKey: ['auth', 'user'],
      queryFn: () => api.get<AuthResponse>('/auth/user'),
    })

    const { name, email, orgs } = res.data
    setUser({ name, email, orgs })
    return true
  } catch (err) {
    if (axios.isAxiosError(err) && err.response?.status === 403) return { name: 'register' }

    // go to login
    return { name: 'login' }
  }
})

export default router
