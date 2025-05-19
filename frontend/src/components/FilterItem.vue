<script setup>
defineProps({
  label: String,
  options: Array,
  id: String,
  modelValue: String,
})

const emit = defineEmits(['update:modelValue'])

const updateValue = (event) => {
  emit('update:modelValue', event.target.value)
}
</script>

<template>
  <div class="sidebar__item">
    <label :for="id" class="sidebar__title">{{ label }}</label>
    <select class="sidebar__select" :id="id" :value="modelValue" @change="updateValue">
      <option v-for="option in options" :key="option" :value="option">{{ option }}</option>
    </select>
  </div>
</template>

<style scoped></style>
