# POCSAG Encoding for Go

**WORK IN PROGRESS - CURRENTLY FAILING TESTS**

This library encodes a given POCSAG pager CAPCODE and text message into the over the air format.

Ported in part and with thanks from https://github.com/faithanalog/pocsag-encoder

## Usage

This library exposes one function that accepts the CAPCODE and text message, and returns a byte array of the encoded message.

`EncodeTransmission(address int, message string) []byte`

