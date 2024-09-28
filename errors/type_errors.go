package errors

import "errors"

var (
  InvalidOperation = errors.New("Unsupported operation between operands")
  UnknownType = errors.New("Operand is of unknown type")
  UnknownField = errors.New("Unknown field")
  UnhomogeneousType = errors.New("List elements are of different types")
)
