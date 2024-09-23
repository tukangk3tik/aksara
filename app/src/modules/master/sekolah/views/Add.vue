<template>
  <div class="grid">
    <div class="col-12">
      <Card class="slim">
        <template #content>
          <Panel
            header="Tambah Master Data Sekolah"
            :toggleable="false"
          >
            <template #icons>
              <Button class="p-button-text p-button-info p-button-rounded p-button-raised button-sm"><span class="material-icons">help</span>
                Info</Button>
            </template>
          </Panel>
          <TabView>
            <TabPanel header="Main Info">
              <div class="grid">
                <div class="col-8 form-mode">
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">badge</span>-->
                      NPSN
                    </span>
                    <InputText
                      v-model="formData.npsn"
                      class="inputtext-sm"
                      placeholder="NPSN"
                    />
                  </div>
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">id_card</span>-->
                      Nama
                    </span>
                    <InputText
                      v-model="formData.nama"
                      class="inputtext-sm"
                      placeholder="Nama Sekolah"
                    />
                  </div>
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">grade</span>-->
                      BP
                    </span>
                    <InputText
                      v-model="formData.bp"
                      class="inputtext-sm"
                      placeholder="BP"
                    />
                  </div>
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">checklist</span>-->
                      Tipe
                    </span>
                    <Dropdown
                      v-model="formData.type"
                      :options="option_type"
                      optionLabel="name"
                      optionValue="id"
                      placeholder="Tipe"
                    />
                  </div>
                </div>
                <div class="col-4 form-mode">
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">library_books</span>-->
                      Jlh Rombel
                    </span>
                    <InputText
                      v-model="formData.jlh_rombel"
                      class="inputtext-sm"
                      placeholder="Jlh Rombel"
                    />
                  </div>
                  <div class="p-inputgroup">
                    <span class="p-inputgroup-addon">
<!--                      <span class="material-icons-outlined material-symbols-outlined">person</span>-->
                      Jlh Guru
                    </span>
                    <InputText
                      v-model="formData.jlh_guru"
                      class="inputtext-sm"
                      placeholder="Jlh Guru"
                    />
                  </div>
                </div>
              </div>
            </TabPanel>
          </TabView>
          <div class="p-grid">
            <div class="p-col-12">
              <div class="flex jc-between">
                <div class="flex-initial flex align-items-center justify-content-center m-2 px-5 py-3">
                  <Button
                    type="button"
                    class="button-raised button-sm p-button-danger px-3"
                    @click="back()"
                  >
                    <span class="material-icons-outlined material-symbols-outlined">arrow_back</span> Back
                  </Button>
                </div>
                <div class="flex-grow-1 flex align-items-center justify-content-center m-2 px-5 py-3"></div>
                <div
                  v-if="allowSave === true"
                  class="flex-initial flex align-items-right justify-content-center m-2 px-5 py-3"
                >
                  <Button
                    type="button"
                    class="button-raised button-sm p-button-info px-3"
                    @click="updateAccountData($event)"
                  >
                    <span class="material-icons-outlined material-symbols-outlined">check_circle</span> Save Data
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </template>
      </Card>
      <ConfirmPopup group="confirm_changes"></ConfirmPopup>
      <ConfirmDialog group="keep_editing"></ConfirmDialog>
    </div>
    <Dialog
      v-model:visible="displayEditorImage"
      header="Avatar Editor"
      :style="{ width: '50vw' }"
      :modal="true"
    >
      <p class="m-0">
        <Cropper @cropImage="setImageData" />
      </p>
      <template #footer>
        <Button
          label="Close"
          icon="pi pi-times"
          class="button-text"
          @click="toggleEditImageWindow"
        />
      </template>
    </Dialog>
  </div>
</template>
<script>
import Card from 'primevue/card'
import Panel from 'primevue/panel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Dropdown from 'primevue/dropdown'
import ConfirmPopup from 'primevue/confirmpopup'
import ConfirmDialog from 'primevue/confirmdialog'
import Dialog from 'primevue/dialog'
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'
import Cropper from '@/components/Cropper.vue'
import MasterSekolahService from '@/modules/master/sekolah/service'

export default {
  name: 'MasterSekolahAdd',
  components: {
    Card,
    Panel,
    InputText,
    Button,
    Dropdown,
    ConfirmPopup,
    ConfirmDialog,
    Cropper,
    Dialog,
    TabView,
    TabPanel,
  },
  data() {
    return {
      displayEditorImage: false,
      formData: {
        npsn: '',
        nama: '',
        type: '',
        status: '',
        bp: '',
        jlh_guru: 0,
        jlh_rombel: 0,
        __v: 0,
      },
      allowSave: false,
      option_type: [{
        id: 'negeri',
        name: 'Negeri'
      }, {
        id: 'swasta',
        name: 'Swasta'
      }],
      lazyParams: {},
      selectedNode: {},
      filtersNode: {
        label: { value: '', matchMode: 'contains' },
        to: { value: '', matchMode: 'contains' },
      },
      expandedKeys: {},
      nodes: null,
      columns: [
        { field: 'label', header: 'Label', expander: true },
        { field: 'to', header: 'To' },
      ],
    }
  },
  async mounted() {
    this.allowSave = true
  },
  methods: {
    back() {
      this.$router.push('/master/sekolah')
    },
    setImageData(value) {
      this.formData.image_edit = true
      this.formData.image = value
    },
    toggleEditImageWindow() {
      this.displayEditorImage = !this.displayEditorImage
    },
    resetForm() {
      this.formData = {
        npsn: '',
        nama: '',
        type: '',
        status: '',
        bp: '',
        jlh_guru: 0,
        jlh_rombel: 0,
        __v: 0,
      }
    },
    updateAccountData: function (event) {
      const target = event.target

      const confirmation = this.$confirm
      if (this.allowSave) {
        confirmation.require({
          group: 'confirm_changes',
          target: target,
          message: `Tambah data sekolah?`,
          icon: 'pi pi-exclamation-triangle',
          acceptClass: 'button-success',
          acceptIcon: 'pi pi-check-circle',
          acceptLabel: 'Ya',
          rejectLabel: 'Batal',
          rejectIcon: 'pi pi-times-circle',
          accept: async () => {
            await MasterSekolahService.addSekolah(this.formData)
              .then(async (response) => {
                this.$confirm.require({
                  group: 'keep_editing',
                  message: `${response.message}. Tambah data sekolah yang lain?`,
                  header: 'Tambah kembali data sekolah',
                  icon: 'pi pi-exclamation-triangle',
                  acceptClass: 'p-button-success',
                  rejectClass: 'p-button-warning',
                  acceptLabel: 'Ya',
                  acceptIcon: 'pi pi-check-circle',
                  rejectLabel: 'Sudah selesai',
                  rejectIcon: 'pi pi-times-circle',
                  accept: () => {
                    this.resetForm()
                  },
                  reject: () => {
                    this.$router.push('/master/sekolah')
                  },
                  onHide: () => {
                    //Callback to execute when dialog is hidden
                  },
                })
              })
          },
          reject: () => {
            // callback to execute when user rejects the action
          },
        })
      }
    },
  },
}
</script>
