<template>
  <Header />
  <RouterView />
  <Footer />
</template>

<script setup lang="ts">
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { onBeforeMount } from 'vue'
import { useAuthStore } from './stores/auth'

const authStore = useAuthStore()

onBeforeMount(() => authStore.validateToken())
</script>
