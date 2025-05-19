<template>
  <AddDog :show="show" @close="close" @created="updated" v-bind="dogProps" />
</template>

<script setup>
import AddDog from './AddDog.vue'
import { defineProps, defineEmits } from 'vue'

const props = defineProps({ show: Boolean, dog: Object })
const emit = defineEmits(['close', 'updated'])

const dogProps = { ...props.dog }
function updated() {
  emit('updated')
}
function close() {
  emit('close')
}
</script>

<style lang="scss" scoped>
</style>
