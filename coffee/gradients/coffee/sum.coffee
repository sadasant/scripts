define ()-> (a...)->
    f = a.pop()
    v = []
    i = 0
    sum = 0
    while true
        v = []
        for _a in a
            return sum if !_a[i]
            v.push _a[i]
        sum += f.apply null, v
        i++
    sum
