<template>
  <div>
    <Card class="slim">
      <template #content>
        <Panel
          header="Manajemen Master Sekolah"
          :toggleable="false"
        >
          <template #icons>
            <Button
              class="p-button-info p-button-rounded p-button-raised button-sm"
              @click="itemAddForm"
            ><span class="material-icons">add</span>
              Tambah Sekolah</Button>
          </template>
          <DataTable
            ref="dt"
            v-model:filters="filters"
            :value="items"
            :lazy="true"
            :paginator="true"
            :rows="20"
            :rowsPerPageOptions="[5, 10, 20, 50]"
            :totalRecords="totalRecords"
            :loading="loading"
            filterDisplay="row"
            :globalFilterFields="['code', 'name']"
            responsiveLayout="scroll"
            @page="onPage($event)"
            @sort="onSort($event)"
            @filter="onFilter($event)"
          >
            <Column
              header="#"
              class="text-right wrap_content"
            >
              <template #body="slotProps">
                <strong class="d-inline-flex">
                  <span class="material-icons-outlined material-symbols-outlined">tag</span>
                  {{ slotProps.data.autonum}}
                </strong>
              </template>
            </Column>
            <Column
              header="Action"
              class="text-right wrap_content">
              <template #body="slotProps">
              <span class="p-buttonset">
                <Button
                  class="p-button-info p-button-sm p-button-raised"
                  @click="itemEditForm(slotProps.data.id)"
                >
                  <span class="material-icons">edit</span> Edit
                </Button>
                <Button
                  class="p-button-danger p-button-sm p-button-raised"
                  @click="itemDelete($event, slotProps.data.id)"
                >
                  <span class="material-icons">delete</span> Delete
                </Button>
              </span>
              </template>
            </Column>
            <Column
              ref="npsn"
              field="npsn"
              header="NPSN"
              filterMatchMode="startsWith"
              :sortable="true"
              style="width: 12.5% !important;"
            >
              <!--              eslint-disable-next-line -->
              <template #filter="{ filterModel, filterCallback }">
                <InputText
                  type="text"
                  class="column-filter"
                  placeholder="Cari NPSN"
                  @keydown.enter="filterCallback()"
                />
              </template>
              <template #body="slotProps">
                <label class="currency-label text-600">{{ slotProps.data.npsn }}</label>
              </template>
            </Column>
            <Column
              ref="name"
              field="name"
              header="Nama"
              filterField="name"
              filterMatchMode="contains"
              :sortable="true"
            >
              <!--              eslint-disable-next-line -->
              <template #filter="{ filterModel, filterCallback }">
                <InputText
                  type="text"
                  class="column-filter"
                  placeholder="Cari Nama"
                  @keydown.enter="filterCallback()"
                />
              </template>
              <template #body="slotProps">
                <span class="text-cyan-500">{{ slotProps.data?.bp }} {{ slotProps.data?.status }}</span><br />
                <b class="font-bold">{{ slotProps.data?.name ?? '-' }}</b>
              </template>
            </Column>
            <Column
              ref="jlh_siswa"
              field="jlh_siswa"
              header="Jlh Siswa"
              :sortable="false"
              style="width: 12.5% !important;"
            >
              <template #body="slotProps">
                <b>{{ slotProps.data?.jlh_siswa ?? 0 }}</b>
              </template>
            </Column>
            <Column
              ref="rombel"
              field="rombel"
              header="Rombel"
              :sortable="false"
              style="width: 12.5% !important;"
            >
              <template #body="slotProps">
                <b>{{ slotProps.data?.rombel ?? 0 }}</b>
              </template>
            </Column>
            <Column
              ref="guru"
              field="guru"
              header="Guru"
              :sortable="false"
              style="width: 12.5% !important;"
            >
              <template #body="slotProps">
                <b>{{ slotProps.data?.guru ?? 0 }}</b>
              </template>
            </Column>
            <Column
              ref="created_at"
              class="text-right wrap_content"
              field="created_at"
              header="Created Date"
              :sortable="true"
            >
              <template #body="slotProps">
                <b>{{ formatDate(slotProps.data.created_at, 'DD MMMM YYYY') }}</b>
              </template>
            </Column>
            <template #footer>
              <div class="text-xs">
                <NumberLabel class="text-cyan-600" lang="ID" code="ID" currency="IDR" :number="lazyParams.page > 0 ? lazyParams.first + 1 : 1" decimal="0" /> - <NumberLabel class="text-cyan-600" lang="ID" code="ID" currency="IDR" :number="lazyParams.first > 0 ? lazyParams.first + lazyParams.rows : lazyParams.rows" decimal="0" /> / <NumberLabel class="text-cyan-600" lang="ID" code="ID" currency="IDR" :number="totalRecords ? totalRecords : 0" decimal="0" /> rows
              </div>
            </template>
          </DataTable>
        </Panel>
      </template>
      <template #footer></template>
    </Card>
    <ConfirmPopup></ConfirmPopup>
  </div>
