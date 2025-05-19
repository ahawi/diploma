<template>
  <div class="interaction">
    <h2>Выберите 5 собак</h2>
    <div class="list">
      <DogCard v-for="d in dogs" :key="d.id" :dog="d" @adopt="select(d)" />
    </div>
    <button :disabled="selected.length !== 5" @click="submit">Отправить</button>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import DogCard from './DogCard.vue'
const dogs = ref([])
const selected = ref([])
async function fetchDogs() {
  dogs.value = await (await fetch('/api/dogs')).json()
}
function select(d) {
  if (!selected.value.includes(d)) {
    if (selected.value.length < 5) selected.value.push(d)
  }
}
async function submit() {
  await fetch('/api/interactions', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(selected.value),
  })
}
onMounted(fetchDogs)
</script>
<style lang="scss" scoped>
.interaction {
  padding: 1rem;
  .list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
  }
  button {
    margin-top: 1rem;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
  }
}
</style>
