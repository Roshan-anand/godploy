<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import 'vue-sonner/style.css'
import { Toaster } from '@/components/ui/sonner'
import ModeToggle from './components/ModeToggle.vue'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
    <ModeToggle />
    <Button @click="() => toast.success('toast successfully shown')">Click me</Button>
  </header>

  <RouterView />
  <Toaster position="bottom-right" />
</template>
