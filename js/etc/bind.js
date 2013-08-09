var log = "";

Function.prototype.bind = function (oThis) {
    if (typeof this !== "function") {
      // closest thing possible to the ECMAScript 5 internal IsCallable function
      throw new TypeError("Function.prototype.bind - what is trying to be bound is not callable");
    }

    var aArgs = Array.prototype.slice.call(arguments, 1);
    var fToBind = this;
    var fNOP = function () {};
    var fBound = function () {
        log += typeof this + " " + Object.prototype.toString.call(this).slice(8, -1) + " " + (this instanceof fNOP) + " oThis: " + typeof oThis + " " + Object.prototype.toString.call(oThis).slice(8, -1) + " " + (oThis instanceof fNOP) + "\n";
        log += (this instanceof fNOP && oThis) + " " + (!!oThis) + " \n";
        // throw typeof this + " " + Object.prototype.toString.call(this).slice(8, -1) + typeof oThis + " " + Object.prototype.toString.call(oThis).slice(8, -1);
        return fToBind.apply(
            // this instanceof fNOP && oThis ? this : oThis,
            oThis,
            // typeof this === "object" && oThis ? this : oThis,
            // Object.prototype.toString.call(this).slice(8, -1) !== "global" && oThis ? this : oThis,
            aArgs.concat(Array.prototype.slice.call(arguments)));
    };

    fNOP.prototype = this.prototype;
    fBound.prototype = new fNOP();

    return fBound;
  };

this.x = 9;
var module = {
  x: 81,
  getX: function() { return this.x; }
};

console.log(module.getX()); // 81

this.getX = module.getX;
console.log(this.getX()); // 9, because in this case, "this" refers to the global object

// create a new function with 'this' bound to module
var boundGetX = this.getX.bind(module);
console.log(boundGetX()); // 81

console.log(log);
