Toy 3D renderer utilizing SDL, written in Go.

Not intended to be used for anything but learning, as all the rendering is done
pixel-by-pixel in software.

Written by building up incrementally from drawing a single pixel (creating a vertex),
to filling in all the pixels between two points (creating a line segment),
to connecting several line segments (creating a triangle),
to connecting multiple triangles (creating a rectangle, or any arbitrary polygon),
to connecting arbitrary polygons (creating the illusion of a 3D polygon).

## Examples

TODO
