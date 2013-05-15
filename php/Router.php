<?PHP
// Router.php
// by Daniel Rodríguez
// sadasant.com/license
//
// Example index.php
/* -----------------

include('lib/Router.php');

$Router->get('/', 'index');
$Router->get('/data/customers/all.json', 'apiCustomers');
$Router->otherwise('F404'); // But tsverp has it's own 404
$Router->run();

function index() {
	echo "Hello";
}

/* -- */

class Router {

	// Function Routes
	private $FUNCS = array(
		"GET"  => array(),
		"HEAD" => array(),
		"POST" => array(),
		"PUT"  => array(),
	);

	function get ($path, $func) { $this->FUNCS["GET" ][$path] = $func; }
	function head($path, $func) { $this->FUNCS["HEAD"][$path] = $func; }
	function post($path, $func) { $this->FUNCS["POST"][$path] = $func; }
	function put ($path, $func) { $this->FUNCS["PUT" ][$path] = $func; }
	function otherwise($func)   { $this->FUNCS["OTHERWISE"  ] = $func; }

	function run() {
		$request = ($path = explode('.php/', strtolower($_SERVER['PHP_SELF']))) ? $path[1] : '';
		$method  = $_SERVER['REQUEST_METHOD'];
		$func    = $this->FUNCS['OTHERWISE'];
		// The function exists in our functions file?
		foreach ($this->FUNCS[$method] as $path => $f) {
			if ($path[0] != '/') {
				$path .= '/'.$path;
			}
			$len = strlen($path);
			if ($len == 1 || $path[$len - 1] != '/') {
				$path .= '/';
			}
			if (preg_match($path, $request)) {
				$func = $f;
				break;
			}
		}
		return call_user_func($func, $path);
	}

}

// Creating the Object
$Router = new Router();
?>
