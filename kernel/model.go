/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kernel

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/mowang666/rocketmq-client/rlog"
	"github.com/mowang666/rocketmq-client/utils"
)

// SendStatus of message
type SendStatus int

const (
	SendOK SendStatus = iota
	SendFlushDiskTimeout
	SendFlushSlaveTimeout
	SendSlaveNotAvailable

	FlagCompressed = 0x1
	MsgIdLength    = 8 + 8
)

// SendResult RocketMQ send result
type SendResult struct {
	Status        SendStatus
	MsgID         string
	MessageQueue  *MessageQueue
	QueueOffset   int64
	TransactionID string
	OffsetMsgID   string
	RegionID      string
	TraceOn       bool
}

// SendResult send message result to string(detail result)
func (result *SendResult) String() string {
	return fmt.Sprintf("SendResult [sendStatus=%d, msgIds=%s, offsetMsgId=%s, queueOffset=%d, messageQueue=%s]",
		result.Status, result.MsgID, result.OffsetMsgID, result.QueueOffset, result.MessageQueue.String())
}

// PullStatus pull status
type PullStatus int

// predefined pull status
const (
	PullFound PullStatus = iota
	PullNoNewMsg
	PullNoMsgMatched
	PullOffsetIllegal
	PullBrokerTimeout
)

// PullResult the pull result
type PullResult struct {
	NextBeginOffset      int64
	MinOffset            int64
	MaxOffset            int64
	Status               PullStatus
	SuggestWhichBrokerId int64
	messageExts          []*MessageExt
}

func (result *PullResult) GetMessageExts() []*MessageExt {
	return result.messageExts
}

func (result *PullResult) SetMessageExts(msgExts []*MessageExt) {
	result.messageExts = msgExts
}

func (result *PullResult) GetMessages() []*Message {
	if result.messageExts == nil || len(result.messageExts) == 0 {
		return make([]*Message, 0)
	}
	return toMessages(result.messageExts)
}

func (result *PullResult) String() string {
	return ""
}

func decodeMessage(data []byte) []*MessageExt {
	msgs := make([]*MessageExt, 0)
	buf := bytes.NewBuffer(data)
	count := 0
	for count < len(data) {
		msg := &MessageExt{}

		// 1. total size
		binary.Read(buf, binary.BigEndian, &msg.StoreSize)
		count += 4

		// 2. magic code
		buf.Next(4)
		count += 4

		// 3. body CRC32
		binary.Read(buf, binary.BigEndian, &msg.BodyCRC)
		count += 4

		// 4. queueID
		binary.Read(buf, binary.BigEndian, &msg.QueueId)
		count += 4

		// 5. Flag
		binary.Read(buf, binary.BigEndian, &msg.Flag)
		count += 4

		// 6. QueueOffset
		binary.Read(buf, binary.BigEndian, &msg.QueueOffset)
		count += 8

		// 7. physical offset
		binary.Read(buf, binary.BigEndian, &msg.CommitLogOffset)
		count += 8

		// 8. SysFlag
		binary.Read(buf, binary.BigEndian, &msg.SysFlag)
		count += 4

		// 9. BornTimestamp
		binary.Read(buf, binary.BigEndian, &msg.BornTimestamp)
		count += 8

		// 10. born host
		hostBytes := buf.Next(4)
		var port int32
		binary.Read(buf, binary.BigEndian, &port)
		msg.BornHost = fmt.Sprintf("%s:%d", utils.GetAddressByBytes(hostBytes), port)
		count += 8

		// 11. store timestamp
		binary.Read(buf, binary.BigEndian, &msg.StoreTimestamp)
		count += 8

		// 12. store host
		hostBytes = buf.Next(4)
		binary.Read(buf, binary.BigEndian, &port)
		msg.StoreHost = fmt.Sprintf("%s:%d", utils.GetAddressByBytes(hostBytes), port)
		count += 8

		// 13. reconsume times
		binary.Read(buf, binary.BigEndian, &msg.ReconsumeTimes)
		count += 4

		// 14. prepared transaction offset
		binary.Read(buf, binary.BigEndian, &msg.PreparedTransactionOffset)
		count += 8

		// 15. body
		var length int32
		binary.Read(buf, binary.BigEndian, &length)
		msg.Body = buf.Next(int(length))
		if (msg.SysFlag & FlagCompressed) == FlagCompressed {
			msg.Body = utils.UnCompress(msg.Body)
		}
		count += 4 + int(length)

		// 16. topic
		_byte, _ := buf.ReadByte()
		msg.Topic = string(buf.Next(int(_byte)))
		count += 1 + int(_byte)

		var propertiesLength int16
		binary.Read(buf, binary.BigEndian, &propertiesLength)
		if propertiesLength > 0 {
			msg.Properties = parseProperties(buf.Next(int(propertiesLength)))
		}
		count += 2 + int(propertiesLength)

		msg.MsgId = createMessageId(hostBytes, msg.CommitLogOffset)
		//count += 16

		msgs = append(msgs, msg)
	}

	return msgs
}

func createMessageId(addr []byte, offset int64) string {
	return "msgID" // TODO
}

func parseProperties(data []byte) map[string]string {
	return make(map[string]string, 0)
}

func toMessages(messageExts []*MessageExt) []*Message {
	msgs := make([]*Message, 0)

	return msgs
}

// MessageQueue message queue
type MessageQueue struct {
	Topic      string `json:"topic"`
	BrokerName string `json:"brokerName"`
	QueueId    int    `json:"queueId"`
}

func (mq *MessageQueue) String() string {
	return fmt.Sprintf("MessageQueue [topic=%s, brokerName=%s, queueId=%d]", mq.Topic, mq.BrokerName, mq.QueueId)
}

func (mq *MessageQueue) HashCode() int {
	result := 1
	result = 31*result + utils.HashString(mq.BrokerName)
	result = 31*result + mq.QueueId
	result = 31*result + utils.HashString(mq.Topic)

	return result
}

func (mq *MessageQueue) Equals(queue *MessageQueue) bool {
	// TODO
	return true
}

type FindBrokerResult struct {
	BrokerAddr    string
	Slave         bool
	BrokerVersion int32
}

type (
	// groupName of consumer
	producerData string

	consumeType string

	ServiceState int
)

const (
	StateCreateJust ServiceState = iota
	StateStartFailed
	StateRunning
	StateShutdown
)

type SubscriptionData struct {
	ClassFilterMode bool
	Topic           string
	SubString       string
	Tags            map[string]bool
	Codes           map[int32]bool
	SubVersion      int64
	ExpType         string
}

type consumerData struct {
	GroupName         string              `json:"groupName"`
	CType             consumeType         `json:"consumeType"`
	MessageModel      string              `json:"messageModel"`
	Where             string              `json:"consumeFromWhere"`
	SubscriptionDatas []*SubscriptionData `json:"subscriptionDataSet"`
	UnitMode          bool                `json:"unitMode"`
}

type heartbeatData struct {
	ClientId      string         `json:"clientID"`
	ProducerDatas []producerData `json:"producerDataSet"`
	ConsumerDatas []consumerData `json:"consumerDataSet"`
}

func (data *heartbeatData) encode() []byte {
	d, err := json.Marshal(data)
	if err != nil {
		rlog.Errorf("marshal heartbeatData error: %s", err.Error())
		return nil
	}
	rlog.Info(string(d))
	return d
}
