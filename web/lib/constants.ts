import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

const { API_URL } = publicRuntimeConfig;

export const ENVIRONMENT = {
  api_url: API_URL || "http://localhost:1323",
};

console.log("env:", ENVIRONMENT);