// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
	"github.com/mitchellh/hashstructure"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Hash function
func (m *Extensions) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetConfigs() {
			innerHash.Reset()

			if h, ok := interface{}(v).(interface {
				Hash(innerHash hash.Hash64) (uint64, error)
			}); ok {
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if val, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if err := binary.Write(innerHash, binary.LittleEndian, val); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Extension) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	if h, ok := interface{}(m.GetConfig()).(interface {
		Hash(hasher hash.Hash64) (uint64, error)
	}); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetConfig(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
