cd lib
rm *.js

function get() {
    echo $1
    wget -q -O $@
}

get jstat.js   http://www.jstat.org/jstat-1.0.0.min.js
get require.js http://requirejs.org/docs/release/2.1.8/minified/require.js
get text.js    https://raw.github.com/requirejs/text/latest/text.js
