Server.local.waitForBoot({

  { [SinOsc.ar(440, 0, 0.2), SinOsc.ar(442, 0, 0.2)] }.play;

  f = { "Function evaluated".postln; };
  f.value;
  f.value;
  f.value;

});
