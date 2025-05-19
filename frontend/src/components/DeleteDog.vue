<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h2>Удалить собаку?</h2>
      <p>Вы уверены, что хотите удалить карточку {{ dog.name }}?</p>
      <div class="actions">
        <button @click="confirm">Да</button>
        <button @click="close">Нет</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
const props = defineProps({ show: Boolean, dog: Object })
const emit = defineEmits(['close', 'deleted'])
async function confirm() {
  await fetch(`/api/dogs/${props.dog.id}`, { method: 'DELETE' })
  emit('deleted')
  close()
}
function close() {
  emit('close')
}
</script>

<style lang="scss" scoped>
.actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}
</style>
