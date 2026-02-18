<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ThemeProvider from '@/components/ThemeProvider.vue'
import ModeToggle from '@/components/ModeToggle.vue'
</script>

<template>
  <ThemeProvider default-theme="system" storage-key="vite-ui-theme">
    <header>
      <div class="wrapper">
        <nav>
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/about">About</RouterLink>
        </nav>
      </div>
      <ModeToggle />
    </header>

    <RouterView />
  </ThemeProvider>
</template>
