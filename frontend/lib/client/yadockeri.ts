import axios, { AxiosResponse, AxiosError } from 'axios'
import { AuthenticationError } from './error'

/**
 * Yadockeri API client
 */
export default class Yadockeri {
  public static async get<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.get<T>(url, params).catch((err: AxiosError) => {
      if (err.response!.status === 401) {
        throw new AuthenticationError(err.message)
      } else {
        throw err
      }
    })
  }

  public static async put<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.put<T>(url, params).catch((err: AxiosError) => {
      if (err.response!.status === 401) {
        throw new AuthenticationError(err.message)
      } else {
        throw err
      }
    })
  }

  public static async patch<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.patch<T>(url, params).catch((err: AxiosError) => {
      if (err.response!.status === 401) {
        throw new AuthenticationError(err.message)
      } else {
        throw err
      }
    })
  }

  public static async post<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.post<T>(url, params).catch((err: AxiosError) => {
      if (err.response!.status === 401) {
        throw new AuthenticationError(err.message)
      } else {
        throw err
      }
    })
  }

  public static async delete<T>(
    url: string,
    params = {}
  ): Promise<AxiosResponse<T>> {
    return axios.delete<T>(url, params).catch((err: AxiosError) => {
      if (err.response!.status === 401) {
        throw new AuthenticationError(err.message)
      } else {
        throw err
      }
    })
  }
}
