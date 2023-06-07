package kvstore

import "encoding/json"

const (
	ModuleName = "kvstore"
	StoreKey   = ModuleName
)

type AppModule struct {
	// AppModule fields
}

func NewAppModule() AppModule {
	return AppModule{}
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(ModuleName, NewHandler(am.keeper))
}

func (am AppModule) QuerierRoute() string {
	return ModuleName
}

func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONMarshaler, _ json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (am AppModule) ExportGenesis(_ sdk.Context, _ codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
