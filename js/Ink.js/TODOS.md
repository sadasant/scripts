### Ink.js TODOS
  

Clear background per animation, per figure
------------------------------------------

Clearing the background of a full resolution canvas
gets harder as the resolution is wider; the solution
is to carefully clear the box in which each object is drawn.

_Possible issues:_
-   Does this have to be made before drawing all the figures? I think yes.
-   Optional per figure or an option once Ink.js is initialized? (I think the second method is the wiser, but I could be wrong)
  

Redraw the canvas only when figure are moving
---------------------------------------------

Currently, the canvas is being refreshed no matter if all the figures
are still, this consumes resources unnecessarily, it could be better if
we managed to make it refresh only if a figure animation function is triggered,
or only if any of the attributes of possition or style are being changed.

_Possible issues:_
-   The analysis or operation detection should take less resources than the redraw itself.
  

Allow extensibility
-------------------

It could be very cool if we always keep an easy way to extend the library,
in order to build plugins for complex figures, but also keep a minimal core
as possible.

_Possible issues:_
-   This could mean to extend the basic prototype, or to reduce even more the basic prototype
    in order to make a separated plugin for animations and colission detection.
  

Text object and text over figures
---------------------------------

I would like to have shapes with text inside them, this could mean
making a general object or function to handle canvas texts, but also,
some extra animation conditionals and function to move the shape with the figures.
  

The alternative to use images all over the place
------------------------------------------------

In order to set background images per canvas or per shape, also, to move them or make
them interact with other object by collision handling and other animation functions.
  

Manage multiple canvas elements separatedly
-------------------------------------------

What if users want to handle many canvas? This could be interesting,
maybe it could need to build separated instances of itself per canvas.
  

Interact with object with the mouse (or touch)
----------------------------------------------

See the way to interact with object with the mouse or allow other kind
of interactions without consuming much memory and without harming the
behavior of uninteractive objects.
  
