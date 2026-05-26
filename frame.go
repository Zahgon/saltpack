// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

type headerOrFooterMarker string

const (
	headerMarker   headerOrFooterMarker = "BEGIN"
	footerMarker   headerOrFooterMarker = "END"
	maxFrameLength int                  = 512 // applies to header and footer
	maxBrandLength int                  = 128
)

func pop(v *([]string), n int) (ret []string) { _ = "STUB: not implemented"; return nil }

func shift(v *([]string), n int) (ret []string) { _ = "STUB: not implemented"; return nil }

func makeFrame(which headerOrFooterMarker, typ MessageType, brand string) string {
	_ = "STUB: not implemented"
	return ""
}

// MakeArmorHeader makes the armor header for the message type for the given "brand"
func MakeArmorHeader(typ MessageType, brand string) string { _ = "STUB: not implemented"; return "" }

// MakeArmorFooter makes the armor footer for the message type for the given "brand"
func MakeArmorFooter(typ MessageType, brand string) string { _ = "STUB: not implemented"; return "" }

func getStringForType(typ MessageType) string { _ = "STUB: not implemented"; return "" }

func parseFrame(m string, typ MessageType, hof headerOrFooterMarker) (brand string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// replace blocks of characters in the set [>\n\r\t ] with a single space, so that Go
// can easily parse each piece
