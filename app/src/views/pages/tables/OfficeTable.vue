<script setup>

const props = defineProps({
  items: Array,
  totalItems: Number,
  loading: Boolean,
})

// Pagination state
const itemsPerPage = defineModel("itemsPerPage")
const page = defineModel("page")

const headers = [
  { title: 'Nama', key: 'name' },
  { title: 'Kode', key: 'code' },
  { title: 'Email', key: 'email' },
  { title: 'Telepon', key: 'phone' },
  { title: 'Provinsi', key: 'province' },
  { title: 'Kabupaten/Kota', key: 'regency' },
  { title: 'Kecamatan', key: 'district' },
]
</script>

<template>
  <VDataTable 
    :headers="headers" 
    :items="props.items"
    :items-per-page="itemsPerPage"
    :page="page"
    :items-length="totalItems"
    :server-items-length="totalItems"
    :loading="loading"
    loading-text="Loading... Mohon Tunggu"
  />
</template>
