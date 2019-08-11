import axios, { AxiosResponse } from 'axios'

/**
 * Interface for API client
 */
export interface YadockeriInstance {
  get<T = any>(url: string, params: object): Promise<AxiosResponse<T>>
  put<T = any>(url: string, params: object): Promise<AxiosResponse<T>>
  patch<T = any>(url: string, params: object): Promise<AxiosResponse<T>>
  post<T = any>(url: string, params: object): Promise<AxiosResponse<T>>
  delete<T = any>(url: string, params: object): Promise<AxiosResponse<T>>
}

/**
 * Yadockeri API client
 */
export default class Yadockeri implements YadockeriInstance {
  public async get<T>(url: string, params = {}): Promise<AxiosResponse<T>> {
    return axios.get<T>(url, params)
  }

  public async put<T>(url: string, params = {}): Promise<AxiosResponse<T>> {
    return axios.put<T>(url, params)
  }

  public async patch<T>(url: string, params = {}): Promise<AxiosResponse<T>> {
    return axios.patch<T>(url, params)
  }

  public async post<T>(url: string, params = {}): Promise<AxiosResponse<T>> {
    return axios.post<T>(url, params)
  }

  public async delete<T>(url: string, params = {}): Promise<AxiosResponse<T>> {
    return axios.delete<T>(url, params)
  }
}
