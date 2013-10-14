define ()->
    toString=()->
        s = []
        this.forEach (e)->
            s.push e.join && e.join(", ") || e
        s.join("\n")
    (x, y)->
        y = y || x
        x -= 1
        m = Object.create Array.prototype
        m = Array.apply m, [x]
        for i in [0..x]
            m[i] = Array.apply(null, new Array(y)).map Number.prototype.valueOf, 0
        m.toString = toString
        m
