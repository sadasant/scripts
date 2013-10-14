define ['Matrix'], (Matrix)->
    (n)->
        m = new Matrix n
        for i in [0..n-1]
            m[i][i] = 1
        m
