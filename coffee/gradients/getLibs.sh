function get() {
    echo $1
    wget -q -O $@
}

cd lib; rm *.*
get jquery.js    http://code.jquery.com/jquery-1.10.2.min.js
get require.js   http://requirejs.org/docs/release/2.1.8/minified/require.js
get sylvester.js http://cdnjs.cloudflare.com/ajax/libs/sylvester/0.1.3/sylvester.js
