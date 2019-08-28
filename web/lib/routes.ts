const route = (path: string) => {
  const template = path;
  const create = (obj?: any) => {
    let result = template;
    if (!obj) return result;
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        result = result.replace(`:${key}`, obj[key]);
      }
    }
    return result;
  };
  return { template, create };
};

export const APP_ROUTES = {
  home: route("/"),
  signUp: route("/sign_up"),
  admin: {
    dashboard: route("/admin/dashboard"),
    photos: {
      index: route("/admin/photos"),
      upload: route("/admin/photos/upload"),
      show: route("/admin/photos/:photoId"),

    },
  },
  auth: {
    login: route("/login"),
    logout: route("/logout"),
  },
};

export const API_ROUTES = {
  photos: {
    index: route("/photos/index"),
    create: route("/photos/create"),
    add_tags: route("/admin/photos/:photoId/tags"),
    add_apps: route("/admin/photos/:photoId/apps"),
    remove_tag: route("/admin/photos/:photoId/tags/:tagId"),
    remove_app: route("/admin/photos/:photoId/apps/:appId"),
    upload_photo: route("/admin/photos/user/:userId"),
  },
  login: route("/login"),
  sign_up: route("/sign_up"),
  sign_up_confirmation: route("/sign_up_confirmation"),
};
