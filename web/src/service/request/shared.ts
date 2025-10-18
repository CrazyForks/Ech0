import { localStg } from '@/utils/storage'

export const getAuthToken = () => {
  const token = localStg.getItem<string>('token')
  return token ? `Bearer ${token}` : ''
}

export const saveAuthToken = (token: string) => {
  if (token) {
    localStg.setItem('token', token)
  }
}

export const getApiUrl = () => {
  const baseUrl = import.meta.env.VITE_SERVICE_BASE_URL
  const resolvedBaseUrl = baseUrl.replace(/\/+$/, '') // 正则去除末尾的斜杠

  // 检查是否使用正向代理
  if (import.meta.env.VITE_PROXY === 'YES') {
    // BaseURL + ProxyURL
    const proxyUrl = import.meta.env.VITE_PROXY_URL
    if (!proxyUrl) {
      throw new Error('Proxy URL is not defined')
    }
    return `${resolvedBaseUrl}${proxyUrl}`
  }
  return resolvedBaseUrl
}

export const getSystemReadyStatus = () => {
  // 检查localStorage中是否有系统状态
  const systemStatus = localStg.getItem<boolean>('systemStatus')
  if (systemStatus !== null) {
    // 如果有，直接使用localStorage中的值
    return systemStatus
  } else {
    // 如果没有，默认设置为false
    return false
  }
}

// src/utils/ws.ts
export function getWsUrl(path: string) {
  // 取出基础地址
  const baseUrl = import.meta.env.VITE_SERVICE_BASE_URL

  // 根据当前协议选择 ws 或 wss
  const wsProtocol = location.protocol === 'https:' ? 'wss:' : 'ws:'

  // 如果是相对路径（生产环境配置为 "/"），自动拼上当前域名
  if (baseUrl === '/' || baseUrl.startsWith('/')) {
    return `${wsProtocol}//${location.host}${path}`
  }

  // 否则使用开发环境配置的完整 baseUrl
  const httpUrl = new URL(baseUrl)
  return `${wsProtocol}//${httpUrl.host}${path}`
}
