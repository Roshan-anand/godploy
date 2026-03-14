import { api } from '@/lib/axios'
import { useQuery } from '@tanstack/vue-query'

export function useAuthGithubApp() {
  return useQuery({
    queryKey: ['github-app'],
    queryFn: () => api.get('/provider/github/app/create'),
    enabled: false,
  })
}
