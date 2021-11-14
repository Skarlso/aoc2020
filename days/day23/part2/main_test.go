package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	first := &cup{
		label: 1,
		next: &cup{
			label: 2,
			next: &cup{
				label: 4,
				next: &cup{
					label: 5,
				},
			},
		},
	}
	first.next.next.next.next = first
	result := first.Display(4)
	assert.Equal(t, "512", result)
}
