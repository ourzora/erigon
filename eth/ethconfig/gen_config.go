// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"math/big"

	"github.com/c2h5oh/datasize"
	"github.com/erigontech/erigon-lib/chain"
	"github.com/erigontech/erigon-lib/common"
	"github.com/erigontech/erigon-lib/common/datadir"
	"github.com/erigontech/erigon-lib/downloader/downloadercfg"
	"github.com/erigontech/erigon-lib/txpool/txpoolcfg"
	"github.com/erigontech/erigon/cl/beacon/beacon_router_configuration"
	"github.com/erigontech/erigon/cl/clparams"
	"github.com/erigontech/erigon/consensus/ethash/ethashcfg"
	"github.com/erigontech/erigon/core/types"
	"github.com/erigontech/erigon/eth/gasprice/gaspricecfg"
	"github.com/erigontech/erigon/ethdb/prune"
	"github.com/erigontech/erigon/params"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                        *types.Genesis `toml:",omitempty"`
		NetworkID                      uint64
		EthDiscoveryURLs               []string
		Prune                          prune.Mode
		BatchSize                      datasize.ByteSize
		ImportMode                     bool
		BadBlockHash                   common.Hash
		Snapshot                       BlocksFreezing
		Downloader                     *downloadercfg.Cfg
		BeaconRouter                   beacon_router_configuration.RouterConfiguration
		CaplinConfig                   clparams.CaplinConfig
		Dirs                           datadir.Dirs
		ExternalSnapshotDownloaderAddr string
		Whitelist                      map[uint64]common.Hash `toml:"-"`
		Miner                          params.MiningConfig
		Ethash                         ethashcfg.Config
		Clique                         params.ConsensusSnapshotConfig
		Aura                           chain.AuRaConfig
		DeprecatedTxPool               DeprecatedTxPoolConfig
		TxPool                         txpoolcfg.Config
		GPO                            gaspricecfg.Config
		RPCGasCap                      uint64  `toml:",omitempty"`
		RPCTxFeeCap                    float64 `toml:",omitempty"`
		StateStream                    bool
		HeimdallURL                    string
		WithoutHeimdall                bool
		WithHeimdallMilestones         bool
		WithHeimdallWaypointRecording  bool
		PolygonSync                    bool
		PolygonSyncStage               bool
		Ethstats                       string
		InternalCL                     bool
		CaplinDiscoveryAddr            string
		CaplinDiscoveryPort            uint64
		CaplinDiscoveryTCPPort         uint64
		SentinelAddr                   string
		SentinelPort                   uint64
		OverridePragueTime             *big.Int `toml:",omitempty"`
		SilkwormExecution              bool
		SilkwormRpcDaemon              bool
		SilkwormSentry                 bool
		SilkwormVerbosity              string
		SilkwormNumContexts            uint32
		SilkwormRpcLogEnabled          bool
		SilkwormRpcLogDirPath          string
		SilkwormRpcLogMaxFileSize      uint16
		SilkwormRpcLogMaxFiles         uint16
		SilkwormRpcLogDumpResponse     bool
		SilkwormRpcNumWorkers          uint32
		SilkwormRpcJsonCompatibility   bool
		DisableTxPoolGossip            bool
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkID = c.NetworkID
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.Prune = c.Prune
	enc.BatchSize = c.BatchSize
	enc.ImportMode = c.ImportMode
	enc.BadBlockHash = c.BadBlockHash
	enc.Snapshot = c.Snapshot
	enc.Downloader = c.Downloader
	enc.BeaconRouter = c.BeaconRouter
	enc.CaplinConfig = c.CaplinConfig
	enc.Dirs = c.Dirs
	enc.ExternalSnapshotDownloaderAddr = c.ExternalSnapshotDownloaderAddr
	enc.Whitelist = c.Whitelist
	enc.Miner = c.Miner
	enc.Ethash = c.Ethash
	enc.Clique = c.Clique
	enc.Aura = c.Aura
	enc.DeprecatedTxPool = c.DeprecatedTxPool
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.StateStream = c.StateStream
	enc.HeimdallURL = c.HeimdallURL
	enc.WithoutHeimdall = c.WithoutHeimdall
	enc.WithHeimdallMilestones = c.WithHeimdallMilestones
	enc.WithHeimdallWaypointRecording = c.WithHeimdallWaypointRecording
	enc.PolygonSync = c.PolygonSync
	enc.PolygonSyncStage = c.PolygonSyncStage
	enc.Ethstats = c.Ethstats
	enc.InternalCL = c.InternalCL
	enc.CaplinDiscoveryAddr = c.CaplinDiscoveryAddr
	enc.CaplinDiscoveryPort = c.CaplinDiscoveryPort
	enc.CaplinDiscoveryTCPPort = c.CaplinDiscoveryTCPPort
	enc.SentinelAddr = c.SentinelAddr
	enc.SentinelPort = c.SentinelPort
	enc.OverridePragueTime = c.OverridePragueTime
	enc.SilkwormExecution = c.SilkwormExecution
	enc.SilkwormRpcDaemon = c.SilkwormRpcDaemon
	enc.SilkwormSentry = c.SilkwormSentry
	enc.SilkwormVerbosity = c.SilkwormVerbosity
	enc.SilkwormNumContexts = c.SilkwormNumContexts
	enc.SilkwormRpcLogEnabled = c.SilkwormRpcLogEnabled
	enc.SilkwormRpcLogDirPath = c.SilkwormRpcLogDirPath
	enc.SilkwormRpcLogMaxFileSize = c.SilkwormRpcLogMaxFileSize
	enc.SilkwormRpcLogMaxFiles = c.SilkwormRpcLogMaxFiles
	enc.SilkwormRpcLogDumpResponse = c.SilkwormRpcLogDumpResponse
	enc.SilkwormRpcNumWorkers = c.SilkwormRpcNumWorkers
	enc.SilkwormRpcJsonCompatibility = c.SilkwormRpcJsonCompatibility
	enc.DisableTxPoolGossip = c.DisableTxPoolGossip
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                        *types.Genesis `toml:",omitempty"`
		NetworkID                      *uint64
		EthDiscoveryURLs               []string
		Prune                          *prune.Mode
		BatchSize                      *datasize.ByteSize
		ImportMode                     *bool
		BadBlockHash                   *common.Hash
		Snapshot                       *BlocksFreezing
		Downloader                     *downloadercfg.Cfg
		BeaconRouter                   *beacon_router_configuration.RouterConfiguration
		CaplinConfig                   *clparams.CaplinConfig
		Dirs                           *datadir.Dirs
		ExternalSnapshotDownloaderAddr *string
		Whitelist                      map[uint64]common.Hash `toml:"-"`
		Miner                          *params.MiningConfig
		Ethash                         *ethashcfg.Config
		Clique                         *params.ConsensusSnapshotConfig
		Aura                           *chain.AuRaConfig
		DeprecatedTxPool               *DeprecatedTxPoolConfig
		TxPool                         *txpoolcfg.Config
		GPO                            *gaspricecfg.Config
		RPCGasCap                      *uint64  `toml:",omitempty"`
		RPCTxFeeCap                    *float64 `toml:",omitempty"`
		StateStream                    *bool
		HeimdallURL                    *string
		WithoutHeimdall                *bool
		WithHeimdallMilestones         *bool
		WithHeimdallWaypointRecording  *bool
		PolygonSync                    *bool
		PolygonSyncStage               *bool
		Ethstats                       *string
		InternalCL                     *bool
		CaplinDiscoveryAddr            *string
		CaplinDiscoveryPort            *uint64
		CaplinDiscoveryTCPPort         *uint64
		SentinelAddr                   *string
		SentinelPort                   *uint64
		OverridePragueTime             *big.Int `toml:",omitempty"`
		SilkwormExecution              *bool
		SilkwormRpcDaemon              *bool
		SilkwormSentry                 *bool
		SilkwormVerbosity              *string
		SilkwormNumContexts            *uint32
		SilkwormRpcLogEnabled          *bool
		SilkwormRpcLogDirPath          *string
		SilkwormRpcLogMaxFileSize      *uint16
		SilkwormRpcLogMaxFiles         *uint16
		SilkwormRpcLogDumpResponse     *bool
		SilkwormRpcNumWorkers          *uint32
		SilkwormRpcJsonCompatibility   *bool
		DisableTxPoolGossip            *bool
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkID != nil {
		c.NetworkID = *dec.NetworkID
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.Prune != nil {
		c.Prune = *dec.Prune
	}
	if dec.BatchSize != nil {
		c.BatchSize = *dec.BatchSize
	}
	if dec.ImportMode != nil {
		c.ImportMode = *dec.ImportMode
	}
	if dec.BadBlockHash != nil {
		c.BadBlockHash = *dec.BadBlockHash
	}
	if dec.Snapshot != nil {
		c.Snapshot = *dec.Snapshot
	}
	if dec.Downloader != nil {
		c.Downloader = dec.Downloader
	}
	if dec.BeaconRouter != nil {
		c.BeaconRouter = *dec.BeaconRouter
	}
	if dec.CaplinConfig != nil {
		c.CaplinConfig = *dec.CaplinConfig
	}
	if dec.Dirs != nil {
		c.Dirs = *dec.Dirs
	}
	if dec.ExternalSnapshotDownloaderAddr != nil {
		c.ExternalSnapshotDownloaderAddr = *dec.ExternalSnapshotDownloaderAddr
	}
	if dec.Whitelist != nil {
		c.Whitelist = dec.Whitelist
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.Ethash != nil {
		c.Ethash = *dec.Ethash
	}
	if dec.Clique != nil {
		c.Clique = *dec.Clique
	}
	if dec.Aura != nil {
		c.Aura = *dec.Aura
	}
	if dec.DeprecatedTxPool != nil {
		c.DeprecatedTxPool = *dec.DeprecatedTxPool
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.StateStream != nil {
		c.StateStream = *dec.StateStream
	}
	if dec.HeimdallURL != nil {
		c.HeimdallURL = *dec.HeimdallURL
	}
	if dec.WithoutHeimdall != nil {
		c.WithoutHeimdall = *dec.WithoutHeimdall
	}
	if dec.WithHeimdallMilestones != nil {
		c.WithHeimdallMilestones = *dec.WithHeimdallMilestones
	}
	if dec.WithHeimdallWaypointRecording != nil {
		c.WithHeimdallWaypointRecording = *dec.WithHeimdallWaypointRecording
	}
	if dec.PolygonSync != nil {
		c.PolygonSync = *dec.PolygonSync
	}
	if dec.PolygonSyncStage != nil {
		c.PolygonSyncStage = *dec.PolygonSyncStage
	}
	if dec.Ethstats != nil {
		c.Ethstats = *dec.Ethstats
	}
	if dec.InternalCL != nil {
		c.InternalCL = *dec.InternalCL
	}
	if dec.CaplinDiscoveryAddr != nil {
		c.CaplinDiscoveryAddr = *dec.CaplinDiscoveryAddr
	}
	if dec.CaplinDiscoveryPort != nil {
		c.CaplinDiscoveryPort = *dec.CaplinDiscoveryPort
	}
	if dec.CaplinDiscoveryTCPPort != nil {
		c.CaplinDiscoveryTCPPort = *dec.CaplinDiscoveryTCPPort
	}
	if dec.SentinelAddr != nil {
		c.SentinelAddr = *dec.SentinelAddr
	}
	if dec.SentinelPort != nil {
		c.SentinelPort = *dec.SentinelPort
	}
	if dec.OverridePragueTime != nil {
		c.OverridePragueTime = dec.OverridePragueTime
	}
	if dec.SilkwormExecution != nil {
		c.SilkwormExecution = *dec.SilkwormExecution
	}
	if dec.SilkwormRpcDaemon != nil {
		c.SilkwormRpcDaemon = *dec.SilkwormRpcDaemon
	}
	if dec.SilkwormSentry != nil {
		c.SilkwormSentry = *dec.SilkwormSentry
	}
	if dec.SilkwormVerbosity != nil {
		c.SilkwormVerbosity = *dec.SilkwormVerbosity
	}
	if dec.SilkwormNumContexts != nil {
		c.SilkwormNumContexts = *dec.SilkwormNumContexts
	}
	if dec.SilkwormRpcLogEnabled != nil {
		c.SilkwormRpcLogEnabled = *dec.SilkwormRpcLogEnabled
	}
	if dec.SilkwormRpcLogDirPath != nil {
		c.SilkwormRpcLogDirPath = *dec.SilkwormRpcLogDirPath
	}
	if dec.SilkwormRpcLogMaxFileSize != nil {
		c.SilkwormRpcLogMaxFileSize = *dec.SilkwormRpcLogMaxFileSize
	}
	if dec.SilkwormRpcLogMaxFiles != nil {
		c.SilkwormRpcLogMaxFiles = *dec.SilkwormRpcLogMaxFiles
	}
	if dec.SilkwormRpcLogDumpResponse != nil {
		c.SilkwormRpcLogDumpResponse = *dec.SilkwormRpcLogDumpResponse
	}
	if dec.SilkwormRpcNumWorkers != nil {
		c.SilkwormRpcNumWorkers = *dec.SilkwormRpcNumWorkers
	}
	if dec.SilkwormRpcJsonCompatibility != nil {
		c.SilkwormRpcJsonCompatibility = *dec.SilkwormRpcJsonCompatibility
	}
	if dec.DisableTxPoolGossip != nil {
		c.DisableTxPoolGossip = *dec.DisableTxPoolGossip
	}
	return nil
}
