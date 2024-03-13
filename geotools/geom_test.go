package geotools

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/peterstace/simplefeatures/geom"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/airspace-link-inc/geompb/gen/go/geompb/v1"
)

func TestGeomToPB(t *testing.T) {
	tests := []struct {
		name string
		wkt  string
		want *geompb.Geometry
	}{
		{
			"a point in NYC",
			"POINT(-73.8677438649466 40.6213222726945)",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_POINT,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{-73.8677438649466, 40.6213222726945},
					},
				},
			},
		},

		{
			"a point Z at 200  in NYC",
			"POINT Z (-73.8677438649466 40.6213222726945 200)",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_POINTZ,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{-73.8677438649466, 40.6213222726945, 200},
					},
				},
			},
		},

		{
			"two points NYC & Toronto as multipoint",
			"MULTIPOINT((-79.41376755607581 43.70810361067444),(-73.99607041999752 40.73124501446793))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTIPOINT,
				Geometries: []*geompb.Geometry{
					{
						Type: geompb.Geometry_TYPE_POINT,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{-79.41376755607581, 43.70810361067444},
							},
						},
					},
					{
						Type: geompb.Geometry_TYPE_POINT,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{-73.99607041999752, 40.73124501446793},
							},
						},
					},
				},
			},
		},

		{
			"a linestring from NYC to Toronto",
			"LINESTRING(-73.8677438649466 40.6213222726945,-79.70528594465897 43.80736006424203)",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_LINESTRING,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{-73.8677438649466, 40.6213222726945, -79.70528594465897, 43.80736006424203},
					},
				},
			},
		},

		{
			"a linestring from NYC to Toronto with Z 400",
			"LINESTRING Z (-73.8677438649466 40.6213222726945 400,-79.70528594465897 43.80736006424203 400)",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_LINESTRINGZ,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{-73.8677438649466, 40.6213222726945, 400, -79.70528594465897, 43.80736006424203, 400},
					},
				},
			},
		},

		{
			"two linestring in costa rica as multilinestring",
			"MULTILINESTRING((-84.14407289771565 9.96284871125465,-83.91011520366128 9.951092998161073,-83.93048248205758 10.233833731929877),(-84.13263094264897 9.860770781545611,-83.88992876201335 9.843089602092192,-83.92361196891666 9.537433160168405))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTILINESTRING,
				Geometries: []*geompb.Geometry{
					{
						Type: geompb.Geometry_TYPE_LINESTRING,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{-84.14407289771565, 9.96284871125465, -83.91011520366128, 9.951092998161073, -83.93048248205758, 10.2338337319298770},
							},
						},
					},
					{
						Type: geompb.Geometry_TYPE_LINESTRING,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{-84.13263094264897, 9.860770781545611, -83.88992876201335, 9.843089602092192, -83.92361196891666, 9.537433160168405},
							},
						},
					},
				},
			},
		},

		{
			"a valid polygon no holes around Orleans, France",
			"POLYGON((1.8743786640940527 47.93047756294925,1.8743786640940527 47.87658239261785,1.9696535783850493 47.87658239261785,1.9696535783850493 47.93047756294925,1.8743786640940527 47.93047756294925))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{1.8743786640940527, 47.93047756294925, 1.8743786640940527, 47.87658239261785, 1.9696535783850493, 47.87658239261785, 1.9696535783850493, 47.93047756294925},
					},
				},
			},
		},

		{
			"A valid polygon with a hole in Africa",
			"POLYGON((11.4652180932167 12.195518360028313,11.330778012973777 -2.834688304275655,11.330778012973777 -2.834688304275655,11.330778012973777 -2.834688304275655,24.774786037265972 -3.9083316114949413,24.505905876780123 12.326891775851083,11.4652180932167 12.195518360028313),(18.5233223059701 6.894362482583444,20.405483429371003 3.1451573972067783,20.405483429371003 3.1451573972067783,13.81791949746783 3.8161139779436444,18.5233223059701 6.894362482583444))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{11.4652180932167, 12.195518360028313, 11.330778012973777, -2.834688304275655, 11.330778012973777, -2.834688304275655, 11.330778012973777, -2.834688304275655, 24.774786037265972, -3.9083316114949413, 24.505905876780123, 12.326891775851083},
					},
					{
						Sequences: []float64{18.5233223059701, 6.894362482583444, 20.405483429371003, 3.1451573972067783, 20.405483429371003, 3.1451573972067783, 13.81791949746783, 3.8161139779436444},
					},
				},
			},
		},

		{
			"A valid polygon with a hole in Africa and a Z of 200",
			"POLYGON Z ((11.4652180932167 12.195518360028313 200,11.330778012973777 -2.834688304275655 200,11.330778012973777 -2.834688304275655 200,11.330778012973777 -2.834688304275655 200,24.774786037265972 -3.9083316114949413 200,24.505905876780123 12.326891775851083 200,11.4652180932167 12.195518360028313 200),(18.5233223059701 6.894362482583444 200,20.405483429371003 3.1451573972067783 200,20.405483429371003 3.1451573972067783 200,13.81791949746783 3.8161139779436444 200,18.5233223059701 6.894362482583444 200))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGONZ,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{11.4652180932167, 12.195518360028313, 200, 11.330778012973777, -2.834688304275655, 200, 11.330778012973777, -2.834688304275655, 200, 11.330778012973777, -2.834688304275655, 200, 24.774786037265972, -3.9083316114949413, 200, 24.505905876780123, 12.326891775851083, 200},
					},
					{
						Sequences: []float64{18.5233223059701, 6.894362482583444, 200, 20.405483429371003, 3.1451573972067783, 200, 20.405483429371003, 3.1451573972067783, 200, 13.81791949746783, 3.8161139779436444, 200},
					},
				},
			},
		},

		{
			"a multipolygon, one polygon with a hole and one without around africa",
			"MULTIPOLYGON(((40 40,20 45,45 30,40 40)),((20 35,10 30,10 10,30 5,45 20,20 35),(30 20,20 15,20 25,30 20)))",
			&geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTIPOLYGON,
				Geometries: []*geompb.Geometry{
					{
						Type: geompb.Geometry_TYPE_POLYGON,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{40, 40, 20, 45, 45, 30},
							},
						},
					},
					{
						Type: geompb.Geometry_TYPE_POLYGON,
						Coordinates: []*geompb.Coordinates{
							{
								Sequences: []float64{20, 35, 10, 30, 10, 10, 30, 5, 45, 20},
							},
							{
								Sequences: []float64{30, 20, 20, 15, 20, 25},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := geom.UnmarshalWKT(tt.wkt)
			require.NoError(t, err)

			pb := GeomToPB(g)
			if diff := cmp.Diff(pb, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("unexpected difference in protocol buffer:\n%v", diff)
			}

			// test back
			g2, err := PBToGeom(pb)
			require.NoError(t, err)

			require.Equal(t, tt.wkt, g2.AsText())
		})
	}
}

