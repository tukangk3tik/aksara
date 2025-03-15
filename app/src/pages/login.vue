<script setup>
import { useAuthStore } from '@/stores/auth.store';
import logo from '@images/logo.svg?raw'
import { nextTick, ref } from 'vue'
import { useRouter } from 'vue-router';
// import authV1BottomShape from '@images/svg/auth-v1-bottom-shape.svg?url'
// import authV1TopShape from '@images/svg/auth-v1-top-shape.svg?url'

const authStore = useAuthStore()
const router = useRouter()

const form = ref(null)
const email = ref('')
const password = ref('')
const remember = ref(false)
const isLoading = ref(false)
const isPasswordVisible = ref(false)
const loginMessage = ref('')
const loginStatus = ref(null)

const rules = {
  required: (v) => !!v || "Harus diisi",
  email: (v) => /.+@.+\..+/.test(v) || "E-mail harus valid",
  minLength: (length) => (v) =>
    (v && v.length >= length) || `Password harus terdiri dari ${length} characters`,
}

const handleLogin = async () => {
  isLoading.value = true
  loginStatus.value = null

  try {
    loginMessage.value = ''
    const { valid } = await form.value.validate()
    if (!valid) return;

    await authStore.login({
      email: email.value,
      password: password.value
    })

    loginStatus.value = 'success'
    loginMessage.value = 'Login berhasil!'
    
    await nextTick()

    setTimeout(() => {
      router.push('/')
    }, 800)
  } catch (error) {
    setTimeout(() => {
      router.push('/')
      loginStatus.value = 'error'
      loginMessage.value = error
      isLoading.value = false
    }, 800)
  } 
}

</script>

<template>
  <div class="auth-wrapper d-flex align-center justify-center pa-4">
    <div class="position-relative my-sm-16">

      <!-- 👉 Top shape -->
      <!-- <VImg
        :src="authV1TopShape"
        class="text-primary auth-v1-top-shape d-none d-sm-block"
      /> -->

      <!-- 👉 Bottom shape -->
      <!-- <VImg
        :src="authV1BottomShape"
        class="text-primary auth-v1-bottom-shape d-none d-sm-block"
      /> -->

      <!-- 👉 Auth Card -->
      <VCard
        class="auth-card"
        max-width="460"
        :class="$vuetify.display.smAndUp ? 'pa-6' : 'pa-0'"
      >
        <VCardItem class="justify-center">
          <RouterLink
            to="/"
            class="app-logo">
            <!-- eslint-disable vue/no-v-html -->
            <div
              class="d-flex"
              v-html="logo"
            />
            <h1 class="app-logo-title">
              Aksara 
            </h1>
          </RouterLink>
        </VCardItem>

        <VCardText>
          <h4 class="text-h4 mb-1">
            Selamat datang di Aksara! 👋🏻
          </h4>
          <p class="mb-0">
            Silahkan login ke akun Anda
          </p>
        </VCardText>

        <VCardText>
          <VForm ref="form" @submit.prevent="handleLogin">
            <VRow>
              <!-- email -->
              <VCol cols="12">
                <VTextField
                  v-model="email"
                  autofocus
                  label="Email"
                  type="email"
                  placeholder="budi@aksara.com"
                  :rules="[rules.required, rules.email]"
                />
              </VCol>

              <!-- password -->
              <VCol cols="12">
                <VTextField
                  v-model="password"
                  label="Password"
                  placeholder="············"
                  :type="isPasswordVisible ? 'text' : 'password'"
                  autocomplete="password"
                  :append-inner-icon="isPasswordVisible ? 'bx-hide' : 'bx-show'"
                  @click:append-inner="isPasswordVisible = !isPasswordVisible"
                  :rules="[rules.required, rules.minLength(8)]"
                />

                <!-- remember me checkbox -->
                <div class="d-flex align-center justify-space-between flex-wrap my-6">
                  <VCheckbox
                    v-model="remember"
                    label="Ingat saya" />

                  <a
                    class="text-primary"
                    href="javascript:void(0)">
                    Lupa Password?
                  </a>
                </div>

                <!-- login button -->
                <VBtn
                  block
                  type="submit"
                  :disabled="isLoading"
                  :loading="isLoading"
                >
                  Login
                </VBtn>
              </VCol>

              <!-- error message -->
              <VCol cols="12">
                <VAlert :type="loginStatus" v-if="loginMessage">
                {{ loginMessage }}
                </VAlert>
              </VCol>
            </VRow>
          </VForm>
        </VCardText>
      </VCard>
    </div>
  </div>
</template>

<style lang="scss">
@use "@core/scss/template/pages/page-auth";
</style>
