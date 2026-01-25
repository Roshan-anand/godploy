<script setup lang="ts">
const testRoute = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/test')
    const data = await res.json()
    console.log(data)
  } catch (error) {
    console.error('Error in testRoute:', error)
  }
}
</script>

<template>
  <main>
    <button @click="testRoute">click</button>
  </main>
</template>
