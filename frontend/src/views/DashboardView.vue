<script setup lang="ts">
import DashBoardSidebar from '@/components/DashBoardSideBar.vue'
import { SidebarProvider, SidebarInset, SidebarTrigger } from '@/components/ui/sidebar'
import ModeToggle from '@/components/ModeToggle.vue'
import { Button } from '@/components/ui/button'

const githuRedirect = () => (window.location.href = '/api/provider/github/app/create')
</script>

<template>
  <SidebarProvider>
    <DashBoardSidebar />
    <SidebarInset>
      <header class="flex justify-between items-center p-2">
        <SidebarTrigger class="-ml-1" />
        <ModeToggle />
      </header>
      <main>
        <Button variant="secondary" @click="githuRedirect"> connect github </Button>
      </main>
    </SidebarInset>
  </SidebarProvider>
</template>
