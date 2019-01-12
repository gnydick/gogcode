# This is a very rough start to some g-code utilities, but...
The goal of this project is to create more and more g-code processing capabilities to help solve problems Makers have. Ultimately, the beginnings of the structs and utils could morph into a custom g-code generator. I've know idea how to write a slicer, but there is a lot that can be done with the base g-code to make very interesting things happen e.g. velocity painting.

**gogcode:** main directory, top level intellij project with util libs and main structs

**addZhop:** scans your g-code and detects travel movements and adds a z-hop around them so your print head doesn't scrape along the print

**ironLayer:** adds a duplicate pass for each layer without extruding so it "irons" the layer down*

**lapping:** this generates a routine to lap your nozzle. If you're concerned the tip of your nozzle may not be perfectly flat or it's dirty, this utility will generate orbital sander like patterns. Just tape some fine sandpaper to your bed, move the print head to Z0 or the appropriate position for your printer, and run the routine.

The above programs aren't very generalized yet, but hopefully when you browse the code, you can see where I'd like to take it.
