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
        title: 'Office',
        path: '/master/office',
      },
      {
        id: 'school',
        title: 'School',
        path: '/master/school',
      },
      // {
      //   id: 'student',
      //   title: 'Student',
      //   path: '/master/student',
      // },
      {
        id: 'teacher',
        title: 'Teacher',
        path: '/master/teacher',
      },
    ],
  },
  {
    id: 'users',
    title: 'Users',
    path: '/users',
    icon: 'UserGroupIcon',
  },
  {
    id: 'settings',
    title: 'Settings',
    path: '/settings',
    icon: 'CogIcon',
  },
];