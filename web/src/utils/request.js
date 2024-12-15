import axios from 'axios'
import { ElLoading, ElMessage } from 'element-plus'
import { useAdminUserStore } from '@/pinia/modules/admin_user.js' // 引入axios
const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API, timeout: 99999
})

let activeAxios = 0
let timer
let loadingInstance

const showLoading = (option) => {
  activeAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      loadingInstance = ElLoading.service(option)
    }
  }, 400)
}

const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    clearTimeout(timer)
    loadingInstance && loadingInstance.close()
  }
}

service.interceptors.request.use((config) => {
  if (!config.donNotShowLoading) {
    showLoading(config.loadingOption)
  }

  config.headers = {
    'Content-Type': 'application/json', 'Auth': useAdminUserStore().token, ...config.headers
  }
  return config
}, (error) => {
  if (!error.config.donNotShowLoading) {
    closeLoading()
  }
  ElMessage({
    showClose: true, message: error, type: 'error'
  })
  return Promise.reject(error)
})

// http response 拦截器
service.interceptors.response.use((response) => {
  if (!response.config.donNotShowLoading) {
    closeLoading()
  }

  if (response.data.code !== 0) {
    ElMessage({
      showClose: true, message: response.data.message, type: 'error'
    })
    return Promise.reject(response.data)
  }

  return Promise.resolve(response.data.data)
}, (error) => {
  if (!error.config.donNotShowLoading) {
    closeLoading()
  }
  return Promise.reject(error)
})

export default service