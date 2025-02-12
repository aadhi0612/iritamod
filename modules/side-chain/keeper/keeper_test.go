package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/aadhi0612/iritamod/modules/perm"
	permtypes "github.com/aadhi0612/iritamod/modules/perm/types"
	"github.com/aadhi0612/iritamod/modules/side-chain/keeper"
	"github.com/aadhi0612/iritamod/simapp"
)

var (
	rootAdmin = sdk.AccAddress(tmhash.SumTruncated([]byte("rootAdmin")))
	accAvata  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_avata"))) // side chain user: cosmos1j0898zyz64cyxy2s2km99t2c3s6tn5tzfppw9h
	accXvata  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_xvata"))) // side chain user
	accAlice  = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_alice"))) // cosmos16877jxzrdetmzsl3pntv4n402m8d0cpvwd74w7
	accBob    = sdk.AccAddress(tmhash.SumTruncated([]byte("acc_bob")))   // cosmos1a53v8ksyd6x47sju572t48s3ynmqyan0n2c6kx

	avataSpaceId   = uint64(1) //
	avataSpaceName = "Avata Space"
	avataSpaceUri  = "https://space.avata.com"
)

type TestSuite struct {
	suite.Suite

	ctx        sdk.Context
	cdc        *codec.LegacyAmino
	keeper     keeper.Keeper
	permKeeper perm.Keeper
	app        *simapp.SimApp
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// Note: After setting up, we have:
// 1. the spaceId(1) has been created and belongs to accAvata
// 2. layer1 nft holds class `badKids` and two nfts from mocked file `mock_data/nfts.json`
// 3. class `badKids` is owned by `alice`
func (s *TestSuite) SetupTest() {
	app := simapp.Setup(false)

	s.cdc = app.LegacyAmino()
	s.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	s.app = app
	s.keeper = app.SideChainKeeper
	s.permKeeper = app.PermKeeper

	s.prepareRoles()
	s.prepareSideChain()
}

func (s *TestSuite) prepareSideChain() {
	id, err := s.keeper.CreateSpace(s.ctx, avataSpaceName, avataSpaceUri, accAvata)
	s.Require().NoError(err)
	s.Require().Equal(avataSpaceId, id)
}

func (s *TestSuite) prepareRoles() {
	err := s.permKeeper.Authorize(s.ctx, accAvata, rootAdmin, permtypes.RoleSideChainUser)
	if err != nil {
		panic("failed to authorize role")
	}
	err = s.permKeeper.Authorize(s.ctx, accXvata, rootAdmin, permtypes.RoleSideChainUser)
	if err != nil {
		panic("failed to authorize role")
	}
}
