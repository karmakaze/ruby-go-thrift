// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Flop struct {
	A int16 `thrift:"a,1" json:"a"`
	B int16 `thrift:"b,2" json:"b"`
}

func NewFlop() *Flop {
	return &Flop{}
}

func (p *Flop) GetA() int16 {
	return p.A
}

func (p *Flop) GetB() int16 {
	return p.B
}
func (p *Flop) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Flop) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.A = v
	}
	return nil
}

func (p *Flop) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.B = v
	}
	return nil
}

func (p *Flop) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("flop"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Flop) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("a", thrift.I16, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:a: %s", p, err)
	}
	if err := oprot.WriteI16(int16(p.A)); err != nil {
		return fmt.Errorf("%T.a (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:a: %s", p, err)
	}
	return err
}

func (p *Flop) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("b", thrift.I16, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:b: %s", p, err)
	}
	if err := oprot.WriteI16(int16(p.B)); err != nil {
		return fmt.Errorf("%T.b (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:b: %s", p, err)
	}
	return err
}

func (p *Flop) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Flop(%+v)", *p)
}

type ExampleException struct {
}

func NewExampleException() *ExampleException {
	return &ExampleException{}
}

func (p *ExampleException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ExampleException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ExampleException"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ExampleException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExampleException(%+v)", *p)
}

func (p *ExampleException) Error() string {
	return p.String()
}
