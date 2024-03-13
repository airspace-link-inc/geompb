package geotools

import (
	"fmt"

	"github.com/airspace-link-inc/geompb/gen/go/geompb/v1"
	"github.com/peterstace/simplefeatures/geom"
)

// GeomToPB serializes a Geometry suitable for Protocol Buffer
// returns nil on unknown geometries
func GeomToPB(g geom.Geometry) *geompb.Geometry {
	var gpb *geompb.Geometry

	switch g.Type() {
	case geom.TypePoint:
		gpb = PointToPB(g.MustAsPoint())
	case geom.TypeMultiPoint:
		gpb = &geompb.Geometry{
			Type: geompb.Geometry_TYPE_MULTIPOINT,
		}
		mp, _ := g.AsMultiPoint()
		gpb.Geometries = make([]*geompb.Geometry, mp.NumPoints())
		for i := 0; i < mp.NumPoints(); i++ {
			gpb.Geometries[i] = PointToPB(mp.PointN(i))
		}
		// test if we have 3 dimensions
		if mp.CoordinatesType() == geom.DimXYZ {
			gpb.Type = geompb.Geometry_TYPE_MULTIPOINTZ
		}
	case geom.TypeLineString:
		gpb = LineStringToPB(g.MustAsLineString())
	case geom.TypeMultiLineString:
		gpb = &geompb.Geometry{
			Type: geompb.Geometry_TYPE_MULTILINESTRING,
		}
		mls, _ := g.AsMultiLineString()
		gpb.Geometries = make([]*geompb.Geometry, mls.NumLineStrings())
		for i := 0; i < mls.NumLineStrings(); i++ {
			gpb.Geometries[i] = LineStringToPB(mls.LineStringN(i))
		}
		// test if we have 3 dimensions
		if mls.CoordinatesType() == geom.DimXYZ {
			gpb.Type = geompb.Geometry_TYPE_MULTILINESTRINGZ
		}
	case geom.TypePolygon:
		gpb = PolygonToPB(g.MustAsPolygon())
	case geom.TypeMultiPolygon:
		gpb = &geompb.Geometry{
			Type: geompb.Geometry_TYPE_MULTIPOLYGON,
		}
		mp, _ := g.AsMultiPolygon()
		gpb.Geometries = make([]*geompb.Geometry, mp.NumPolygons())

		for i := 0; i < mp.NumPolygons(); i++ {
			gpb.Geometries[i] = PolygonToPB(mp.PolygonN(i))
		}
		// test if we have 3 dimensions
		if mp.CoordinatesType() == geom.DimXYZ {
			gpb.Type = geompb.Geometry_TYPE_MULTIPOLYGONZ
		}
	default:
		return nil
	}

	return gpb
}

func LineStringToPB(l geom.LineString) *geompb.Geometry {
	gpb := &geompb.Geometry{}

	gseq := l.Coordinates()

	// testing for 2D or 3D
	switch l.CoordinatesType() {
	case geom.DimXY:
		gpb.Type = geompb.Geometry_TYPE_LINESTRING

		seq := make([]float64, gseq.Length()*2)

		for j := 0; j < gseq.Length(); j++ {
			xy := gseq.GetXY(j)
			seq[j*2] = xy.X
			seq[j*2+1] = xy.Y
		}
		gpb.Coordinates = []*geompb.Coordinates{{Sequences: seq}}
	case geom.DimXYZ:
		gpb.Type = geompb.Geometry_TYPE_LINESTRINGZ

		seq := make([]float64, gseq.Length()*3)

		for j := 0; j < gseq.Length(); j++ {
			xyz := gseq.Get(j)
			seq[j*3] = xyz.X
			seq[j*3+1] = xyz.Y
			seq[j*3+2] = xyz.Z
		}
		gpb.Coordinates = []*geompb.Coordinates{{Sequences: seq}}
	}

	return gpb
}

func PointToPB(p geom.Point) *geompb.Geometry {
	gpb := &geompb.Geometry{}

	// testing for 2D or 3D
	switch p.CoordinatesType() {
	case geom.DimXY:
		gpb.Type = geompb.Geometry_TYPE_POINT

		xy, _ := p.XY()

		gpb.Coordinates = []*geompb.Coordinates{{
			Sequences: []float64{xy.X, xy.Y},
		}}
	case geom.DimXYZ:
		gpb.Type = geompb.Geometry_TYPE_POINTZ

		xyz := p.DumpCoordinates().Get(0)

		gpb.Coordinates = []*geompb.Coordinates{{
			Sequences: []float64{xyz.X, xyz.Y, xyz.Z},
		}}
	}

	return gpb
}

