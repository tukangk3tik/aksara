<template>
  <div class="grid">
    <div class="col-12">
      <Card>
        <template #title>Tambah Sekolah</template>
        <template #content>
          <div class="grid">
            <div class="col-4">

            </div>
            <div class="col-8 form-mode">
              <div class="inputgroup">
                <span class="inputgroup-addon">
                  <span class="material-icons-outlined material-symbols-outlined">mail</span>
                </span>
                <!-- <InputText class="inputtext-sm" @input="updateAccount($event.target.value)" v-model="accountDetail.email"
                  placeholder="Email" /> -->
                <InputText
                  v-model="formData.email"
                  class="inputtext-sm"
                  placeholder="Email"
                />
              </div>
              <div class="inputgroup">
                <span class="inputgroup-addon">
                  <span class="material-icons-outlined material-symbols-outlined">person</span>
                </span>
                <InputText
                  v-model="formData.first_name"
                  class="inputtext-sm"
                  placeholder="First Name"
                />
                <InputText
                  v-model="formData.last_name"
                  class="inputtext-sm"
                  placeholder="Last Name"
                />
              </div>
              <div class="inputgroup">
                <span class="inputgroup-addon">
                  <span class="material-icons-outlined material-symbols-outlined">supervised_user_circle</span>
                </span>
                <Dropdown
                  v-model="formData.authority"
                  :options="authorityData"
                  optionLabel="name"
                  optionValue="id"
                  placeholder="Select authority"
                />
              </div>
              <Button
                class="button button-info button-sm button-raised"
                @click="updateAccountData"
              >
                <span class="material-icons">fact_check</span> Apply from authority
              </Button>
            </div>
            <div class="col-12">
              <div class="d-flex jc-between">
                <div>
                  <Button
                    type="button"
                    class="button-raised button-sm p-button-danger px-3"
                    @click="back()"
                  >
                    <span class="material-icons-outlined material-symbols-outlined">arrow_back</span> Back
                  </Button>
                </div>
                <div v-if="allowSave === true">
                  <Button
                    type="button"
                    class="button-raised button-sm button-info px-3"
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
      <ConfirmPopup></ConfirmPopup>
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
import InputText from 'primevue/inputtext'
import Dropdown from 'primevue/dropdown'
import Button from 'primevue/button'
import Checkbox from 'primevue/checkbox'
import Column from 'primevue/column'
import ConfirmPopup from 'primevue/confirmpopup'
import Dialog from 'primevue/dialog'
import Cropper from '@/components/Cropper.vue'
import { getCurrentTimestamp } from '@/util/time'

import { mapActions, mapGetters, mapState } from 'vuex'
export default {
  name: 'MasterSekolahAdd',
  components: {
    Card,
    InputText,
    Button,
    Dropdown,
    ConfirmPopup,
    Checkbox,
    Column,
    Cropper,
    Dialog,
  },
  data() {
    return {
      displayEditorImage: false,
      formData: {
        id: 0,
        email: '',
        authority: '',
        first_name: '',
        last_name: '',
        image: '',
        menuTree: [],
      },
      allowSave: false,
      authorityData: [],
      selectedParent: {},
      selectedMenu: [],
      selectedPage: [],
      selectedPerm: [],
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
  computed: {
    // ...mapState('accountModule', ['menu_list', 'menu_tree']),
    // ...mapGetters({
    //   menuTree: 'accountModule/getMenuTree',
    //   accountDetail: 'accountModule/getAccountDetail',
    //   authorityList: 'accountModule/getAuthorityList',
    // }),
  },
  // watch: {
  //   accountDetail: {
  //     handler(getDetail) {
  //       if (getDetail) {
  //         //
  //       } else {
  //         this.allowSave = false
  //       }
  //     },
  //     immediate: true,
  //   },
  // },
  async mounted() {
    this.allowSave = false
    // this.$store.dispatch('accountModule/fetchMenu')
    // await this.$store.dispatch(
    //   'accountModule/fetchAccountDetail',
    //   this.$route.query.id
    // )
    // await this.$store.dispatch('accountModule/fetchMenuTree')
    // await this.$store.dispatch('accountModule/fetchAuthority')
    this.displayEditorImage = false
  },
  methods: {
    ...mapActions('accountModule', [
      'updateAccount',
      'updatePermission',
      'updateAccess',
    ]),
    back() {
      this.$router.push('/master/sekolah')
    },
    updateAccountData: function (event) {
      if (this.allowSave) {
        // this.formData.image = this.formData.image.replace(/^data:image\/\w+;base64,/, '')
        this.updateAccount(this.formData).then((response) => {
          if (response.status === 200) {
            const responseAccess = []
            const responsePermission = []
            // // Update Access and Permission if exists
            for (const a in this.selectedMenu) {
              this.updateAccess({
                account: this.$route.query.id,
                menu: this.selectedMenu[a],
              }).then((response) => {
                responseAccess.push(response)
              })
            }

            for (const a in this.selectedPerm) {
              this.updatePermission({
                account: this.$route.query.id,
                permission: this.selectedPerm[a],
              }).then((response) => {
                responsePermission.push(response)
              })
            }

            this.$store.dispatch('accountModule/fetchMenuTree', this.lazyParams)
            this.$confirm.require({
              target: event.target,
              message: `${response.data.message}. Back to account list?`,
              icon: 'pi pi-exclamation-triangle',
              acceptClass: 'button-success',
              acceptLabel: 'Yes',
              rejectLabel: 'Keep Editing',
              accept: () => {
                this.$router.push('/account')
              },
              reject: () => {
                // callback to execute when user rejects the action
              },
            })
          }
        })
      }
    },
  },
}
</script>
<style>
.profile-display {
  position: relative;
  background: red;
}

.profile-display img {
  position: absolute;
  border-radius: 100%;
  width: 200px;
  height: 200px;
  background: #f2f2f2;
  margin: 0 25%;
}
</style>