func TestPBToGeom(t *testing.T) {
	tests := []struct {
		name      string
		geom      geompb.Geometry
		wantError bool
	}{
		{
			"a point with only one coordinate",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_POINT,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{-73.867743864946},
					},
				},
			},
			true,
		},

		{
			"a LINESTRING with only one coordinate",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_LINESTRING,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{2},
					},
				},
			},
			true,
		},

		{
			"a LINESTRING 2d with 3 coordinates",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_LINESTRING,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{2, 2, 400},
					},
				},
			},
			true,
		},

		{
			"a LINESTRINGZ 3d with 2 coordinates",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_LINESTRINGZ,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{2, 2},
					},
				},
			},
			true,
		},

		{
			"a POLYGON with 2 points",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{2, 2, 1, 1},
					},
				},
			},
			true,
		},

		{
			"a POLYGON with 3 coordinates",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{2, 2, 1},
					},
				},
			},
			true,
		},

		{
			"a CW POLYGON",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{
					{
						Sequences: []float64{139.60639440628944, 35.75578402827527, 139.80886784025313, 35.75578402827527, 139.80886784025313, 35.57437067701714, 139.60639440628944, 35.57437067701714, 139.60639440628944, 35.75578402827527},
					},
				},
			},
			true,
		},

		{
			"an empty POINT",
			geompb.Geometry{
				Type:        geompb.Geometry_TYPE_POINT,
				Coordinates: []*geompb.Coordinates{},
			},
			true,
		},

		{
			"an empty LINESTRING",
			geompb.Geometry{
				Type:        geompb.Geometry_TYPE_LINESTRING,
				Coordinates: []*geompb.Coordinates{},
			},
			true,
		},

		{
			"an empty POLYGON",
			geompb.Geometry{
				Type:        geompb.Geometry_TYPE_POLYGON,
				Coordinates: []*geompb.Coordinates{},
			},
			true,
		},

		{
			"a multi point with a pointz",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTIPOINT,
				Geometries: []*geompb.Geometry{
					{
						Type:        geompb.Geometry_TYPE_POINTZ,
						Coordinates: []*geompb.Coordinates{{Sequences: []float64{1, 2}}},
					},
				},
			},
			true,
		},

		{
			"a multi linestring with a pointz",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTILINESTRING,
				Geometries: []*geompb.Geometry{
					{
						Type:        geompb.Geometry_TYPE_POINTZ,
						Coordinates: []*geompb.Coordinates{{Sequences: []float64{1, 2}}},
					},
				},
			},
			true,
		},

		{
			"a multi polygon with a pointz",
			geompb.Geometry{
				Type: geompb.Geometry_TYPE_MULTIPOLYGON,
				Geometries: []*geompb.Geometry{
					{
						Type:        geompb.Geometry_TYPE_POINTZ,
						Coordinates: []*geompb.Coordinates{{Sequences: []float64{1, 2}}},
					},
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PBToGeom(&tt.geom)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected to receive an error")
				}
			} else {
				if err != nil {
					t.Fatal("should not fail")
				}
			}
		})
	}
}
