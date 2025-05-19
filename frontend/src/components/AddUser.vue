<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h2>Добавить пользователя</h2>
      <form @submit.prevent="submit">
        <label>Имя <input v-model="form.name" required /></label>
        <label>Email <input type="email" v-model="form.email" required /></label>
        <label
          >Роль
          <select v-model="form.role">
            <option>admin</option>
            <option>staff</option>
          </select></label
        >
        <div class="actions">
          <button type="submit">Создать</button>
          <button type="button" @click="close">Отмена</button>
        </div>
      </form>
    </div>
  </div>
</template>
<script setup>
import { ref, defineProps, defineEmits } from 'vue'
const props = defineProps({ show: Boolean })
const emit = defineEmits(['close', 'created'])
const form = ref({ name: '', email: '', role: 'staff' })
async function submit() {
  await fetch('/api/users', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(form.value),
  })
  emit('created')
  close()
}
function close() {
  emit('close')
}
</script>
<style lang="scss" scoped></style>
