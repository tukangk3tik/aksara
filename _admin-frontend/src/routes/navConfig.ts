import { NavItem } from '../types/navigation';

export const navigationItems: NavItem[] = [
  {
    id: 'dashboard',
    title: 'Dashboard',
    path: '/',
    icon: 'HomeIcon',
  },
  {
    id: 'master',
    title: 'Master Data',
    icon: 'DatabaseIcon',
    children: [
      {
        id: 'offices',
        title: 'Kantor',
        path: '/master/office',
      },
      {
        id: 'school',
        title: 'Sekolah',
        path: '/master/school',
      },
      {
        id: 'teacher',
        title: 'Guru',
        path: '/master/teacher',
      },
    ],
  },
  {
    id: 'users',
    title: 'Pengguna',
    path: '/users',
    icon: 'UserGroupIcon',
  },
  {
    id: 'settings',
    title: 'Pengaturan',
    path: '/settings',
    icon: 'CogIcon',
  },
];