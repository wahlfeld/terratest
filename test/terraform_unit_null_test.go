package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wahlfeld/terratest/modules/terraform"
)

func TestUnitNullInput(t *testing.T) {
	t.Parallel()

	foo := map[string]interface{}{
		"nullable_string":    nil,
		"nonnullable_string": "foo",
	}
	options := &terraform.Options{
		TerraformDir: "./fixtures/terraform-null",
		Vars:         map[string]interface{}{"foo": foo},
	}
	terraform.InitAndApply(t, options)

	fooOut := terraform.OutputMap(t, options, "foo")
	assert.Equal(t, fooOut, map[string]string{"nonnullable_string": "foo", "nullable_string": "<nil>"})

	barOut := terraform.Output(t, options, "bar")
	assert.Equal(t, barOut, "I AM NULL")
}
