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
    TYPE_POINT = 1; // coordinates[0], 2 coordinates: longitude, latitude
    TYPE_POINTZ = 2; // coordinates[0], 3 coordinates: longitude, latitude, altitude
    TYPE_MULTIPOINT = 3; // coordinates is empty, use geometries with POINT types 
    TYPE_MULTIPOINTZ = 4; // coordinates is empty, use geometries with POINTZ types 
    TYPE_POLYGON = 5; // coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW 
    TYPE_POLYGONZ = 6; // coordinates[0] is the outer loop must be CCW, first point is automatically the last point as well, coordinates[0:] are holes must be CW 
    TYPE_MULTIPOLYGON = 7; // coordinates is empty, use geometries with POLYGON types
    TYPE_MULTIPOLYGONZ = 8; // coordinates is empty, use geometries with POLYGONZ types
    TYPE_LINESTRING = 9; // coordinates[0], list of longitude, latitude
    TYPE_LINESTRINGZ = 10; // coordinates[0], list of longitude, latitude, altitude
    TYPE_MULTILINESTRING = 11; // coordinates is empty, use geometries with LINESTRING types
    TYPE_MULTILINESTRINGZ = 12; // coordinates is empty, use geometries with LINESTRINGZ types
 }
}

message Coordinates {
  repeated double sequences = 1; // lng, lat or lng lat altitude
}
```

## Using with Go
In this repository you will see generated code for Go and helper functions to transform Geometry from and to a [Go geometry package](https://github.com/peterstace/simplefeatures).
