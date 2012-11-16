Server.local.waitForBoot({

  { var ampOsc, frecOsc, o;
    ampOsc  = SinOsc.kr(0.5, 1.5pi, 0.5, 0.5);
    frecOsc = SinOsc.kr(0.1, 1.5pi, 0.5, 0.5);
    o = frecOsc * 440;
    SinOsc.ar([o + 220, o + 210], 0, ampOsc);
  }.play;


});
