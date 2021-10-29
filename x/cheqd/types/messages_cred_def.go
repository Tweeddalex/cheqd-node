package types

import (
	"github.com/cheqd/cheqd-node/x/cheqd/utils"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var _ IdentityMsg = &MsgCreateCredDef{}

func NewMsgCreateCredDef(id string, schemaId string, tag string, signatureType string, controller []string, value types.Any) *MsgCreateCredDef {
	return &MsgCreateCredDef{
		Id:         id,
		SchemaId:   schemaId,
		Tag:        tag,
		Type:       signatureType,
		Value:      &value,
		Controller: controller,
	}
}

func (msg *MsgCreateCredDef) GetSigners() []Signer {
	result := make([]Signer, len(msg.Controller))

	for i, signer := range msg.Controller {
		result[i] = Signer{
			Signer: signer,
		}
	}

	return result
}

func (msg *MsgCreateCredDef) GetDid() string {
	return utils.GetDidFromCredDef(msg.Id)
}

func (msg *MsgCreateCredDef) Validate(namespace string) error {
	if !utils.IsCredDef(msg.Id) {
		return ErrBadRequest.Wrap("Id must end with resource type '?service=CL-CredDef'")
	}

	if utils.IsNotDid(namespace, msg.GetDid()) {
		return ErrBadRequestIsNotDid.Wrap("Id")
	}

	if msg.Value == nil || msg.Value.Size() == 0 {
		return ErrBadRequestIsRequired.Wrap("Value")
	}

	if len(msg.SchemaId) == 0 {
		return ErrBadRequestIsRequired.Wrap("SchemaId")
	}

	if len(msg.Type) == 0 {
		return ErrBadRequestIsRequired.Wrap("SignatureType")
	}

	if utils.IsNotCredDefSignatureType(msg.Type) {
		return ErrBadRequest.Wrapf("%s is not allowed type", msg.Type)
	}

	if len(msg.Controller) == 0 {
		return ErrBadRequestIsRequired.Wrap("Controller")
	}

	if notValid, i := utils.ArrayContainsNotDid(namespace, msg.Controller); notValid {
		return ErrBadRequestIsNotDid.Wrapf("Controller item %s", msg.Controller[i])
	}

	return nil
}

func (msg *MsgCreateCredDef) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}