</template>

<script>
/* eslint-disable */
import DateManagement from '@/modules/function'
import Card from 'primevue/card'
import Panel from 'primevue/panel'
import ConfirmPopup from 'primevue/confirmpopup'
import Toolbar from 'primevue/toolbar'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import MasterSekolahService from '@/modules/master/sekolah/service'
import DataTableFilterMeta from 'primevue/datatable'
import {mapState} from "vuex"
import NumberLabel from '@/components/Number.vue'

export default {
  components: {
    DataTable,
    Panel,
    Column,
    InputText,
    Button,
    Card,
    Toolbar,
    ConfirmPopup,
    NumberLabel,
  },
  data() {
    return {
      loading: false,
      totalRecords: 0,
      items: [],
      filters: {
        code: { value: '', matchMode: 'contains' },
        name: { value: '', matchMode: 'contains' },
      },
      lazyParams: {},
      columns: [
        { field: 'code', header: 'Code' },
        { field: 'name', header: 'Name' },
        { field: 'created_at', header: 'Join Date' },
      ],
    }
  },
  computed: {
    ...mapState('storeCredential', {
      credential: state => state
    }),
  },
  mounted() {
    this.lazyParams = {
      first: 0,
      page: 0,
      rows: this.$refs.dt.rows,
      sortField: 'created_at',
      sortOrder: 1,
      filters: this.filters,
    }

    this.loadLazyData()
  },
  methods: {
    itemAddForm() {
      this.$router.push('/master/sekolah/add')
    },
    itemEditForm(id) {
      // this.$router.push(`/master/sekolah/edit/${id}`)
      this.$router.push({
        path: `/master/sekolah/edit/${id}`,
        query: {
          id: id,
        },
      })
    },
    itemDelete(event, id) {
      this.$confirm.require({
        target: event.currentTarget,
        message: 'Hapus data sekolah ini?',
        icon: 'pi pi-exclamation-triangle',
        acceptClass: 'button-danger',
        acceptLabel: 'Ya. Hapus',
        rejectLabel: 'Batal',
        accept: async () => {
          this.loading = true
          await MasterSekolahService.deleteSekolah(id).then((detail) => {
            this.loading = false
          })
        },
        reject: () => {
          // Reject
        },
      })
    },
    formatDate(date, format) {
      return DateManagement.formatDate(date, format)
    },
    loadLazyData() {
      this.loading = true

      MasterSekolahService.getSekolahList(this.lazyParams).then((response) => {
        if (response) {
          const data = response.payload.data
          const totalRecords = response.payload.totalRecords
          this.items = data
          this.totalRecords = totalRecords
        } else {
          this.items = []
          this.totalRecords = 0
        }
        this.loading = false
      })
    },
    onPage(event) {
      this.lazyParams = event
      this.loadLazyData()
    },
    onSort(event) {
      this.lazyParams = event
      this.loadLazyData()
    },
    onFilter() {
      this.lazyParams.page = 0
      this.lazyParams.first = 0
      this.lazyParams.filters = this.filters
      this.loadLazyData()
    },
  },
}
</script>
