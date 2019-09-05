const withCSS = require("@zeit/next-css");

module.exports = withCSS({
    publicRuntimeConfig: {
        API_URL: process.env.API_URL,
    }
});
