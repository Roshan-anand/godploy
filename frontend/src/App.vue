<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ModeToggle from './components/ModeToggle.vue'
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
  </header>

  <RouterView />
</template>
