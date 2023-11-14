export interface Paginated<D> {
  data: D[]
  total: number
}

export interface Pagination {
    page: number
    size: number
}