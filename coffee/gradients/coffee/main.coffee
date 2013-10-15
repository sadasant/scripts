define [
    'load'
    'sum'
], (load, sum)->

    # Identity Matrix
    console.log Matrix.I(5).inspect()

    data = load '../data/ex1data1.txt'
    X = data.map (e)-> [1, e[0]]
    Y = data.map (e)-> e[1]
    m = Y.length
    h = Matrix.Zero 2, 1
    J = 1/(2*m)
    J *= sum X, Y, (x, y)->
        Math.pow h.x($M x) - y, 2

    # Cost is: 31.745461082787738
    console.log 'Cost is:', J
