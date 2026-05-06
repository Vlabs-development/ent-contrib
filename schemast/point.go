// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schemast

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

type Point struct {
	Lat float64 `graphql:"lat"`
	Lon float64 `graphql:"lon"`
}

// Scan implements the Scanner interface.
func (p *Point) Scan(value any) error {
	bin, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid binary value for point")
	}
	var op orb.Point
	if err := wkb.Scanner(&op).Scan(bin[4:]); err != nil {
		return err
	}
	p.Lon, p.Lat = op.X(), op.Y()
	return nil
}

// Value implements the driver Valuer interface.
func (p Point) Value() (driver.Value, error) {
	op := orb.Point{p.Lon, p.Lat}
	return wkb.Value(op).Value()
}

// FormatParam implements the sql.ParamFormatter interface to tell the SQL
// builder that the placeholder for a Point parameter needs to be formatted.
func (p Point) FormatParam(placeholder string, info *sql.StmtInfo) string {
	if info.Dialect == dialect.MySQL {
		return "ST_GeomFromWKB(" + placeholder + ")"
	}
	return placeholder
}

// SchemaType defines the schema-type of the Point object.
func (Point) SchemaType() map[string]string {
	return map[string]string{
		dialect.MySQL: "POINT",
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (p *Point) UnmarshalGQL(v any) error {
	pt, ok := v.(map[string]any)
	if !ok {
		return fmt.Errorf("invalid format")
	}

	*p = Point{
		Lat: pt["lat"].(float64),
		Lon: pt["lon"].(float64),
	}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (p Point) MarshalGQL(w io.Writer) {
	b, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}
