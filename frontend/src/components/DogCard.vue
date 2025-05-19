<template>
  <div class="dog-card">
    <img :src="dog.photo" alt="фото собаки" />
    <h3>{{ dog.name }}</h3>
    <p>{{ dog.breed }}, {{ dog.age }} лет</p>
    <div class="actions">
      <button @click="$emit('favorite', dog)">❤</button>
      <button @click="$emit('adopt', dog)">Усыновить</button>
    </div>
  </div>
</template>
<script setup>
import { defineProps } from 'vue'
const props = defineProps({ dog: Object })
</script>
<style lang="scss" scoped>
.dog-card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 1rem;
  text-align: center;
  img {
    width: 100%;
    height: auto;
    border-radius: 4px;
  }
  .actions {
    margin-top: 0.5rem;
    display: flex;
    justify-content: space-around;
  }
}
</style>
