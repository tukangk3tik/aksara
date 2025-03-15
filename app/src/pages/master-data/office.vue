<script setup>
// import OfficeTable from '@/views/pages/tables/OfficeTable.vue';
import { useOfficeStore } from '@/stores/office.store';
import { storeToRefs } from 'pinia';
import { onMounted, watch } from 'vue'

const officeStore = useOfficeStore()
const { isLoading, offices, totalItems, limit, page } = storeToRefs(officeStore)

onMounted(() => {
  officeStore.fetchOffices()
})

// Watch for pagination changes and fetch new data
watch([page, limit], () => {
  officeStore.fetchOffices();
});

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
  <VRow>
    <VCol cols="12">
      <VCard title="Daftar Kantor Wilayah">
        <VDataTable 
          :headers="headers" 
          :loading="isLoading"
          :items="offices"
          :items-per-page="limit"
          :page="page"
          :items-length="totalItems"
          :server-items-length="totalItems"
          loading-text="Loading... Mohon Tunggu"
          @update:page="page = $event"
          @update:items-per-page="limit = $event"
        />
      </VCard>
    </VCol>
  </VRow>
</template>
