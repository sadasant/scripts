// Ink.js
// By Daniel R. (sadasant.com)
// License: http://opensource.org/licenses/mit-license.php
// Homepage: https://github.com/sadasant/Ink.js

~function (W) {

  // Global Variables
  var D = W.document
    , M = W.Math
    , O = W.Object
    , U // Undefined

  // Private Variables
  var draw_stack = {}   // Stack of shapes
    , ids        = 0    // Current last id of shapes
    , removed    = {}   // Shapes' graveyard
    , can               // Inner Canvas holder
    , con               // Inner Context holder
    , dcan              // Document's Canvas holder
    , dcon              // Document's Context holder
    , reqAF             // Request Animation Frame
    , windowSize        // Window size object
    , started           // Started Flag
    , busy              // Busy Flag

  // Math variables
  var PI180 = M.PI/180
    , PI2   = M.PI*2
    , COS   = M.cos
    , SIN   = M.sin
    , ABS   = M.abs


  // Request Animation Frame
  // Based on code by Paul Irish
  reqAF = function() {
    return W.requestAnimationFrame
        || W.webkitRequestAnimationFrame
        || W.mozRequestAnimationFrame
        || W.oRequestAnimationFrame
        || W.msRequestAnimationFrame
        || function(f) { W.setTimeout(f, 1000 / 60) }
  }().bind(W)


  // Getting the window size
  windowSize = function() {
    return {
      w : W.innerWidth  || (D.documentElement && D.documentElement.offsetWidth ) || (D.body && D.body.offsetWidth ) || 630
    , h : W.innerHeight || (D.documentElement && D.documentElement.offsetHeight) || (D.body && D.body.offsetHeight) || 460
    }
  }()


  // Looping through the stack of shapes
  function drawStack() {
    if (busy) return reqAF(drawStack)
    reqAF(drawStack)
    var k
    for (k in draw_stack) {
      if (draw_stack[k]) {
        con.save()
        draw_stack[k].draw()
        con.restore()
      }
    }
    dcon.drawImage(can, 0, 0)
  }


  // Insert in the stack of shapes
  function draw(obj) {
    var a = arguments
      , i = 0
      , l = a.length
    for (; i < l; i++) {
      draw_stack[a[i].id] = a[i]
    }
  }


  // Move the shape to the graveyard
  function remove(obj) {
    if (obj) {
      removed[obj.id] = true
      delete draw_stack[obj.id]
    }
  }


  // Prototypes' holder
  var Main = {
    // Geometric shapes' prototype
    geometrics : {
      // Always inside the canvas
      foreverInScope: function() {
        if (this.x > can.width -10)  this.x -= can.width
        if (this.x < -10          )  this.x += can.width -10
        if (this.y > can.height   )  this.y -= can.height
        if (this.y < -10          )  this.y += can.height
      }
      // degrees 2 rad
    , rotateTo: function (deg) {
        this.rotation = deg * PI180
      }
      // move forward this much with this maxSpeed
    , accel: function(much, maxSpeed) {
        if (maxSpeed) {
          this.maxSpeed.x = maxSpeed
          this.maxSpeed.y = maxSpeed
        }
        if (this.rotateInEdge) {
          var a = this.rotation
            , x = - much * SIN(a)
            , y =   much * COS(a)
          this.speed.x += x
          this.speed.y += y
        } else {
          this.speed.x += much
          this.speed.y += much
        }
        var abs = {
              x: ABS(this.speed.x)
            , y: ABS(this.speed.y)
          }
        if (abs.x > this.maxSpeed.x) {
          this.speed.x = (abs.x/this.speed.x)*this.maxSpeed.x
        }
        if (abs.y > this.maxSpeed.y) {
          this.speed.y = (abs.y/this.speed.y)*this.maxSpeed.y
        }
      }
    , stop: function() {
        this.speed.x = 0
        this.speed.y = 0
      }
    , repos: function() {
        if (this.rotateInEdge) {
          this.x -= this.speed.x
          this.y -= this.speed.y
        } else {
          var a = this.rotation
          this.x += this.speed.x * SIN(a)
          this.y -= this.speed.y * COS(a)
        }
        this._x = this.x - this.half_width
        this._y = this.y - this.half_height
      }
      // lame collider
    , addCollider: function(obj) {
        this.colliders[this.colliders.length] = obj
      }
      // onCollide by default just turns red the shape
    , onCollide: function() {
        this.fill   = "rgba(255, 0, 0, 0.3)"
        this.stroke = "rgba(255, 0, 0, 1)"
      }
    , collideArea: 15
    , collide: function() {
        var col = this.colliders
          , coli
          , i = 0
          , l = col.length
          , diffx
          , diffy
        for (; i !== l; i += 1) {
          if (coli = col[i]) {
            diffx = ABS(coli.x - this.x)
            diffy = ABS(coli.y - this.y)
            if (diffx < this.collideArea && diffy < this.collideArea) {
              if (removed[coli.id]) {
                delete this.colliders[i]
              } else {
                this.onCollide(coli)
              }
            }
          }
        }
      }
    }
  // Gradients' prototype
  , grad : {
      // arguments = [ float1, color1 ], [ float2, color2 ], ...
      colors : function(colors) {
        var i = 0
          , l = colors.length
          , c
        for (; i !== l; i += 1) {
          c = colors[i]
          this.addColorStop(c[0], c[1])
        }
        return this
      }
    }
  }


  // Rectangle
  // a = {
  //   x            : Number // Position x
  // , y            : Number // Position y
  // , width        : Number
  // , height       : Number
  // , fill         : String || Ink.grad
  // , border_color : String || Ink.grad
  // , border_width : Number
  // }
  function Rect(a) {
    if (!a) a       = {}
    var s           = O.create(Main.geometrics)
      , path_op
      , width_border
      , height_border
    s.id            = '' + (ids += 1)
    s.x             = a.x
    s.y             = a.y
    s.width         = a.width
    s.height        = a.height
    s.half_width    = a.width / 2
    s.half_height   = a.height / 2
    s._x            = a._x !== U ? a._x : - s.half_width
    s._y            = a._y !== U ? a._y : - s.half_height
    s.fill          = a.fill || "rgba(0, 0, 0, 0.2)"
    s.border_color  = a.border_color  || "rgba(0, 0, 0, 0.2)"
    s.border_width  = a.border_width  || 0
    s.border_radius = a.border_radius || 0
    if (s.border_radius) {
      // In order to provide a border radius, it must be a path
      // so we auto-generate the path with the code below.
      // This could be the beggining of more auto-generated forms.
      path_op = [
          - s.half_width
        , - s.half_width + s.border_radius
        ,   s.half_width - s.border_radius
        , s.half_width
        , - s.half_height
        , - s.half_height + s.border_radius
        ,   s.half_height - s.border_radius
        , s.half_height
        ]
      s.path = [
        [ path_op[1], path_op[4]] // Starting point
      , [ path_op[2], path_op[4], path_op[3], path_op[4], path_op[3], path_op[5]]
      , [ path_op[3], path_op[6], path_op[3], path_op[7], path_op[2], path_op[7]]
      , [ path_op[1], path_op[7], path_op[0], path_op[7], path_op[0], path_op[6]]
      , [ path_op[0], path_op[5], path_op[0], path_op[4], path_op[1], path_op[4]]
      ]
      delete path_op
    }
    s.rotation      = 0
    s.speed         = { x:  0, y:  0 }
    s.maxSpeed      = { x: 10, y: 10 }
    s.infiniteScope = null
    s.colliders     = []
    // Draw method
    // We must use Path_draw function
    // in order to draw a rectangle with border radius.
    s.draw          = s.border_radius ? Path_draw : Rect_draw
    s.afterDraw     = U
    return s // this
  }
  function Rect_draw() {
    // Every element has the following 18 lines
    // I've been thinking on making a separated function to
    // reduce the code, but that would make it slower, no?
    if (this.infiniteScope) {
      // Check if it's out of the scope
      this.foreverInScope()
    }
    if (this.speed.x || this.speed.y) {
      // Move it according to it's speed
      this.repos()
    }
    // Translating the context
    con.translate(this.x, this.y)
    if (this.rotation) {
      // Rotate
      con.rotate(this.rotation)
    }
    if (this.colliders.length) {
      // Collide
      this.collide()
    }
    // draw
    con.fillStyle   = this.fill
    con.strokeStyle = this.border_color
    con.lineWidth   = this.border_width
    con.fillRect(this._x, this._y, this.width, this.height)
    this.afterDraw && this.afterDraw()
  }


  // Circle
  // a = {
  //   x            : Number // Position x
  // , y            : Number // Position y
  // , r            : Number // Radius
  // , fill         : String || Ink.grad
  // , border_color : String || Ink.grad
  // , border_width : Number
  // }
  function Circ(a) {
    if (!a) a       = {}
    var s           = Object.create(Main.geometrics)
    s.id            = (ids += 1).toString()
    s.x             = a.x || 0
    s.y             = a.y || 0
    s.r             = a.r || 0
    s.fill          = a.fill || "rgba(255, 255, 255, 0.3)"
    s.border_color  = a.border_color || "rgba(0, 0, 0, 0.2)"
    s.border_width  = a.border_width || 0
    s.rotation      = 0
    s.speed         = { x:  0, y:  0 }
    s.maxSpeed      = { x: 10, y: 10 }
    s.infiniteScope = null
    s.colliders     = []
    s.draw          = Circ_draw // draw method
    s.afterDraw     = U
    return s // this
  }
  function Circ_draw() {
    if (this.infiniteScope) {
      // Check if it's out of the scope
      this.foreverInScope()
    }
    if (this.speed.x || this.speed.y) {
      // Move it according to it's speed
      this.repos()
    }
    // Translating the context
    con.translate(this.x, this.y)
    if (this.rotation) {
      // Rotate
      con.rotate(this.rotation)
    }
    if (this.colliders.length) {
      // Collide
      this.collide()
    }
    // draw
    con.fillStyle   = this.fill
    con.strokeStyle = this.border_color
    con.lineWidth   = this.border_width
    con.beginPath()
    con.arc(0, 0, this.r, 0, PI2)
    con.closePath()
    con.stroke()
    con.fill()
    this.afterDraw && this.afterDraw()
  }


  // Path
  // a = {
  //   x            : Number // Position x
  // , y            : Number // Position y
  // , path         : Array
  // , fill         : String || Ink.grad
  // , border_color : String || Ink.grad
  // , border_width : Number
  // }
  function Path(a) {
    if (!a) a       = {}
    var s           = Object.create(Main.geometrics)
    s.id            = (ids += 1).toString()
    s.x             = a.x
    s.y             = a.y
    s.path          = a.path || []
    s.fill          = a.fill || "rgba(255, 255, 255, 0.3)"
    s.border_color  = a.border_color || "rgba(0, 0, 0, 0.2)"
    s.border_width  = a.border_width || 0
    s.rotation      = 0
    s.speed         = { x:  0, y:  0 }
    s.maxSpeed      = { x: 10, y: 10 }
    s.infiniteScope = null
    s.colliders     = []
    s.draw          = Path_draw // draw method
    s.afterDraw     = U
    return s // this
  }
  function Path_draw() {
    if (this.infiniteScope) {
      // Check if it's out of the scope
      this.foreverInScope()
    }
    if (this.speed.x || this.speed.y) {
      // Move it according to it's speed
      this.repos()
    }
    // Translating the context
    con.translate(this.x, this.y)
    if (this.rotation) {
      // Rotate
      con.rotate(this.rotation)
    }
    if (this.colliders.length) {
      // Collide
      this.collide()
    }
    // draw
    con.fillStyle   = this.fill
    con.strokeStyle = this.border_color
    con.lineWidth   = this.border_width
    con.beginPath()
    var i = 0
      , l = this.path.length
      , vi
    for (; i !== l; i += 1){
      vi = this.path[i]
      if (i) con.lineTo(vi[0], vi[1])
      else   con.moveTo(vi[0], vi[1])
      if (vi.length === 6) {
        con.quadraticCurveTo(vi[2], vi[3], vi[4], vi[5])
      }
    }
    con.closePath()
    con.stroke()
    con.fill()
    this.afterDraw && this.afterDraw()
  }

  // grad
  // a = {
  //   type = String
  // , x1   = Number
  // , x2   = Number
  // , y1   = Number
  // , y2   = Number
  // , r1   = Number
  // , r2   = Number
  // }
  function grad(a) {
    if (!a) a = {}
    if (typeof a.type !== 'string') return
    a.type = a.type.toUpperCase()
    var grad
    switch (a.type) {
      case 'LINEAR':
        grad = con.createLinearGradient(a.x1, a.y1, a.x2, a.y2)
        break
      case 'RADIAL':
        grad = con.createRadialGradient(a.x1, a.y1, a.r1, a.x2, a.y2, a.r2)
        break
    }
    grad.colors = Main.grad.colors
    return grad
  }

  function doneInit() {
    busy = false
    if (!started) {
      started = true
      drawStack()
    }
  }

  // Initialization
  // a = {
  //   id     : String // HTML Element Id
  // , width  : Number
  // , height : Number
  // , clear  : String // To clear the canvas with a certain color
  // }
  function init(a) {
    if (!a) a   = {}
    busy        = true
    draw_stack  = {}
    ids         = 0
    removed     = {}
    this.dcan   = dcan = this.dcan || D.getElementById(a.id || 'canvas')
    this.dcon   = dcon = this.dcon || dcan.getContext('2d')
    this.can    = can  = this.can  || D.createElement('canvas')
    this.con    = con  = this.con  || can.getContext('2d')
    can.height  = dcan.height = a.height || windowSize.h
    can.width   = dcan.width  = a.width  || windowSize.w
    // If there's no width or height, the user wants the canvas to take all the window,
    // for which the body must hide the overflow.
    if (!(a.width && a.height)) {
      D.body.style.overflow = 'hidden'
      dcan.style.margin = '-8px -8px'
    }
    this.center = { x: windowSize.w/2, y: windowSize.h/2 }
    if (a.clear) {
      Ink.draw(new Ink.Rect({
        _x : 0
      , _y : 0
      , width  : can.width
      , height : can.height
      , fill   : a.clear
      }))
    }
    return setTimeout(doneInit, 100)
  }

  // API
  var Ink = W.Ink = {
    VERSION : '0.3.5'
  , can     : null
  , con     : null
  , init    : init
  , draw    : draw
  , remove  : remove
  , Path    : Path
  , Rect    : Rect
  , Circ    : Circ
  , grad    : grad
  , reqAF   : reqAF
  }

}(this)
