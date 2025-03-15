import { useAuthStore } from '@/stores/auth.store';

export function authGuard(to, from, next) {
  const authStore = useAuthStore()
  // const user = authStore.user

  if (!localStorage.getItem("access_token")) {
    return next('/login')
  }

  next()
}