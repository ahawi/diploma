<template>
  <nav class="navbar">
    <router-link to="/animals">Животные</router-link>
    <router-link to="/profile">Личный кабинет</router-link>
    <UserMenu />
  </nav>
</template>
<script setup>
import UserMenu from './UserMenu.vue'
</script>
<style lang="scss" scoped>
.navbar {
  display: flex;
  gap: 1rem;
  align-items: center;
  padding: 1rem;
  background: #333;
  color: #fff;
  a {
    color: #fff;
    text-decoration: none;
  }
}
</style>
