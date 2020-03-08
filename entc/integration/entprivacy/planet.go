// Code generated by entc, DO NOT EDIT.

package entprivacy

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/entprivacy/planet"
)

// Planet is the model entity for the Planet schema.
type Planet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age uint `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlanetQuery when eager-loading is set.
	Edges PlanetEdges `json:"edges"`
}

// PlanetEdges holds the relations/edges for other nodes in the graph.
type PlanetEdges struct {
	// Neighbors holds the value of the neighbors edge.
	Neighbors []*Planet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NeighborsOrErr returns the Neighbors value or an error if the edge
// was not loaded in eager-loading.
func (e PlanetEdges) NeighborsOrErr() ([]*Planet, error) {
	if e.loadedTypes[0] {
		return e.Neighbors, nil
	}
	return nil, &NotLoadedError{edge: "neighbors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Planet) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullInt64{},  // age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Planet fields.
func (pl *Planet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(planet.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pl.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pl.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[1])
	} else if value.Valid {
		pl.Age = uint(value.Int64)
	}
	return nil
}

// QueryNeighbors queries the neighbors edge of the Planet.
func (pl *Planet) QueryNeighbors() *PlanetQuery {
	return (&PlanetClient{config: pl.config}).QueryNeighbors(pl)
}

// Update returns a builder for updating this Planet.
// Note that, you need to call Planet.Unwrap() before calling this method, if this Planet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Planet) Update() *PlanetUpdateOne {
	return (&PlanetClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pl *Planet) Unwrap() *Planet {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("entprivacy: Planet is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Planet) String() string {
	var builder strings.Builder
	builder.WriteString("Planet(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pl.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Planets is a parsable slice of Planet.
type Planets []*Planet

func (pl Planets) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
