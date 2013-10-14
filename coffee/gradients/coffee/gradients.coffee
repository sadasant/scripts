define ()->
    normal = ()->
        norm = new NormalDistribution(0,1)
        norm.getQuantile(0.95)
    {
        hello: (s)-> console.log('hello '+s)
        normal: normal
    }
