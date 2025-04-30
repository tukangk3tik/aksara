export interface NavItem {
  id: string;
  title: string;
  path?: string;
  icon?: string;
  children?: NavItem[];
}

export interface BreadcrumbItem {
  title: string;
  path?: string;
}