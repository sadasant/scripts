define [
    'eye'
    'text!../data/ex1data1.txt'
], (eye, data)->
    console.log(eye(5).toString())

    data = (data.split '\n').map (e)->
        e = e.split ','
        x = parseFloat e[0]
        y = parseFloat e[1]
        [x, y]

