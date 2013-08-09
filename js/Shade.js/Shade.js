// Shade.js
// by Daniel Rodríguez
// http://sadasant.com/license

(function(W, D, J, U) { // window, document, JSON, undefined
  "use strict"

  // Empty Values
  // Avoiding building unnecesary objects
  var str   = ""
    , space = " "
    , eight = 8
    , one   = 1

  // Type Strings
  // This way, when checking for a type, we don't need to build a
  // string object from scratch.
  var type_func = "Function"
    , type_str  = "String"
    , type_obj  = "Object"

  // JSON Functions
  // Because repeating stringify in the code makes it heavier
  var J2O = J.parse
    , O2J = J.stringify

  // encodeURIComponent
  // Because repeating encodeURIComponent in the code makes it
  // heavier.
  var enComp = encodeURIComponent

  // eq
  // Built in testing framework (at least assertions).
  // It converts objects to json to compare them as strings.
  // It converts functions to strings to compare them.
  function eq(a, b, message, callback) {
    var ta = typeOf(a)
      , tb = typeOf(b)
    if (ta === type_obj  && tb === type_obj) {
      return callback(O2J(a) === O2J(b), message)
    }
    if (ta === type_func && tb === type_func) {
      return callback(str + a === str + b, message)
    }
    callback(a === b, message)
  }

  // neq
  // Because sometiems we want to check two things are really
  // different stuff.
  function neq(a, b, message, callback) {
    return eq(a, b, message, function(eq, message) {
      return callback(!eq, message)
    })
  }

  // Type checker
  // If undefined, returns "undefined"
  // If null, return "null"
  // Else, return the segment of it's type by parsing it to string
  function typeOf(object) {
    return object === U || object === null
      ? str + object
      : Object.prototype.toString.call(object).slice(eight, -one)
  }


  // Storage
  // localStorage handler
  // Shade.db.set(key, value)  // Set a Value with a Key
  // Shade.db.get(key, value)  // Get a Value from a Key
  // Shade.db.drop(key, value) // Deletes a Value from a Key
  function DB() {
    var storage = W.localStorage || {}
      , that = this
      , value
    that.get = function(key) {
      value = storage[key]
      return value
        ? typeOf(value) === type_func
          ? value
          : J2O(value)
        : value
    }
    that.set = function(key, value) {
      if (!value) {
        return false
      }
      storage[key] = typeOf(value) === type_str
        ? value
        : O2J(value)
    }
    that.drop = function(key) {
      for (key in storage) {
        delete storage[key]
      }
    }
  }


  // DOM doggie
  // Checks if it can be gotten with default functions
  // It doesn't cover all css stuff because some browsers doesn't
  // like document.querySelector
  function find(str, node) {
    node = node || D;
    var type  = typeOf(node)
      , strip = str.split(space)
      , head  = strip.shift()
      , tail  = strip.join(space)
      , index = 0
      , _node = []
      , found
    if (!~type.indexOf("HTML") || type === "HTMLCollection" && !node.length) {
      return
    }
    switch (head[0]) {
      case "." : node = node.getElementsByClassName(head.slice(one)) ; break;
      case "#" : node = node.getElementById(head.slice(one))         ; break;
      default  : node = node.getElementsByTagName(head)              ; break;
    }
    if (tail.length) {
      if (node.length) {
        _node = []
        for (; node[index]; index += one) {
          found = find(tail, node[index])
          if (found == U) {
            continue
          }
          if (found.length) {
            _node = _node.concat(found)
          } else {
            _node[_node.length] = found
          }
        }
        return _node.length > 1 ? _node : _node[0]
      } else {
        return find(tail, node)
      }
    } else {
      if (node.length === U) {
        return node
      } else {
        return node.length > 1 ? _node.slice.call(node) : node[0]
      }
    }
  }


  // Cross browser class helpers
  // Shade.hasClass()
  // Shade.addClass()
  // Shade.removeClass()
  function hasClass(node, className) {
    if (node.classList) {
      return node.classList.contains(className)
    } else {
      return !!~node.className.indexOf(className)
    }
  }
  function addClass(node, className) {
    if (node.classList) {
      node.classList.add(className)
    } else {
      node.className += space + className + space
    }
  }
  function removeClass(node, className) {
    if (node.classList) {
      node.classList.remove(className)
    } else {
      node.className = node.className.replace(new RegExp(className, "g"), str)
    }
  }


  // Events helper
  function addEvent(node, type, callback) {
    if (node.addEventListener) {
      node.addEventListener(type, callback, false)
    } else {
      node.attachEvent("on" + type, callback)
    }
  }


  // Object To URL
  // Used in Shade.http
  function urlify(object) {
    var url = []
      , key
    for (key in object) {
      url[url.length] = enComp(key) + "=" + enComp(object[key])
    }
    return url.join("&")
  }


  // Shade.http ajax
  // Parameters headers, data and calback are optionals.
  // Headers and data are objects
  // data can be a string like: "a=1&b=2"
  // You can use any HTTP Method
  function http(method, url, headers, data, callback) {
    var xhr = new XMLHttpRequest()
      , key
    if (data === U) {
      callback = headers
      headers = data = U
    } else if (callback === U) {
      callback = data
      data = headers
      headers = U
    }
    headers || (headers = { "Content-Type": "application/x-www-form-urlencoded" })
    xhr.onreadystatechange = function() {
      if (xhr.readyState === 4 && typeOf(callback) === type_func) {
        callback(xhr.status < 300, xhr.responseText, xhr)
      }
    }
    xhr.open(method, url, true)
    for (key in headers) {
      xhr.setRequestHeader(key, headers[key])
    }
    xhr.send(typeOf(data) === type_obj ? urlify(data) : data)
    return xhr
  }


  // Setting the shade
  W.Shade = {
    eq          : eq
  , neq         : neq
  , typeOf      : typeOf
  , db          : new DB()
  , find        : find
  , hasClass    : hasClass
  , addClass    : addClass
  , removeClass : removeClass
  , addEvent    : addEvent
  , http        : http
  }

})(window, document, JSON)
