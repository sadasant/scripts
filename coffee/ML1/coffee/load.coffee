define ()-> (path)->
    data = null

    $.ajax
        url: path
        type: 'get'
        async: false
        success: (d)-> data = d

    (data.split '\n').map (e)->
        e = e.split ','
        x = parseFloat e[0]
        y = parseFloat e[1]
        [x, y]
