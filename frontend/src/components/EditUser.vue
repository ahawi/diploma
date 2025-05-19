<template>
  <AddUser :show="show" @close="close" @created="updated" v-bind="userProps" />
</template>
<script setup>
import AddUser from './AddUser.vue'
import { defineProps, defineEmits } from 'vue'
const props = defineProps({ show: Boolean, user: Object })
const emit = defineEmits(['close', 'updated'])
const userProps = { ...props.user }
function updated() {
  emit('updated')
}
function close() {
  emit('close')
}
</script>
<style lang="scss" scoped></style>
