# S.js

Simple CSJS scripts for building and testing websites.

S.js is focused on simplicity, not backwards browser compatibility, nor stability.

### MAIN API

---
### `S.q`

Query DOM Elements using `document.querySelectorAll`.  
If it finds just one, it will be returned alone (not inside a NodeList).  

_Example:_

```javascript
S.q('*')
S.q('[name=user]')
S.q('[name^=pass]')
```

---
### `S.t`

Returns the type of a given argument using it's object prototype.  

_Example:_  

```javascript
S.t('')           // 'String'
S.t([])           // 'Array'
S.t(/\//g)        // 'RegEx'
S.t(new Date())   // 'Date'
S.t(localStorage) // 'Storage'
S.t(S.q('*'))     // 'NodeList'

S.t(null)         // null
S.t()             // undefined
```

---
### `S.u`

Makes an URL Query String using a given Object.  
It won't work well if you put objects inside the object.  

_Parammeters:_

1. An object `{ hello : 'hello', goodbye : 'goodbye' }`

_Returns:_  

The stirng "hello=hello&goodbye=goodbye"

---
### `S.wipe`

Deletes all the keys in the `S.s` storage object.

---
### `S.req`

A monolithic function for making HTTP Requests.  

_Parammeters:_

1. HTTP Method String: 'GET', 'POST', etc.

2. URI String

3. Set of headers: `{ "Content-Type": "application/x-www-form-urlencoded" }`

4. Data Object or String (If you send and object, it will be converted to an URL query String using `S.u`)

5. Callback function, which will receive:

    1. True if the ready state is 4

    2. The response text

    4. The current XMLHttpRequest object

6. You can send here a different XMLHttpRequest object

_Returns:_  

Returns the XMLHttpRequest Object that is used to make the request.  

_Example:_  

```javascript
S.req('POST', '/ajax', null, { 'l\\l' : 'l/l', lal : 'l&l' }, function(ok, txt) {
  console.log(JSON.parse(txt))
})
```

---
### Who uses this?

I've been using it on [my website](http://sadasant.com/) and my [node.js chat](http://talk.nodester.com/).

---
### Thank you

Thanks for stopping by :)

---
### License

Copyright (c) 2012,  
Daniel Jesús Rodríguez Sivira  
http://sadasant.com/  
All Rights Reserved.  

GNU GPL 3 Licensed.  
<http://sadasant.com/license>
