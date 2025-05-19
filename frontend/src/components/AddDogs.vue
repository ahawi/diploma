<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h2>Добавить новую собаку</h2>
      <form @submit.prevent="submit">
        <label>Имя <input v-model="form.name" required /></label>
        <label>Возраст <input type="number" v-model.number="form.age" required /></label>
        <label>Порода <input v-model="form.breed" required /></label>
        <label
          >Пол
          <select v-model="form.gender">
            <option value="male">Мальчик</option>
            <option value="female">Девочка</option>
          </select>
        </label>
        <label>Описание <textarea v-model="form.description"></textarea></label>
        <label>Особенности <textarea v-model="form.features"></textarea></label>
        <label>Фото <input type="file" @change="onFileChange" /></label>
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

const form = ref({
  name: '',
  age: null,
  breed: '',
  gender: 'male',
  description: '',
  features: '',
  photo: null,
})

function onFileChange(e) {
  form.value.photo = e.target.files[0]
}

async function submit() {
  const data = new FormData()
  Object.entries(form.value).forEach(([k, v]) => data.append(k, v))
  await fetch('/api/dogs', { method: 'POST', body: data })
  emit('created')
  close()
}

function close() {
  emit('close')
}
</script>

<style lang="scss" scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}
.modal-content {
  background: #fff;
  padding: 2rem;
  border-radius: 8px;
  width: 400px;
  form {
    display: flex;
    flex-direction: column;
    label {
      margin-bottom: 1rem;
      display: flex;
      flex-direction: column;
      input,
      textarea,
      select {
        margin-top: 0.5rem;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
      }
    }
    .actions {
      display: flex;
      justify-content: flex-end;
      gap: 1rem;
    }
  }
}
</style>
