import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

const { API_URL, S3_HOST, S3_BUCKET } = publicRuntimeConfig;

export const ENVIRONMENT = {
  api_url: API_URL || "http://localhost:1323",
  s3_url: S3_HOST ? `${S3_HOST}/${S3_BUCKET}/` : "http://localhost:9000/originals/",
};

console.log("env:", ENVIRONMENT);