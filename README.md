This is a very rough start to some g-code utilities.

gogcode - main directory, top level intellij project

addZhop - scans your g-code and detects travel movements and adds a z-hop around them so your print head doesn't scrape along the print

ironLayer - adds a duplicate pass for each layer without extruding so it "irons" the layer down

lapping - this generates a routine to lap your nozzle. If you're concerned the tip of your nozzle may not be perfectly flat or it's dirty, this utility will generate orbital sander like patterns. Just tape some fine sandpaper to your bed, move the print head to Z0 or the appropriate position for your printer, and run the routine.
