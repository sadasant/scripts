/**
 * dom.js 0.0.46.5
 * Copyright (C) 2011 Daniel Rodríguez (sadasant.com)
 * Licence: GNU General Public Licence
 */

(function(){

  // Environment settings
  //·--------------------

  // Enviroment reference
  var _env  = this
    , _doc  = document

  // TODO: Store old `dom` and further integration testing.

  // Main namespace for the public API
  // TODO: CommonJS friendly?
  var dom  = _env.dom = { VERSION: '0.0.46.5' }


  // Tools
  //·-----

  // type checker
  const is = dom.is = (function() {
    var common = {
        s: 'string'
      , n: 'number'
      , o: 'object'
      , f: 'function'
      }
    , type = {}
    , checkType = function(str_type, e) { return typeof e == str_type }
    type.h = function(e) { return e && (e.toString().indexOf('HTML') == 8 || e.tagName) ? true : false }
    type.a = function(e) { return Array.isArray(e) }
    for (var k in common) {
      type[k] = checkType.bind(null, common[k])
    }
    return type
  })()

  // String bindings
  const string = dom.string = {
    trim  : function(s) { return (s.trim) ? s.trim() : s.replace(/^\s+|\s+$/g,"") }
  }

  // Object bindings
  const object = dom.object = {
    single : function(name, val) { var o = {}; o[name] = val; return o}
  }

  // Array bindings
  const array = dom.array = (function() {
    const proto = Array.prototype
    return {
      slice  : function(a) { return proto.slice.call(a) }
    , filter : function(a, f) { return proto.filter.call(a, f) }
    , clean  : function(a) { return a.filter(function (e) { return e }) }
    , unique : function(a) {
        var r = []
          , len = a.length
        if (len < 2) return a
        o:for (var i = 0; i < len; i++) {
          for (var j = 0; j < r.length; j++) {
            if (r[j] === a[i]) continue o
          }
          r[r.length] = a[i]
        }
        return r
      }
    }
  })()


  // TODO: Settings
  //·--------------


  // Methods
  //·-------

  // Get and set attrs
  dom.att = (function() {
    const attributes = {
        '#': 'id'
      , '.': 'class'
      }
    , notAtt = { tag: 1, html: 1 }

    // dom.att method
    return function () {
      var args = this._args || array.slice(arguments)
        , att = this._att
        , el = [].concat(this._el) || []
        , str_el = this._str_el
        , str_att = this._str_att || []
        , index = this._index

      // digest arguments
      var a, i = 0
      while (a = args[i++]) {
        if (is.s(a)) {
          str_att[str_att.length] = a
        } else
        if (is.a(a)) { args = args.concat(a) } else
        if (is.h(a)) { el[el.length] = a } else
        if (is.o(a)) { att = a }
      }

      if (str_att.length == 2) {
        if (!att) att = {}
        att[str_att[0]] = str_att[1]
      }

      if (el) {
        el = array.clean(el)
        if (el.length < 2) el = el[0]
        else
        if (el.length > 1) return el.map(function(e){return dom.att.apply({_el:e, _args:[att], _str_att:str_att})})
      }

      // set attributes
      if (att) {
        // getting attributes
        for (var k in att) {
          var attk = att[k]
          if (!attk || !~attk.indexOf(':')) continue
          attk = attk.split(':')
          if (attk[1] == 'i') { att[k] = attk[0] + index } else
          if (attk[1] == 'h') { att[k] = attk[0] + att.html } else
          if (att[attk[1]])   { att[k] = attk[0] + att[attk[1]] }
        }
        // setting attributes
        if (el) {
          for (var k in att) {
            var attk = att[k]
            if (!attk) continue
            if (notAtt[k]) { if (k == 'html') el.innerHTML = attk }
            else { el.setAttribute(k, attk) }
          }
        }
        return el
      } else

      // get attributes
      if (el){
        var att = {}
          , el_att = el.attributes
        for (var k in el_att) {
          att[el_att[k].name] = el_att[k].textContent
        }
        att.html = el.innerHTML
        if (str_att.length) att = att[str_att[0]]
        return att
      } else

      // read attributes on a given string using CSS selectors and other dom.js proper syntax
      if (str_el) {
        var isAtt
        el = str_el
        for (var i = 0; i < el.length; i++) {
          if(isAtt = attributes[el[i]]) {
            el = el.split(el[i])
            if (!att) att = {}
            att[isAtt] = el[1]
            if (i && el[0]) att.tag = el[0]
            break
          }
        }
        if (!att) return { tag: el }
        for (var k in att) {
          var _att = dom.att.apply({ _str_el:att[k] })
          for (var _k in _att) {
            att[(_k == 'tag') ? k : _k] = _att[_k]
          }
        }
        return att
      }
    }
  })()


  // Returns one or more elements found
  dom.get = (function() {
    var special_fun = {
        active        : function() { return _doc.activeElement }
      , selection     : function() { return _env.getSelection ? _env.getSelection() : _doc.getSelection ? _doc.getSelection : _doc.selection.createRange().text }
      // TODO: Move this to dom.pos()
      , position      : function() {
          var x = 0
            , y = 0
            , el = this._el
          while (el && !isNaN(el.offsetLeft) && !isNaN(el.offsetTop)) {
            x += el.offsetLeft - el.scrollLeft
            y += el.offsetTop  - el.scrollTop
            el = el.offsetParent
          }
          return {x:x, y:y}
        }
      }
    , specials = []

    // Array with special keys
    for (var k in special_fun) {
      specials[specials.length] = k
    }

    // dom.get method
    return function() {
      var args = this._args || array.slice(arguments)
        , doc = this._doc || _doc
        , els = []
        , limit
        , cb
        , a
        , i = 0

      // digest arguments
      while (a = args[i++]) {
        if (is.n(a)) { limit = a } else
        if (is.s(a)) {
          // special strings behavior
          if (~specials.indexOf(a)) { return special_fun[a].apply({_el:doc}) }
          // normal strings behavior
          if (!is.a(doc)) doc = [doc]
          for (var j = 0; j < doc.length; j++) {
            var el = array.slice(doc[j].querySelectorAll(a))
            els = els.concat(el)
          }
        } else
        if (is.a(a)) { args = args.concat(a) } else
        if (is.h(a)) { els[els.length] = a   } else
        if (is.f(a)) { cb = a          } else
        if (is.o(a)) {
          // for ubiquity requests and not-query related.
          var el, u
          if (a.id)    { el = _doc.getElementById(a.id) }
          if (a.class) {
            if (!el) { el = array.slice(_doc.getElementsByClassName(a.class)) } else
            if (!~el.className.indexOf(a.class)) el = u
          }
          if (a.tag && (a.tag = a.tag.toUpperCase())) {
            if (!el) { el = array.slice(_doc.getElementsByTagName(a.tag)) }
            else {
              if (!is.a(el)) el = [el]
              for (var i = 0; i < el.length; i++) {
                if (el[i].tagName !== a.tag) el[i] = u
              }
              el = array.clean(el)
            }
          }
          // TODO: Move this to dom.pos
          if (a.x && a.y) {
            var _el = _doc.elementFromPoint ? _doc.elementFromPoint(a.x, a.y) : _doc.getElementFromPoint(a.x, a.y)
            if (el && _el !== el) el = u
            else el = _el
          }
          return cb ? cb(el) : el
        }
      }

      // ending
      els = els.length < 2 ? els[0] : array.unique(els)
      if (limit) els = els.slice(0,limit)
      return cb ? cb(els) : els
    }
  })()

  // Appends or write into an object or parent
  dom.add = function() {
    var args = this._args || array.slice(arguments)
      , dad = [].concat(this._dad || _doc.body || _doc.firstChild)
      , els = []
      , cb
      , a
      , i = 0

    // parsing arguments
    while (a = args[i++]) {
      if (is.s(a)) {
        a.split('\n').forEach(function(e, i, a) {
          els[els.length] = _doc.createTextNode(e)
          if (i < a.length - 1) {
            els[els.length] = dom.new('br')
          }
        })
      } else
      if (is.a(a)) { args = args.concat(a) } else
      if (is.h(a)) { els[els.length] = a   }
      if (is.f(a)) { cb = a                }
    }

    // ending
    for (var i = 0; i < dad.length; i++) {
      for (var j = 0; j < els.length; j++) {
        // clone if parents are many
        dad[i].appendChild(!i ? els[j] : els[j].cloneNode(true))
      }
    }
    return cb ? cb(true) : true
  }

  // Remove and clone one or more elements from its place
  dom.dig = function() {
    var args       = this._args || array.slice(arguments)
      , doc        = this._doc  || _doc
      , dont_clone = this._dont_clone
      , els        = []
      , cb

    function dig(o) {
      if (!dont_clone) {
        _o = o.cloneNode(true)
        els[els.length] = _o
      }
      o.parentNode.removeChild(o)
    }

    var a
      , i = 0
    while (a = args[i++]) {
      if (is.s(a)) { [].concat(dom.get(a)).map(dig) } else
      if (is.a(a)) { args = args.concat(a) } else
      if (is.h(a)) { dig(a) } else
      if (is.f(a)) { cb = a }
    }

    if (els.length < 2) els = els[0]

    return cb ? cb(els) : els
  }

  // Remove one or more HTML elements
  dom.rm = dom.dig.bind({_dont_clone:true})

  // Returns an HTML element found by id, class or tag name.
  dom.new = function() {
    var args = this._args || array.slice(arguments)
      , doc  = this._doc  || _doc
      , els = []
      , el
      , att = {}
      , i = 0
      , index
      , cb

    // main element
    el = args.splice(0, 1)[0]

    if (is.s(el)) el = el.split(',').map(function(e){ return string.trim(e) })
    if (is.h(el) && !args.length) return el
    el = el[0]

    // get attributes
    att = dom.att.apply({_str_el:el})
    if (att.tag) { el = att.tag }

    // parsing arguments
    var a, i = 0
    while (a = args[i++]) {
      if (is.a(a)) {
        for (var j = 0; j < a.length; j++) {
          if (!is.h(a[j])) { els[els.length] = a[j] }
          else {
            els[els.length] = dom.in(a[j]).att()
          }
        }
      } else
      if (is.s(a)) { !att.html && (att.html = a)        } else
      if (is.n(a)) { index = a                          } else
      if (is.h(a)) { els[els.length] = a                } else
      if (is.o(a)) { for (var k in a) { att[k] = a[k] } } else
      if (is.f(a)) { cb = a }
    }

    // ending
    if (els && els.length) {
      return els.map(function (e, i) {
        return dom.new(el, att, e, i)
      })
    }
    el = doc.createElement(el)
    el = dom.att.apply({_el:el, _index:index}, [att])
    return cb ? cb(el) : el
  }

  // Returns one object with dom.js methods.
  dom.in = function(el) {
    !is.h(el) && (el = el ? dom.get(el) : _doc)
    return {
      att : dom.att.bind({ _el : el }) // the _args in the prototype is to avoid repeating the slice() method
    , get : dom.get.bind({ _doc: el })
    , add : dom.add.bind({ _dad: el })
    , new : function() {
        var _el = dom.new.call({ _args: array.slice(arguments) })
        dom.in(el).add(_el)
        return _el
      }
      // dom.in(dom.new('b')).out() // <b>
    , out : function() { return el }
    }
  }

})();
