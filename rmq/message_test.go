// Copyright © 2017 Meltwater
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

package rmq

import (
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

var (
	myStringHeader   = "RABBITIO.amqp.headers.string.myStringHeader"
	myStringEqHeader = "RABBITIO.amqp.headers.string.myStringEqHeader"
	myInt32Header    = "RABBITIO.amqp.headers.int.myInt32Header"
	myInt64Header    = "RABBITIO.amqp.headers.int.myInt64Header"
	myFloat32Header  = "RABBITIO.amqp.headers.float.myFloat32Header"
	myFloat64Header  = "RABBITIO.amqp.headers.float.myFloat64Header"
	myBoolHeader     = "RABBITIO.amqp.headers.bool.myBoolHeader"
)

func TestToPAXRecords(t *testing.T) {
	messageHeaders := make(amqp.Table)
	messageHeaders["myStringHeader"] = "myString"
	messageHeaders["myStringEqHeader"] = "my=String"
	messageHeaders["myInt32Header"] = int32(32)
	messageHeaders["myInt64Header"] = int64(64)
	messageHeaders["myFloat32Header"] = float32(32.32)
	messageHeaders["myFloat64Header"] = float64(64.64)
	messageHeaders["myBoolHeader"] = true
	message := &Message{Headers: messageHeaders}

	var attrHeaders = make(map[string]string)
	attrHeaders[myStringHeader] = "myString"
	attrHeaders[myStringEqHeader] = "my=String"
	attrHeaders[myInt32Header] = "32"
	attrHeaders[myInt64Header] = "64"
	attrHeaders[myFloat32Header] = "32.32"
	attrHeaders[myFloat64Header] = "64.64"
	attrHeaders[myBoolHeader] = "true"

	pax := message.ToPAXRecords()

	assert.NoError(t, messageHeaders.Validate(), "should be valid Headers")

	assert.Equal(t, attrHeaders[myStringHeader], pax[myStringHeader])
	assert.Equal(t, attrHeaders[myStringEqHeader], pax[myStringEqHeader])
	assert.Equal(t, attrHeaders[myInt32Header], pax[myInt32Header])
	assert.Equal(t, attrHeaders[myInt64Header], pax[myInt64Header])
	assert.Equal(t, attrHeaders[myFloat32Header], pax[myFloat32Header])
	assert.Equal(t, attrHeaders[myFloat64Header], pax[myFloat64Header])
	assert.Equal(t, attrHeaders[myBoolHeader], pax[myBoolHeader])
}

func TestNewMessage(t *testing.T) {
	var headers = make(map[string]string)
	headers["RABBITIO.amqp.routingkey"] = "routingKey from tarball PAXRecords"
	headers[myStringHeader] = "myString"
	headers[myStringEqHeader] = "my=String"
	headers[myInt32Header] = "3232"
	headers[myInt64Header] = "6464"
	headers[myFloat32Header] = "32.123"
	headers[myFloat64Header] = "64.123"
	headers[myBoolHeader] = "true"

	m := NewMessage([]byte("Message"), headers)

	assert.Equal(t, "routingKey from tarball PAXRecords", m.RoutingKey)
	assert.Equal(t, []byte("Message"), m.Body)
	assert.NoError(t, m.Headers.Validate())
}
