export const APP_ROUTES = {
  home: "/",
  signUp: "/sign_up",
  admin: {
    dashboard: "/admin/dashboard",
    photos: {
      index: "/admin/photos",
      upload: "/admin/photos/upload",
    },
  },
  auth: {
    login: "/login",
    logout: "/logout",
  },
};

export const API_ROUTES = {
  photos: {
    index: "/photos/index",
    create: "/photos/create",
  },
};
