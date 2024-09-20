const moduleRoute = [
  {
    path: '/master/sekolah',
    name: 'MasterSekolah',
    component: () =>
      import(
        /* webpackChunkName: "master_item" */ '@/modules/master/sekolah/Module.vue'
      ),
    meta: {
      pageTitle: 'Master Data Sekolah',
      requiresAuth: true,
      breadcrumb: [
        {
          label: 'Master Sekolah',
          to: '/master/sekolah',
        },
      ],
    },
    children: [
      {
        path: '',
        name: 'MasterSekolah',
        meta: {
          pageTitle: 'Master Data Sekolah',
          requiresAuth: true,
          breadcrumb: [
            {
              label: 'Master Data Sekolah',
              to: '/master/sekolah/',
            },
          ],
        },
        component: () =>
          import(
            /* webpackChunkName: "master_item" */ '@/modules/master/sekolah/views/Index.vue'
          ),
      },
      {
        path: 'add',
        name: 'MasterSekolahAdd',
        meta: {
          pageTitle: 'Tambah Master Sekolah',
          requiresAuth: true,
          breadcrumb: [
            {
              label: 'Master Sekolah',
              to: '/master/sekolah/',
            },
          ],
        },
        component: () =>
          import(
            /* webpackChunkName: "master_item" */ '@/modules/master/sekolah/views/Add.vue'
          ),
      },
      {
        path: 'edit/:id',
        name: 'MasterSekolahEdit',
        meta: {
          pageTitle: 'Edit Master Sekolah',
          requiresAuth: true,
          breadcrumb: [
            {
              label: 'Master Sekolah',
              to: '/master/sekolah/',
            },
          ],
        },
        component: () =>
          import(
            /* webpackChunkName: "master_item" */ '@/modules/master/sekolah/views/Edit.vue'
          ),
      },
    ],
  },
]

export default moduleRoute
