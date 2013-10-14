define [
    'gradients'
    'eye'
], (gradients, eye)->
    gradients.hello('world')
    console.log(gradients.normal())
    console.log(eye(5).toString())
