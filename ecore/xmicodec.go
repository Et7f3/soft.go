// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type XMICodec struct {
}

func (d *XMICodec) NewEncoder(options map[string]interface{}) EResourceEncoder {
	return NewXMIEncoder(options)
}
func (d *XMICodec) NewDecoder(options map[string]interface{}) EResourceDecoder {
	return NewXMIDecoder(options)
}
