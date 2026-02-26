<script setup lang="ts">
import { Icon } from '@iconify/vue'
import { useColorMode } from '@vueuse/core'
import { Button } from '@/components/ui/button'

const mode = useColorMode({ disableTransition: false })
</script>

<template>
  <Button variant="secondary" @click="() => (mode = mode == 'light' ? 'dark' : 'light')">
    <Icon v-if="mode == 'light'" icon="radix-icons:sun" />
    <Icon v-else icon="radix-icons:moon" />
  </Button>
</template>
