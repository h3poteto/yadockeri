export type Project = {
  id: number
  user_id: number
  title: string
  base_url: string
  repository_owner: string
  repository_name: string
  helm_repository_url: string
  helm_directory_name: string
  namespace: string
  values: Array<OverrideValue>
}

export type OverrideValue = {
  key: string
  value: string
}
