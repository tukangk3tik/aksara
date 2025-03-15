<script setup>
const offices = [
  {
    id: 1,
    code: 'KP01',
    name: 'Kantor Pusat',
    email: 'kantor_pusat@cerdascek.com',
    phone: '+62',
    province: 'DKI JAKARTA',
    regency: 'KOTA ADM. JAKARTA PUSAT',
    district: 'Kemayoran',
  },
  {
    id: 2,
    code: 'KDB01',
    name: 'Kantor Kab. Belu',
    email: 'kd_belu01@cerdascek.com',
    phone: '+62',
    province: 'Nusa Tenggara Timur',
    regency: 'Kab. Belu',
    district: 'Kota Atambua',
  },
]

const headers = [
  { title: 'Nama', key: 'name' },
  { title: 'Kode', key: 'code' },
  { title: 'Email', key: 'email' },
  { title: 'Telepon', key: 'phone' },
  { title: 'Provinsi', key: 'province' },
  { title: 'Kabupaten/Kota', key: 'regency' },
  { title: 'Kecamatan', key: 'district' },
];

const totalItem = ref(0)
const itemPerPage = ref(10)
const loading = ref(false)

// need to call the API
</script>

<template>
  <VDataTable 
    :headers="headers" 
    :items="offices"
    :loading="loading"
    loading-text="Loading... Mohon Tunggu"

    />
</template>
