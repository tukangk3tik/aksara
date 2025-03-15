
export class DataTableDto {
  constructor({page, limit, search, sort}) {
    this.page = page ?? 1;
    this.limit = limit ?? 10;
    this.search = search ?? '';
    this.sort = sort ?? '';
  }

  toQueryParams() {
    return `page=${this.page}&limit=${this.limit}&search=${this.search}&sort=${this.sort}`;
  }

  getPagination() {
    return {
      page: this.page,
      limit: this.limit,
      search: this.search,
      sort: this.sort
    }
  }
}