# Vortex

Vortex is a strongly-typed declarative programming language for graph databases.

# Language Spec 

Vortex supports `Vertex` and `Edge` connected by `Relation`s.

## Schema

A `Schema` defines a vertex. It can be compared with `struct` in other programming languages. `Schema`s consists of attributes which are strongly-typed.

This is how a schema named `Person` can be defined.

```sql
Schema Person {
    name string
    age  int
}
```

The type of an attribute can be one of the defined base types. Right now, it only supports `string`, `int` and `bool`.

The names can be composed of unicode letters. Each attribute has to be defined in a newline or separated by spaces.

## Vertex 

A `Vertex` is initialized from a `Schema`. It can be thought of as an `object` or `instance` of a `class` in other programming languages.

```sql
Vertex Jintu Person {
    .name = "Jintu"
    .age  = 20
}
```

A vertex uses the `.` identifier followed by the attribute name and a literal of the corrosponding type.

## Edge

An `Edge` is a link between two `Vertex`s. It does not have its own properties (yet) but connects vertices unidirectionally (`OneWay`) or bidirectionally (`TwoWay`).

An `Edge` can be created from one of the predefined base edge. The base edges is like a schema but does not hold any attributes.

```sql
Edge LivesIn OneWay
```

This creates a new edge `LivesIn` that is unidirectional from **left to right**.

The syntax is left simple with extensibility in mind. The obvious addition is to allow attributes in edges just like vertices.

## Relation

A relation binds two `Vertex`s with an `Edge`. If the edge type is `OneWay`, the relation is applied left to right.

Assuming `India` exist as a `Vertex`, you can write this:

```sql
Relation LivesIn {
    Jintu India
}
```

# Showcase

The combination of these entities gives you super-power to write complex graph queries very intuitively. Let's see a few examples of what we can do with it.

## Find all `Person` who have a one degree friend

```sql
Query Person {
    []FriendsWith Person
}
```

## Find all `Person` named "John" who have a mutual friend

A `Vertex` can be aliased for referencing later. This allows defining recursive relations.

```sql
Query Person as A { 
   .name = 'John'
   and []FriendsWith Person {
       FriendsWith A
   }
}
```

## Find all `Person` whose salary is strictly greater than all their friends' salary combined.

`()` nodes matches with any vertex and `[]()` matches with any edge.

```sql
Query Person as A { 
   Sum([]FriendsWith Person {
       []() ()
   }, .salary) < .salary
}
```

A subquery has its own scope. The `.salary` inside sum is scoped to the inner `Person` and the other `.salary` is scoped to the `Person` alised as `A`.

You can be explicit about it with `A.salary` though.

## Try to figure out the query below

```sql
Query Sum(Person as A {
    .name = "Hi"
    or (StartsWith(.name, 'H') and .age > 10)
    and [1..2]LivesIn () {
        [..]Within Country {
            .name="India"
            and []Within Continent
        }
    }
    and Sum([]FriendsWith Person {
        []LivesIn () {
            []Within Country{.name="USA"}
        }
    }, .salary) < .salary
}, .age)
```

The `[..]` syntax states that the edge can be there any number of times, including zero. The default `[]` evaluates to `[1]` to make the edge appear strictly once.
