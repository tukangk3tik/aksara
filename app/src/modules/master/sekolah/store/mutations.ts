import * as types from './types'

export default {
  [types.MASTER_SEKOLAH_LIST] (state: any) {
    return state.items
  },

  [types.MASTER_SEKOLAH_ADD] (state: any, item: any) {
    state.items.push(item)
  },

  [types.MASTER_SEKOLAH_REMOVE] (state: any, id: any) {
    state.items = state.items.filter((item: any) => item.id !== id)
  }
}
