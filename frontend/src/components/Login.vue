<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form
          @submit.prevent="onSubmit"
          ref="form"
          autocomplete="off"
          class="needs-validation"
          novalidate
        >
          <div class="mb-3">
            <label for="email" class="form-label" style="width: 7.5rem"> Email </label>
            <input
              v-model="v$.email.$model"
              type="email"
              name="email"
              required="true"
              style="width: 20rem"
            />
            <div v-if="v$.email.$error" class="form-error">
              <p v-for="error of v$.email.$errors" :key="error.$uid">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <div class="mb-3">
            <label for="password" class="form-label" style="width: 7.5rem"> Password </label>
            <input
              v-model="v$.password.$model"
              type="password"
              name="password"
              required="true"
              style="width: 20rem"
            />
            <div v-if="v$.password.$error" class="form-error">
              <p v-for="error of v$.password.$errors" :key="error.$uid">
                {{ error.$message }}
              </p>
            </div>
          </div>
          <input type="submit" class="btn btn-primary" value="Login" />
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { required, email, minLength } from '@vuelidate/validators'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const form = reactive({
  email: '',
  password: ''
})

const rules = {
  email: { required, email },
  password: { required, minLength: minLength(5) }
}

const v$ = useVuelidate(rules, form)

const emit = defineEmits(['error'])

const onSubmit = async () => {
  const isValid = await v$.value.$validate()

  if (!isValid) {
    return
  }
  authStore.login(form.email, form.password)
}
</script>

<style scoped>
.form-error {
  margin-left: 7.5rem;

  p {
    color: red;
    font-size: 12px;
  }
}
</style>
