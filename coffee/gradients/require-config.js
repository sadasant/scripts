require.config({
    urlArgs: "bust="+(new Date()).getTime(),
    baseUrl: "js-build",
    paths: {
        text: "../lib/text"
    },
    shim: {
        gradients: {
            deps: ["../lib/jstat"]
        }
    }
});
require(["main"]);
