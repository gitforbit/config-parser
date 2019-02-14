// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *Bind) Init() {
	p.data = []types.Bind{}
}

func (p *Bind) GetParserName() string {
	return "bind"
}

func (p *Bind) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Bind) GetOne(index int) (common.ParserData, error) {
	if len(p.data) == 0 {
		return nil, errors.FetchError
	}
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *Bind) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Bind:
		p.data = newValue
	case *types.Bind:
		if index > -1 {
			p.data = append(p.data, types.Bind{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.Bind:
		if index > -1 {
			p.data = append(p.data, types.Bind{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (p *Bind) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "bind" {
		data, err := p.parse(line, parts, comment)
		if err != nil {
			return "", &errors.ParseError{Parser: "Bind", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Bind", Line: line}
}