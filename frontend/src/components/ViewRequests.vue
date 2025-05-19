<template>
  <div class="requests">
    <h2>Заявки на усыновление</h2>
    <table>
      <thead>
        <tr>
          <th>Пользователь</th>
          <th>Собака</th>
          <th>Дата</th>
          <th>Статус</th>
          <th>Действие</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="req in requests" :key="req.id">
          <td>{{ req.user.name }}</td>
          <td>{{ req.dog.name }}</td>
          <td>{{ new Date(req.date).toLocaleDateString() }}</td>
          <td>{{ req.status }}</td>
          <td>
            <select v-model="req.status" @change="update(req)">
              <option>pending</option>
              <option>approved</option>
              <option>rejected</option>
            </select>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
const requests = ref([])
async function fetchRequests() {
  const res = await fetch('/api/adoptions')
  requests.value = await res.json()
}
async function update(req) {
  await fetch(`/api/adoptions/${req.id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
  })
}
onMounted(fetchRequests)
</script>
<style lang="scss" scoped>
.requests {
  padding: 1rem;
  table {
    width: 100%;
    border-collapse: collapse;
    th,
    td {
      padding: 0.5rem;
      border: 1px solid #ddd;
    }
  }
}
</style>
