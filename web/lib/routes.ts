import { param, route } from "typesafe-react-router";

export const APP_ROUTES = {
  home: route(),
  signUp: route("sign_up"),
  admin: {
    dashboard: route("admin", "dashboard"),
    photos: {
      index: route("admin", "photos"),
      upload: route("admin", "photos", "upload"),
    },
  },
  auth: {
    login: route("login"),
    logout: route("logout"),
  },
};

export const API_ROUTES = {
  photos: {
    index: route("photos", "index"),
    create: route("photos", "create"),
    add_tags: route("admin", "photos", param("photoId"), "tags"),
    remove_tag: route("admin", "photos", param("photoId"), "tags", param("tagId")),
    upload_photo: route("admin", "photos", "user", param("userId")),
  },
  login: route("login"),
  sign_up: route("sign_up"),
  sign_up_confirmation: route("sign_up_confirmation")
};
