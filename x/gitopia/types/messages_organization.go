package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOrganization{}

func NewMsgCreateOrganization(creator string, name string, avatarUrl string, followers string, following string, repositories string, repositoryNames string, teams string, members string, location string, email string, website string, verified string, description string, createdAt string, updatedAt string, extensions string) *MsgCreateOrganization {
	return &MsgCreateOrganization{
		Creator:         creator,
		Name:            name,
		AvatarUrl:       avatarUrl,
		Followers:       followers,
		Following:       following,
		Repositories:    repositories,
		RepositoryNames: repositoryNames,
		Teams:           teams,
		Members:         members,
		Location:        location,
		Email:           email,
		Website:         website,
		Verified:        verified,
		Description:     description,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		Extensions:      extensions,
	}
}

func (msg *MsgCreateOrganization) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrganization) Type() string {
	return "CreateOrganization"
}

func (msg *MsgCreateOrganization) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrganization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrganization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOrganization{}

func NewMsgUpdateOrganization(creator string, id uint64, name string, avatarUrl string, followers string, following string, repositories string, repositoryNames string, teams string, members string, location string, email string, website string, verified string, description string, createdAt string, updatedAt string, extensions string) *MsgUpdateOrganization {
	return &MsgUpdateOrganization{
		Id:              id,
		Creator:         creator,
		Name:            name,
		AvatarUrl:       avatarUrl,
		Followers:       followers,
		Following:       following,
		Repositories:    repositories,
		RepositoryNames: repositoryNames,
		Teams:           teams,
		Members:         members,
		Location:        location,
		Email:           email,
		Website:         website,
		Verified:        verified,
		Description:     description,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		Extensions:      extensions,
	}
}

func (msg *MsgUpdateOrganization) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOrganization) Type() string {
	return "UpdateOrganization"
}

func (msg *MsgUpdateOrganization) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOrganization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOrganization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOrganization{}

func NewMsgDeleteOrganization(creator string, id uint64) *MsgDeleteOrganization {
	return &MsgDeleteOrganization{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteOrganization) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOrganization) Type() string {
	return "DeleteOrganization"
}

func (msg *MsgDeleteOrganization) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOrganization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOrganization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
