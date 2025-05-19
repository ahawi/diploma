<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h2>Удалить пользователя?</h2>
      <p>Удалить {{ user.name }}?</p>
      <div class="actions">
        <button @click="confirm">Да</button>
        <button @click="close">Нет</button>
      </div>
    </div>
  </div>
</template>
<script setup>
import { defineProps, defineEmits } from 'vue'
const props = defineProps({ show: Boolean, user: Object })
const emit = defineEmits(['close', 'deleted'])
async function confirm() {
  await fetch(`/api/users/${props.user.id}`, { method: 'DELETE' })
  emit('deleted')
  close()
}
function close() {
  emit('close')
}
</script>
<style lang="scss" scoped></style>
