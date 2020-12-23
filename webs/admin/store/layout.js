export const state = () => ({
  pageInfo: {
    show: true,
    title: '',
    subtitle: '',
    icon: 'mdi-google-circles',
    backTo: '',
    breadcrumbs: [],
  },
  drawer: {
    show: true,
    dense: false,
    clipped: true,
    fixed: true,
    floating: false,
    mini: false,
  },
  menu: {
    items: [
      {
        module: 'dashboard',
        menuID: 'dashboard',
        icon: 'mdi-view-dashboard',
        title: 'Dashboard',
        subtitle: 'Monitor event',
        to: '/',
        user_kind: [
          'tenant-sparepart',
          'tenant-automotive',
          'eventorganizer',
          'appowner',
        ],
      },
      {
        module: 'event',
        menuID: 'event',
        icon: 'mdi-android-auto',
        title: 'Event',
        subtitle: 'Event management',
        to: '/event',
        user_kind: ['eventorganizer', 'appowner'],
      },
      {
        module: 'content',
        menuID: 'content',
        icon: 'mdi-newspaper',
        title: 'Content',
        subtitle: 'Content management',
        to: '/content',
        user_kind: ['eventorganizer', 'appowner'],
      },
      {
        module: 'userinbox',
        menuID: 'userinbox',
        icon: 'mdi-tray-full',
        title: 'User Inbox',
        subtitle: 'User Inbox',
        to: '/userinbox',
        user_kind: ['eventorganizer', 'appowner'],
      },
      {
        title: 'Master',
        module: 'master',
        user_kind: ['eventorganizer', 'appowner'],
        children: [
          {
            module: 'master',
            submodule: 'tenant',
            menuID: 'master-tenant',
            icon: 'mdi-city',
            title: 'Tenant',
            subtitle: 'Tenant management',
            to: '/master/tenant',
            user_kind: ['eventorganizer', 'appowner'],
          },
        ],
      },
      {
        title: 'Access',
        module: 'account',
        user_kind: [
          'tenant-sparepart',
          'tenant-automotive',
          'eventorganizer',
          'appowner',
        ],
        children: [
          {
            module: 'account',
            menuID: 'account',
            icon: 'mdi-account-supervisor',
            title: 'Account',
            subtitle: 'User account management',
            to: '/account/profile',
            user_kind: [
              'tenant-sparepart',
              'tenant-automotive',
              'eventorganizer',
              'appowner',
            ],
          },
        ],
      },
    ],
  },
})

export const mutations = {

}

export const actions = {

}