// PolygonToPB serializes a polygon suitable for Protocol Buffer
func PolygonToPB(p geom.Polygon) *geompb.Geometry {
	gpb := &geompb.Geometry{}

	// Force orientation
	p = p.ForceCCW()

	seqs := p.Coordinates()

	gpb.Coordinates = make([]*geompb.Coordinates, len(seqs))

	// testing for 2D or 3D
	switch p.CoordinatesType() {
	case geom.DimXY:
		gpb.Type = geompb.Geometry_TYPE_POLYGON

		for i, ring := range seqs {
			seq := make([]float64, ring.Length()*2-2)

			for j := 0; j < ring.Length()-1; j++ {
				xy := ring.GetXY(j)
				seq[j*2] = xy.X
				seq[j*2+1] = xy.Y
			}

			gpb.Coordinates[i] = &geompb.Coordinates{Sequences: seq}
		}
	case geom.DimXYZ:
		gpb.Type = geompb.Geometry_TYPE_POLYGONZ

		for i, ring := range seqs {
			seq := make([]float64, ring.Length()*3-3)

			for j := 0; j < ring.Length()-1; j++ {
				xyz := ring.Get(j)
				seq[j*3] = xyz.X
				seq[j*3+1] = xyz.Y
				seq[j*3+2] = xyz.Z
			}

			gpb.Coordinates[i] = &geompb.Coordinates{Sequences: seq}
		}

	default:
		return nil
	}

	return gpb
}

func PBToGeom(pbg *geompb.Geometry) (geom.Geometry, error) {
	// verify proper length for types with coordinates
	switch pbg.Type {
	case geompb.Geometry_TYPE_POINT, geompb.Geometry_TYPE_POINTZ,
		geompb.Geometry_TYPE_LINESTRING, geompb.Geometry_TYPE_LINESTRINGZ,
		geompb.Geometry_TYPE_POLYGON, geompb.Geometry_TYPE_POLYGONZ:

		if len(pbg.GetCoordinates()) == 0 {
			return geom.Geometry{}, fmt.Errorf("empty coordinates")
		}
	}

	// ensuring a multi container only contains the same kind
	switch pbg.Type {
	case geompb.Geometry_TYPE_MULTIPOINT, geompb.Geometry_TYPE_MULTIPOINTZ,
		geompb.Geometry_TYPE_MULTILINESTRING, geompb.Geometry_TYPE_MULTILINESTRINGZ,
		geompb.Geometry_TYPE_MULTIPOLYGON, geompb.Geometry_TYPE_MULTIPOLYGONZ:
		for i, p := range pbg.Geometries {
			// ensuring a multi container only contains the same kind
			if pbg.Type == geompb.Geometry_TYPE_MULTIPOINT && p.Type != geompb.Geometry_TYPE_POINT {
				return geom.Geometry{}, fmt.Errorf("not a POINT type in a MULTIPOINT #%d", i)
			}
			if pbg.Type == geompb.Geometry_TYPE_MULTIPOINTZ && p.Type != geompb.Geometry_TYPE_POINTZ {
				return geom.Geometry{}, fmt.Errorf("not a POINTZ type in a MULTIPOINTZ #%d", i)
			}
			if pbg.Type == geompb.Geometry_TYPE_MULTILINESTRING && p.Type != geompb.Geometry_TYPE_LINESTRING {
				return geom.Geometry{}, fmt.Errorf("not a LINESTRING type in a MULTILINESTRING #%d", i)
			}
			if pbg.Type == geompb.Geometry_TYPE_MULTILINESTRINGZ && p.Type != geompb.Geometry_TYPE_LINESTRINGZ {
				return geom.Geometry{}, fmt.Errorf("not a LINESTRING type in a MULTILINESTRINGZ #%d", i)
			}
			if pbg.Type == geompb.Geometry_TYPE_MULTIPOLYGON && p.Type != geompb.Geometry_TYPE_POLYGON {
				return geom.Geometry{}, fmt.Errorf("not a POLYGON type in a MULTIPOLYGON #%d", i)
			}
			if pbg.Type == geompb.Geometry_TYPE_MULTIPOLYGONZ && p.Type != geompb.Geometry_TYPE_POLYGONZ {
				return geom.Geometry{}, fmt.Errorf("not a POLYGONZ type in a MULTIPOLYGONZ #%d", i)
			}
		}
	}

	switch pbg.Type {
	case geompb.Geometry_TYPE_POINT:
		p, err := pointPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return p.AsGeometry(), nil
	case geompb.Geometry_TYPE_POINTZ:
		p, err := pointZPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return p.AsGeometry(), nil
	case geompb.Geometry_TYPE_LINESTRING:
		ls, err := linestringPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return ls.AsGeometry(), nil
	case geompb.Geometry_TYPE_LINESTRINGZ:
		ls, err := linestringZPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return ls.AsGeometry(), nil
	case geompb.Geometry_TYPE_POLYGON:
		p, err := polygonPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return p.AsGeometry(), nil
	case geompb.Geometry_TYPE_POLYGONZ:
		p, err := polygonZPBToGeom(pbg)
		if err != nil {
			return geom.Geometry{}, err
		}

		return p.AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTIPOINT:
		mp := make([]geom.Point, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := pointPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mp[i] = p
		}
		return geom.NewMultiPoint(mp).AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTIPOINTZ:
		mp := make([]geom.Point, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := pointZPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mp[i] = p
		}
		return geom.NewMultiPoint(mp).AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTILINESTRING:
		mls := make([]geom.LineString, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := linestringPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mls[i] = p
		}
		return geom.NewMultiLineString(mls).AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTILINESTRINGZ:
		mls := make([]geom.LineString, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := linestringZPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mls[i] = p
		}
		return geom.NewMultiLineString(mls).AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTIPOLYGON:
		mp := make([]geom.Polygon, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := polygonPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mp[i] = p
		}
		return geom.NewMultiPolygon(mp).AsGeometry(), nil
	case geompb.Geometry_TYPE_MULTIPOLYGONZ:
		mp := make([]geom.Polygon, len(pbg.Geometries))
		for i, pg := range pbg.Geometries {
			p, err := polygonZPBToGeom(pg)
			if err != nil {
				return geom.Geometry{}, err
			}
			mp[i] = p
		}
		return geom.NewMultiPolygon(mp).AsGeometry(), nil
	default:
		return geom.NewEmptyPoint(geom.DimXY).AsGeometry(), nil
	}
}

