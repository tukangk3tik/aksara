import api from '@/util/api'
import {ISupplier} from "@/model/Supplier";
import {CoreResponse} from "@/model/Response";
import {AxiosResponse} from "axios";
import process from 'process'

class MasterSekolahService {
  getSekolahList(parsedData) {
    return api({ requiresAuth: true })
      .get(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah`, {
        params: {
          lazyEvent: JSON.stringify(parsedData)
        },
      })
      .then((response: any) => {
        const data:CoreResponse = response.data
        return data
      })
      .catch((e) => {
        return Promise.reject(e)
      })
  }

  async getSekolahDetail(id) {
    return await api({ requiresAuth: true })
      .get(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah/${id}`)
      .then((response: any) => {
        return Promise.resolve(response)
      })
  }

  async findSekolah(search: string): Promise<ISupplier[]> {
    return await api({requiresAuth: true}).get(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah`, {
      params: {
        lazyEvent: `{"first":0,"rows":10,"projection": {}, "sortField":"created_at","sortOrder":1,"filters":{"name":{"value":"${search}", "matchMode": "contains"}},"search_term": {}}`,
      }
    }).then((response: AxiosResponse) => {
      return Promise.resolve(response.data)
    })
  }

  async addSekolah(data) {
    return await api({requiresAuth: true}).post(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah`, data)
      .then((response: any) => {
      return Promise.resolve(response.data)
    })
      .catch((e) => {
      return Promise.reject(e)
    })
  }

  async editSekolah(id, data) {
    return await api({requiresAuth: true}).patch(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah/${id}`, data)
      .then((response: any) => {
        return Promise.resolve(response.data)
      })
      .catch((e) => {
        return Promise.reject(e)
      })
  }

  async deleteSekolah(id, data) {
    return await api({requiresAuth: true}).delete(`${process.env.VUE_APP_APIGATEWAY}/v1/master/sekolah/${id}`)
      .then((response: any) => {
        return Promise.resolve(response.data)
      })
      .catch((e) => {
        return Promise.reject(e)
      })
  }
}

export default new MasterSekolahService()
