require.config({
    urlArgs: "bust="+(new Date()).getTime(),
    baseUrl: "js-build",
    shim: {
        main: {
            deps: [
                '../lib/jquery',
                '../lib/sylvester'
            ]
        }
    }
});
require(["main"]);
