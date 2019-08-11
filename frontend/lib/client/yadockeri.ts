import axios, { AxiosResponse } from 'axios'

/**
 * Yadockeri API client
 */
export default class Yadockeri {
  public static async get<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.get<T>(url, params)
  }

  public static async put<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.put<T>(url, params)
  }

  public static async patch<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.patch<T>(url, params)
  }

  public static async post<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.post<T>(url, params)
  }

  public static async delete<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.delete<T>(url, params)
  }
}
