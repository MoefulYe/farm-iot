export interface Paginated<D> {
  data: D[]
  cnt: number
}

export interface Pagination {
    page: number
    size: number
}