func pointPBToGeom(pbg *geompb.Geometry) (geom.Point, error) {
	seq := pbg.GetCoordinates()[0].GetSequences()
	if len(seq) != 2 {
		return geom.Point{}, fmt.Errorf("sequence length for POINT is not 2")
	}

	return geom.NewPoint(geom.Coordinates{
		XY: geom.XY{
			X: seq[0],
			Y: seq[1],
		},
	}), nil
}

func pointZPBToGeom(pbg *geompb.Geometry) (geom.Point, error) {
	seq := pbg.GetCoordinates()[0].GetSequences()
	if len(seq) != 3 {
		return geom.Point{}, fmt.Errorf("sequence length for POINTZ is not 3")
	}

	return geom.NewPoint(geom.Coordinates{
		XY: geom.XY{
			X: seq[0],
			Y: seq[1],
		},
		Z:    seq[2],
		Type: geom.DimXYZ,
	}), nil
}

func linestringPBToGeom(pbg *geompb.Geometry) (geom.LineString, error) {
	seq := pbg.GetCoordinates()[0].GetSequences()
	if len(seq) == 0 || len(seq) < 2 || len(seq)%2 != 0 {
		return geom.LineString{}, fmt.Errorf("invalid sequence for LINESTRING")
	}

	return geom.NewLineString(geom.NewSequence(seq, geom.DimXY)), nil
}

func linestringZPBToGeom(pbg *geompb.Geometry) (geom.LineString, error) {
	seq := pbg.GetCoordinates()[0].GetSequences()
	if len(seq) == 0 || len(seq) < 3 || len(seq)%3 != 0 {
		return geom.LineString{}, fmt.Errorf("invalid sequence for LINESTRINGZ")
	}

	return geom.NewLineString(geom.NewSequence(seq, geom.DimXYZ)), nil
}

func polygonPBToGeom(pbg *geompb.Geometry) (geom.Polygon, error) {
	ls := make([]geom.LineString, len(pbg.GetCoordinates()))

	for i, coords := range pbg.GetCoordinates() {
		// test we have at least 3 points
		seq := coords.GetSequences()
		if len(seq) == 0 || len(seq) < 6 {
			return geom.Polygon{}, fmt.Errorf("invalid sequence for POLYGON at ring #%d", i)
		}

		if len(seq)%2 != 0 {
			return geom.Polygon{}, fmt.Errorf("invalid number of doubles in sequence for POLYGON at ring #%d", i)
		}

		// we need one more coordinate seq to close the loop (first point is last point)
		bseq := append(seq, make([]float64, 2)...)

		bseq[len(bseq)-2] = seq[0]
		bseq[len(bseq)-1] = seq[1]

		ls[i] = geom.NewLineString(geom.NewSequence(bseq, geom.DimXY))
	}

	p := geom.NewPolygon(ls)

	if !p.IsCCW() {
		return geom.Polygon{}, fmt.Errorf("polygon is not CCW")
	}

	return p, nil
}

func polygonZPBToGeom(pbg *geompb.Geometry) (geom.Polygon, error) {
	ls := make([]geom.LineString, len(pbg.GetCoordinates()))

	for i, coords := range pbg.GetCoordinates() {
		// test we have at least 3 points
		seq := coords.GetSequences()
		if len(seq) == 0 || len(seq) < 9 {
			return geom.Polygon{}, fmt.Errorf("invalid sequence for POLYGONZ at ring #%d", i)
		}

		if len(seq)%3 != 0 {
			return geom.Polygon{}, fmt.Errorf("invalid number of doubles in sequence for POLYGON at ring #%d", i)
		}

		// we need one more coordinate seq to close the loop (first point is last point)
		bseq := append(seq, make([]float64, 3)...)

		bseq[len(bseq)-3] = seq[0]
		bseq[len(bseq)-2] = seq[1]
		bseq[len(bseq)-1] = seq[2]

		ls[i] = geom.NewLineString(geom.NewSequence(bseq, geom.DimXYZ))
	}

	p := geom.NewPolygon(ls)

	return p, nil
}
