<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#"> Navbar </a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto">
          <li class="nav-item">
            <RouterLink class="nav-link active" aria-current="page" to="/"> Home </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link active" aria-current="page" to="/books"> Books </RouterLink>
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="navBarDropDown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              Admin
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <RouterLink class="dropdown-item" to="/books"> Manage Users </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item" to="/books"> Add User </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item" to="/books"> Manage Books </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item" to="/books"> Add Book </RouterLink>
              </li>
            </ul>
          </li>
          <li class="nav-item">
            <RouterLink v-if="!authStore.isAuthenticated()" class="nav-link" to="/login">
              Login
            </RouterLink>
            <a v-else href="javascript:void(0)" class="nav-link" @click="logout()"> Logout </a>
          </li>
        </ul>
        <span class="navbar-text">
          {{ authStore.user?.name }}
        </span>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const logout = () => authStore.logout()
</script>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
