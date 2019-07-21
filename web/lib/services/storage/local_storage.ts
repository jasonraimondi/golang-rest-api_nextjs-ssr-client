import { COOKIES } from "../../cookie";

const storagePrefix = "app.";

export const LOCAL_STORAGE = {
  authToken: COOKIES.authToken,
};

export class LocalStorage {
  public static remove(key: string): void {
    localStorage.removeItem(storagePrefix + key);
  }

  public static get<T>(key: string): T | null {
    const item = localStorage.getItem(storagePrefix + key);

    if (!item || item === "null") {
      return null;
    }

    try {
      return JSON.parse(item);
    } catch (err) {
      return null;
    }
  }

  public static set(key: string, value: any): boolean {
    if (value === undefined) {
      value = null;
    } else {
      try {
        value = JSON.stringify(value);
      } catch (err) {
        return false;
      }
    }

    try {
      localStorage.setItem(storagePrefix + key, value);
      return true;
    } catch (err) {
      return false;
    }
  }
}
