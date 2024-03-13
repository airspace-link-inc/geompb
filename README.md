# geompb
A protocol buffer representation for geometries


## Rational
We have evaluated multiple solutions to transport geometries, but did not find any simple ones.  
Some use deltas, some don't support or don't exactly map to simple features.  
This is an attempt to not rebuild the exact same thing everywhere.

## Definition
```protobuf
message Geometry {
  Type type = 1;
  repeated Geometry geometries = 2;
  repeated Coordinates coordinates = 3;

  enum Type {
    TYPE_UNSPECIFIED = 0;
    TYPE_POINT = 1; // a 2d point: coordinates[0], 2 coordinates: longitude, latitude
    TYPE_POINTZ = 2; // a 3d point: coordinates[0], 3 coordinates: longitude, latitude, altitude
    TYPE_MULTIPOINT = 3; // multiple 2d points: coordinates is empty, use geometries with POINT types 
    TYPE_MULTIPOINTZ = 4; // multiple 3d points: coordinates is empty, use geometries with POINTZ types 
    TYPE_POLYGON = 5; // a 2d polygon: coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW 
    TYPE_POLYGONZ = 6; // a 3d polygon: coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW 
    TYPE_MULTIPOLYGON = 7; // multiple 2d polygons: coordinates is empty, use geometries with POLYGON types
    TYPE_MULTIPOLYGONZ = 8; // multiple 3d polygons: coordinates is empty, use geometries with POLYGONZ types
    TYPE_LINESTRING = 9; // a 2d linestring: coordinates[0], list of longitude, latitude
    TYPE_LINESTRINGZ = 10; // a 3d linestring: coordinates[0], list of longitude, latitude, altitude
    TYPE_MULTILINESTRING = 11; // multiple 2d linestrings: coordinates is empty, use geometries with LINESTRING types
    TYPE_MULTILINESTRINGZ = 12; // mulitple 3d linestrings: coordinates is empty, use geometries with LINESTRINGZ types
 }
}

message Coordinates {
  repeated double sequences = 1; // lng, lat or lng lat altitude
}
```

The definition basically follow simple features, the only notable difference is in Polygons, rings are not repeating the last point being the first.

## Generate The Code

```sh
buf generate
```

## Buf

This definitions is published to [buf](https://buf.build/airspacelink/geompb).

You can use it by letting buf import the package for you:

```yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
deps:
  - buf.build/airspacelink/geompb
```

Simply import the package in your own definition:

```protocolbuffer
	import "geompb/v1/geom.proto";
```

## Using with Go
In this repository you will see generated code for Go and helper functions to transform Geometry from and to a [Go geometry package](https://github.com/peterstace/simplefeatures).

```go
// Unmarshal from WKT using  peterstace/simplefeatures/
input := "POLYGON((0 0,0 1,1 1,1 0,0 0))"
g, _ := geom.UnmarshalWKT(input)

// Convert from Geom to PB
geotools.GeomToPB(g)

pbg :=  &geompb.Geometry{
	Type: geompb.Geometry_TYPE_POINT,
	Coordinates: []*geompb.Coordinates{
		{
			Sequences: []float64{-73.8677438649466, 40.6213222726945},
		},
	}
}

// Convert from PB to Geom
geotools.GeomToPB(pbg)
```
