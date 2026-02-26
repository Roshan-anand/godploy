<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useForm } from '@tanstack/vue-form'
import { z } from 'zod'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Checkbox } from '@/components/ui/checkbox'
import AuthBranding from '@/components/AuthBranding.vue'
import { useRegister } from '@/composables/useAuth'
import { toast } from 'vue-sonner'
import { watch } from 'vue'

const { isError, isPending, mutate, error } = useRegister()

const form = useForm({
  defaultValues: {
    name: '',
    email: '',
    password: '',
    organisation: '',
    rememberMe: false,
  },
  onSubmit: async ({ value }) => {
    mutate({
      name: value.name,
      email: value.email,
      password: value.password,
      org_name: value.organisation,
    })
  },
})

watch([isError], () => {
  if (isError.value) {
    toast.error(error.value?.message || 'An error occurred while registering. Please try again.')
  }
})

const Field = form.Field
const Subscribe = form.Subscribe
</script>

<template>
  <div class="grid min-h-svh lg:grid-cols-2">
    <AuthBranding />
    <!-- Form Section -->
    <div class="flex items-center justify-center p-8">
      <div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-100px">
        <div class="flex flex-col space-y-2 text-center">
          <h1 class="text-2xl font-semibold tracking-tight">Create an account</h1>
          <p class="text-sm text-muted-foreground">
            Enter your details below to create your account
          </p>
        </div>

        <form
          @submit="
            (e) => {
              e.preventDefault()
              e.stopPropagation()
              form.handleSubmit()
            }
          "
        >
          <div class="grid gap-4">
            <!-- Name -->
            <Field
              name="name"
              :validators="{ onChange: z.string().min(2, 'Name must be at least 2 characters') }"
            >
              <template v-slot="{ field, state }">
                <div class="grid gap-2">
                  <Label :for="field.name">Name</Label>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="text"
                    placeholder="John Doe"
                    :value="state.value"
                    @blur="field.handleBlur"
                    @input="(e: Event) => field.handleChange((e.target as HTMLInputElement).value)"
                  />
                  <p v-if="state.meta.errors.length" class="text-sm font-medium text-destructive">
                    {{ state.meta.errors[0]?.message }}
                  </p>
                </div>
              </template>
            </Field>

            <!-- Email -->
            <Field
              name="email"
              :validators="{ onChange: z.string().email('Please enter a valid email') }"
            >
              <template v-slot="{ field, state }">
                <div class="grid gap-2">
                  <Label :for="field.name">Email</Label>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="email"
                    placeholder="name@example.com"
                    :value="state.value"
                    @blur="field.handleBlur"
                    @input="(e: Event) => field.handleChange((e.target as HTMLInputElement).value)"
                  />
                  <p v-if="state.meta.errors.length" class="text-sm font-medium text-destructive">
                    {{ state.meta.errors[0]?.message || 'Invalid email' }}
                  </p>
                </div>
              </template>
            </Field>

            <!-- Password -->
            <Field
              name="password"
              :validators="{
                onChange: z.string().min(8, 'Password must be at least 8 characters'),
              }"
            >
              <template v-slot="{ field, state }">
                <div class="grid gap-2">
                  <Label :for="field.name">Password</Label>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="password"
                    :value="state.value"
                    @blur="field.handleBlur"
                    @input="(e: Event) => field.handleChange((e.target as HTMLInputElement).value)"
                  />
                  <p v-if="state.meta.errors.length" class="text-sm font-medium text-destructive">
                    {{ state.meta.errors[0]?.message || 'invalid password' }}
                  </p>
                </div>
              </template>
            </Field>

            <!-- Organisation -->
            <Field
              name="organisation"
              :validators="{
                onChange: z.string().min(2, 'Organisation must be at least 2 characters'),
              }"
            >
              <template v-slot="{ field, state }">
                <div class="grid gap-2">
                  <Label :for="field.name">Organisation</Label>
                  <Input
                    :id="field.name"
                    :name="field.name"
                    type="text"
                    placeholder="Acme Inc."
                    :value="state.value"
                    @blur="field.handleBlur"
                    @input="(e: Event) => field.handleChange((e.target as HTMLInputElement).value)"
                  />
                  <p v-if="state.meta.errors.length" class="text-sm font-medium text-destructive">
                    {{ state.meta.errors[0] }}
                  </p>
                </div>
              </template>
            </Field>

            <!-- Remember Me -->
            <Field name="rememberMe">
              <template v-slot="{ field, state }">
                <div class="flex items-center space-x-2 mt-2">
                  <Checkbox
                    :id="field.name"
                    :checked="state.value"
                    @update:checked="field.handleChange"
                  />
                  <Label
                    :for="field.name"
                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  >
                    Remember me
                  </Label>
                </div>
              </template>
            </Field>

            <Subscribe>
              <template v-slot="{ canSubmit, isSubmitting }">
                <Button type="submit" class="w-full mt-2" :disabled="!canSubmit || isPending">
                  {{ isSubmitting || isPending ? 'Creating account...' : 'Create account' }}
                </Button>
              </template>
            </Subscribe>
          </div>
        </form>

        <p class="px-8 text-center text-sm text-muted-foreground">
          Already have an account?
          <RouterLink to="/login" class="underline underline-offset-4 hover:text-primary">
            Log in
          </RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>
