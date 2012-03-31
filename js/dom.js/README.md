### dom.js v0.0.46.5

dom.js is a simple DOM manipulation library

### API Goal

Some of this goals may change with time.

*v0.0.5 goals*

- `dom.get()`: Find / pick / select / query one or more elements by CSS selectors and other parammeters.
- `dom.att()`: Gets / sets the attributes for one or more elements.
- `dom.pos()`: Get the position of an element or set the position to a given one.
- `dom.add()`: Insert / put / fill / write one or more arguments given.
- `dom.dig()`: Remove and return a clone of one or more elements from its place.
- `dom.rm()` : Remove one or more elements from its place.
- `dom.new()`: Create one or more elements.
- `dom.in()` : Allows you to excecute dom.js API inside other elements, with some additional behaviour:
- `dom.out()`: Returns a copy of the element inside in() with the applied changes.

*v0.0.7 goals*

- `dom.conf`  : Will have all the configurations needed to log and capture the data while running.
- `dom.css()` : Gets / sets the CSS values for one or more elements.
- `dom.on()`  : Gets / sets the event handlers from one or more elements.
- `dom.do()`  : Emulates user actions over the site.
- `dom.wait()`: Create and manage setTimeouts and setIntervals.

*v0.1.0 goals*

- `dom.store()`: Stores data, elements, etc in cookies or localstorage.
- `dom.ajax()` : AJAX.

### Limitations

- **Custom objects:**
    dom.js will be limited to work with real DOM objects,
without a higher level of abstraction, in order to
provide both a simple and quick API for the dom,
and a clean way to understand the DOM workflow.

### Example of usage

The capabilities of this library are under development,
currently it's able to work as described below, but
it will go beyond in the upcoming days.

Code:

```javascript
(function () {

  var options = [
    'filter'
  , 'forEach'
  , 'map'
  , 'some'
  , 'reduce'
  , 'reduceRight'
  ]

  var select = dom.new('select#sel')

  options = dom.in(select).new('option.option#op-:i', options)
  // or:
  //   options = dom.new('option.option#op-:index', options)
  //   dom.in(select).add(options)

  console.log(options)
  // [option#op-0.option, option#op-1.option, option#op-2.option, option#op-3.option, option#op-4.option, option#op-5.option]

  dom.add('Method:\n\n')

  dom.add(select)

})()
```

Output:

```html
Method:
<br/>
<br/>
<select id="sel">
  <option id="op-0" class="option">filter</option>
  <option id="op-1" class="option">forEach</option>
  <option id="op-2" class="option">map</option>
  <option id="op-3" class="option">some</option>
  <option id="op-4" class="option">reduce</option>
  <option id="op-5" class="option">reduceRight</option>
</select>
```

---

You can take a look to the *state of art* inside the *tests* directory.

---

### License

dom.js : a simple DOM manipulation library.

Copyright (c) 2011 Daniel Rodríguez (sadasant.com)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

