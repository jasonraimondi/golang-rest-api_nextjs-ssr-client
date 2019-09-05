const withCSS = require("@zeit/next-css");

module.exports = withCSS({
    publicRuntimeConfig: {
        API_URL: process.env.API_URL,
        S3_HOST: process.env.S3_HOST,
        S3_BUCKET: process.env.S3_BUCKET,
    }
});
